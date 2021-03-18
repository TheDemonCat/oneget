package downloader

import (
	"fmt"
	"go.uber.org/multierr"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type GetConfig struct {
	BasePath string
	Project  string
	Version  VersionFilter
	Filters  []FileFilter
}

type OnegetDownloader struct {
	Login, Password string

	cookie *cookiejar.Jar

	client *Client

	urlCh chan *FileToDownload
	wg    *sync.WaitGroup

	parser *HtmlParser
}

func NewDownloader(login, password string) *OnegetDownloader {

	cj, _ := cookiejar.New(nil)
	parser, _ := NewHtmlParser()

	return &OnegetDownloader{
		cookie:   cj,
		Login:    login,
		Password: password,
		wg:       &sync.WaitGroup{},
		parser:   parser,
	}

}

func (dr *OnegetDownloader) Get(config ...GetConfig) ([]string, error) {

	client, err := NewClient(loginURL, releasesURL, dr.Login, dr.Password)
	if err != nil {
		return nil, err
	}

	dr.client = client

	files := make([]string, 0)

	downloadCh := make(chan *FileToDownload, 100)

	for _, getConfig := range config {
		err := dr.getFiles(getConfig, downloadCh)
		if err != nil {
			return nil, err
		}
	}

	go func() {
		dr.wg.Wait()
		close(downloadCh)
	}()

	limit := make(chan struct{}, 10)
	mu := &sync.Mutex{}
	for fileToDownload := range downloadCh {
		go func(file *FileToDownload) {

			limit <- struct{}{}

			filename, err := dr.downloadFile(file)
			if err != nil {
				dr.wg.Done()
				log.Errorf(err.Error())
			}
			if len(filename) > 0 {
				mu.Lock()
				files = append(files, filename)
				mu.Unlock()
			}
			dr.wg.Done()

			<-limit

		}(fileToDownload)

	}

	downloadCh = nil

	return files, nil

}

// GetListProject ShowUnavailablePrograms bool // reverse for total/hideUnavailablePrograms=true
func (dr *OnegetDownloader) GetListProject(showUnavailablePrograms bool) ([]ProjectInfo, error) {

	client, err := NewClient(loginURL, releasesURL, dr.Login, dr.Password)
	if err != nil {
		return nil, err
	}

	dr.client = client

	totalurl := fmt.Sprintf("%s/total?hideUnavailablePrograms=%s", releasesURL,
		strconv.FormatBool(!showUnavailablePrograms))

	resp, err := dr.client.Get(totalurl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	list, err := dr.parser.ParseTotalReleases(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error parse total releases: %s", err.Error())
	}
	return list, err
}

func (dr *OnegetDownloader) getFiles(config GetConfig, downloadCh chan *FileToDownload) error {

	releases, err := dr.getProjectReleases(config)
	if err != nil {
		return err
	}

	for _, release := range releases {
		dr.wg.Add(1)
		go func(info *ProjectVersionInfo, cfg GetConfig) {
			_ = dr.getReleaseFiles(info, config, downloadCh)
			dr.wg.Done()
		}(release, config)
	}

	return nil
}

func (dr *OnegetDownloader) getReleaseFiles(release *ProjectVersionInfo, config GetConfig, downloadCh chan *FileToDownload) error {

	client := dr.client
	resp, err := client.Get(releasesURL + release.Url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	releaseFiles, err := dr.parser.ParseProjectRelease(resp.Body)
	if err != nil {
		return err
	}
	files := filterReleaseFiles(releaseFiles, config.Filters)

	var merr error

	for _, file := range files {

		err := dr.addFileToChannel(file.url, config, downloadCh)
		if err != nil {
			log.Errorf("Error get file from <%s>: %s", err.Error())
			multierr.Append(merr, err)
		}

	}

	return merr

}

func (dr *OnegetDownloader) getProjectReleases(config GetConfig) ([]*ProjectVersionInfo, error) {

	resp, err := dr.client.Get(projectHrefPrefix + config.Project)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	releases, err := dr.parser.ParseProjectReleases(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error parse project <%s> releases: %s, html: <%s>",
			config.Project, err.Error(), readBodyMustString(resp.Body))
	}

	return filterProjectVersionInfo(releases, config.Version), nil

}
func filterReleaseFiles(list []ReleaseFileInfo, filters []FileFilter) (filteredList []ReleaseFileInfo) {

	if len(filters) == 0 || len(list) == 0 {
		return list
	}

	matchInfo := func(i ReleaseFileInfo) bool {

		for _, filter := range filters {

			matchName := filter.MatchString(i.name)
			matchUrl := filter.MatchString(i.url)

			if matchName || matchUrl {
				return true
			}

		}

		return false
	}

	for _, info := range list {

		if matchInfo(info) {
			filteredList = append(filteredList, info)
		}

	}

	return

}

func filterProjectVersionInfo(list []*ProjectVersionInfo, filter VersionFilter) (filteredList []*ProjectVersionInfo) {

	if len(list) == 0 {
		return list
	}

	return filter.Filter(list)

}

func (dr *OnegetDownloader) getClient() *http.Client {
	return &http.Client{
		Jar:       dr.cookie,
		Transport: nil,
	}
}

func (dr *OnegetDownloader) addFileToChannel(href string, config GetConfig, downloadCh chan *FileToDownload) (err error) {

	downloadHref := []string{releasesURL + href}

	if !directLink(href) {
		downloadHref, err = dr.getDownloadFileLinks(href, config)
		if err != nil {
			return err
		}
	}

	fileName, filePath, err := dr.fileNameFromUrl(href)
	if err != nil {
		return err
	}

	if len(downloadHref) == 0 {
		return nil
	}

	dr.wg.Add(1)

	log.Debugf("Add to download: %s", href)
	downloadCh <- &FileToDownload{
		url:      downloadHref,
		path:     filePath,
		name:     fileName,
		basePath: config.BasePath,
	}

	return nil
}

func directLink(href string) bool {
	return strings.HasSuffix(href, "txt") ||
		strings.HasSuffix(href, "pdf") ||
		strings.HasSuffix(href, "html") ||
		strings.HasSuffix(href, "htm") ||
		(strings.HasSuffix(href, "zip") &&
			strings.Contains(href, "path=ro\\"))
}

func (dr *OnegetDownloader) fileNameFromUrl(rawUrl string) (string, string, error) {

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

func (dr *OnegetDownloader) downloadFile(fileToDownload *FileToDownload) (string, error) {

	workDir := filepath.Join(fileToDownload.basePath, strings.ToLower(fileToDownload.path))
	fileName := filepath.Join(workDir, fileToDownload.name)
	_, err := os.Stat(fileName)
	if os.IsExist(err) {
		return fileName, nil
	}
	if os.IsNotExist(err) {

		if _, err := os.Stat(workDir); os.IsNotExist(err) {
			err = os.MkdirAll(filepath.Join(workDir), 0777)
			if err != nil {
				return "", err
			}
			// https://wenzr.wordpress.com/2018/03/27/go-file-permissions-on-unix/
			_ = os.Chmod(workDir, 0777)
		}

		log.Infof("Getting a file: %s", fileToDownload.name)

		downloadUrl := fileToDownload.url[0]

		log.Debugf("Getting a file from url: %s", downloadUrl)
		client := dr.client
		resp, err := client.Get(downloadUrl)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		if err := SaveToFile(resp.Body, fileName+tempFileSuffix); err != nil {
			log.Debugf("File <%s> saved err: %s", fileName, err.Error())
			return "", err
		}

		log.Debugf("File saved to: %s", fileName)

		err = os.Rename(fileName+tempFileSuffix, fileName)
		if err != nil {
			return "", err
		}

		return fileName, nil

	}

	return "", nil

}

func SaveToFile(reader io.ReadCloser, fileName string) error {
	fd, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer reader.Close()
	defer fd.Close()

	if _, err = io.Copy(fd, reader); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (dr *OnegetDownloader) handleError(err error) {
	if err == nil {
		return
	}
	log.Error(err.Error())
}

func (dr *OnegetDownloader) getDownloadFileLinks(href string, _ GetConfig) ([]string, error) {

	client := dr.client
	resp, err := client.Get(href)
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()

	fileLinks, err := dr.parser.ParseReleaseFiles(resp.Body)
	if err != nil {
		return nil, err
	}

	return fileLinks, nil

}
