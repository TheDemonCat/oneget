package downloader

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/khorevaa/logos"
	"github.com/v8platform/oneget/unpacker"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

var releasesURL = "https://releases.1c.ru"
var loginURL = "https://login.1c.ru"

const projectHrefPrefix = "/project/"
const versionFilesHrefPrefix = "/version_files"
const fileServerHrefPrefix = "/public/file/get"

const tempFileSuffix = ".d1c"

var semaMaxConnections = make(chan struct{}, 10)

var log = logos.New("github.com/v8platform/oneget/downloader").Sugar()

type FileToDownload struct {
	url      []string
	basePath string
	path     string
	name     string
}

type Config struct {
	Login         string
	Password      string
	BasePath      string
	StartDate     time.Time
	Nicks         map[string]bool
	VersionFilter string
	DistribFilter string
	Extract       bool
	ExtractPath   string
	Rename        bool
}


type Downloader struct {
	Config
	httpClient *http.Client
	urlCh      chan *FileToDownload
	wg         sync.WaitGroup
}

func New(config Config) *Downloader {

	cj, _ := cookiejar.New(nil)

	return &Downloader{
		Config: config,
		httpClient: &http.Client{
			Jar: cj,
		},
		wg: sync.WaitGroup{},
	}

}

func (dr *Downloader) Get() ([]os.FileInfo, error) {

	files := make([]os.FileInfo, 0)

	ticketUrl, err := dr.getURL()
	if err != nil {
		dr.handleError(err)
		return files, err
	}

	dr.urlCh = make(chan *FileToDownload, 10000)
	dr.wg.Add(1)
	go dr.findLinks(ticketUrl, dr.findProject)
	go func() {
		dr.wg.Wait()
		close(dr.urlCh)
	}()

	for fileToDownload := range dr.urlCh {
		fileInfo, err := dr.downloadFile(fileToDownload)
		if err != nil {
			return nil, err
		}
		if fileInfo != nil {
			files = append(files, fileInfo)
		}
	}

	dr.urlCh = nil

	return files, nil

}

func (dr *Downloader) getURL() (string, error) {

	type loginParams struct {
		Login       string `json:"login"`
		Password    string `json:"password"`
		ServiceNick string `json:"serviceNick"`
	}

	type ticket struct {
		Ticket string `json:"ticket"`
	}

	postBody, err := json.Marshal(
		loginParams{dr.Login, dr.Password, releasesURL})
	if err != nil {
		return "", err
	}

	acquireSemaConnections()
	resp, err := dr.httpClient.Post(
		loginURL+"/rest/public/ticket/get",
		"application/json",
		bytes.NewReader(postBody))
	releaseSemaConnections()
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	responseBodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%s", string(responseBodyData))
	}

	var tick ticket
	err = json.Unmarshal(responseBodyData, &tick)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(loginURL+"/ticket/auth?token=%s", tick.Ticket), nil

}

func (dr *Downloader) findLinks(rawUrl string, f func(string, string, *html.Node)) {

	defer dr.wg.Done()

	acquireSemaConnections()
	resp, err := dr.httpClient.Get(rawUrl)
	releaseSemaConnections()

	if err != nil {
		dr.handleError(err)
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		dr.handleError(err)
		return
	}

	dr.eachNode(doc, rawUrl, f)

}

func (dr *Downloader) eachNode(node *html.Node, u string, f func(string, string, *html.Node)) {

	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				f(u, attr.Val, node)
				break
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		dr.eachNode(c, u, f)
	}

}

func (dr *Downloader) findProject(_, href string, _ *html.Node) {
	projectName := strings.ToLower(strings.TrimLeft(href, projectHrefPrefix))
	log.Debugf("Finding project in href %s", projectName)
	if (dr.Nicks == nil && strings.HasPrefix(href, projectHrefPrefix)) || dr.Nicks[projectName] {
		dr.wg.Add(1)
		go dr.findLinks(releasesURL+href, dr.findVersion)
	}

}

func (dr *Downloader) findVersion(_, href string, node *html.Node) {

	if strings.HasPrefix(href, versionFilesHrefPrefix) {

		vDateRaw := strings.Trim(node.Parent.NextSibling.NextSibling.FirstChild.Data, " \n")
		vDate, err := time.Parse("02.01.06", vDateRaw)
		if err != nil {
			//dr.handleError(err)
			return
		}

		if dr.VersionFilter != "" {
			log.Debugf("Filtering href %s by %s", href, dr.VersionFilter)
			matched, err := regexp.MatchString(dr.VersionFilter, href)
			if err != nil {
				dr.handleError(err)
				return
			}
			if !matched {
				log.Debugf("Href %s SKIP", href)
				return
			}
		}

		if vDate.After(dr.StartDate) {
			dr.wg.Add(1)
			go dr.findLinks(releasesURL+href, dr.findToDownloadLink)
		}

	}

}

func (dr *Downloader) findToDownloadLink(_, href string, _ *html.Node) {

	lowerHref := strings.ToLower(href)
	isRO := strings.Contains(lowerHref, "path=ro\\")

	if dr.DistribFilter == "" {
		if strings.HasSuffix(lowerHref, "rar") ||
			(strings.HasSuffix(lowerHref, "zip") && !isRO) ||
			strings.HasSuffix(lowerHref, "gz") ||
			strings.HasSuffix(lowerHref, "exe") ||
			strings.HasSuffix(lowerHref, "msi") ||
			strings.HasSuffix(lowerHref, "deb") ||
			strings.HasSuffix(lowerHref, "rpm") ||
			strings.HasSuffix(lowerHref, "epf") ||
			strings.HasSuffix(lowerHref, "erf") {

			dr.wg.Add(1)
			dr.findLinks(releasesURL+href, dr.findFileServerLink)

		} else if strings.HasSuffix(lowerHref, "txt") ||
			strings.HasSuffix(lowerHref, "pdf") ||
			strings.HasSuffix(lowerHref, "html") ||
			strings.HasSuffix(lowerHref, "htm") ||
			(strings.HasSuffix(lowerHref, "zip") && isRO) {
			dr.addFileToChannel(href, releasesURL+href)
		}
	} else {
		matched, err := regexp.MatchString(dr.DistribFilter, href)
		if err != nil {
			dr.handleError(err)
			return
		}
		if matched {
			dr.wg.Add(1)
			dr.findLinks(releasesURL+href, dr.findFileServerLink)
		}
	}

}

func (dr *Downloader) findFileServerLink(u, href string, _ *html.Node) {

	if strings.Contains(href, fileServerHrefPrefix) {
		dr.addFileToChannel(u, href)
	}

}

func (dr *Downloader) addFileToChannel(u, href string) {
	fileName, filePath, err := dr.fileNameFromUrl(u)
	if err == nil {
		fileToDownload := FileToDownload{
			url:  []string{href},
			path: filePath,
			name: fileName,
		}
		dr.urlCh <- &fileToDownload
	} else {
		dr.handleError(err)
	}
}

func (dr *Downloader) fileNameFromUrl(rawUrl string) (string, string, error) {

	fileName := strings.Builder{}
	filePath := strings.Builder{}

	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return "", "", err
	}

	query, err := url.ParseQuery(parsedUrl.RawQuery)
	if err != nil {
		return "", "", err
	}

	path := strings.Split(query.Get("path"), "\\")
	fileName.WriteString(path[len(path)-1])

	nick := query.Get("nick")
	ver := query.Get("ver")

	filePath.WriteString(nick)
	filePath.WriteRune(os.PathSeparator)
	filePath.WriteString(ver)
	filePath.WriteRune(os.PathSeparator)

	return fileName.String(), filePath.String(), nil
}

func (dr *Downloader) downloadFile(fileToDownload *FileToDownload) (os.FileInfo, error) {

	workDir := filepath.Join(dr.BasePath, strings.ToLower(fileToDownload.path))
	fileName := filepath.Join(workDir, fileToDownload.name)
	fileInfo, err := os.Stat(fileName)
	if os.IsExist(err) {

		return fileInfo, nil

	} else if os.IsNotExist(err) {

		if _, err := os.Stat(workDir); os.IsNotExist(err) {
			err = os.MkdirAll(filepath.Join(workDir), 0777)
			if err != nil {
				return nil, err
			}
			// https://wenzr.wordpress.com/2018/03/27/go-file-permissions-on-unix/
			os.Chmod(workDir, 0777)
		}
		log.Debugf("Workspace directory: %s", workDir)
		log.Debugf("Getting a file from url: %s", fileToDownload.url)

		acquireSemaConnections()
		resp, err := dr.httpClient.Get(fileToDownload.url[0])
		if err != nil {
			return nil, err
		}

		f, err := os.Create(fileName + tempFileSuffix)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		_, err = io.Copy(f, resp.Body)
		releaseSemaConnections()
		if err != nil {
			return nil, err
		}

		f.Close()

		log.Debugf("End of receiving file by url: %s", fileToDownload.url)
		log.Debugf("File saved to: %s", fileName)
		log.Debugf("End of receiving file by url: %s", fileToDownload.url)
		log.Debugf("File saved to: %s", fileName)

		err = os.Rename(fileName+tempFileSuffix, fileName)
		if err != nil {
			return nil, err
		}

		if dr.Extract {
			extractDir := dr.ExtractPath
			unpacker.Extract(fileName, extractDir)

			if dr.Rename {
				files, err := ioutil.ReadDir(extractDir)
				if err != nil {
					return nil, err
				}
				for _, file := range files {
					if file.IsDir() {
						continue
					}
					oldName := file.Name()
					newName := unpacker.GetAliasesDistrib(oldName)
					err := os.Rename(
							filepath.Join(extractDir, oldName),
							filepath.Join(extractDir, newName))
					if err != nil {
						return nil, err
					}
				}

			}
		}

		fileInfo, err := os.Stat(fileName)
		if err != nil {
			return nil, err
		}

		return fileInfo, nil

	} else if err != nil {
		log.Debugf("Error download files: %s", err)
	}

	return nil, nil

}

func (dr *Downloader) handleError(err error) {
	if err == nil {
		return
	}
	log.Error(err.Error())
}

func acquireSemaConnections() {
	semaMaxConnections <- struct{}{}
}

func releaseSemaConnections() {
	_ = <-semaMaxConnections
}
