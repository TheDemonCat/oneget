package downloader

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"strings"
	"time"
)

type HtmlParser struct {
	HtmlParserConfig
}

type HtmlParserConfig struct {
	ReleaseTableSelector   string // [id$='actualTable']
	ProjectTableSelector   string // [id$='versionsTable']
	ProjectReleaseSelector string // ".files-container .formLine a"
	ReleaseFilesSelector   string // ".downloadDist a"

}

var defaultConfig = HtmlParserConfig{

	ReleaseTableSelector:   "[id$='actualTable']",
	ProjectTableSelector:   "[id$='versionsTable']",
	ProjectReleaseSelector: ".files-container .formLine a",
	ReleaseFilesSelector:   ".downloadDist a",
}

func NewHtmlParser(config ...HtmlParserConfig) (*HtmlParser, error) {

	cfg := defaultConfig

	if len(config) == 1 {
		cfg = config[0]
	}

	return &HtmlParser{
		cfg,
	}, nil
}

func (p *HtmlParser) ParseTotalReleases(body io.Reader) (rows []ProjectInfo, err error) {

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}

	var tableHtml *goquery.Selection

	doc.Find(p.ReleaseTableSelector).Each(func(i int, html *goquery.Selection) {
		tableHtml = html
		return
	})

	if tableHtml == nil {
		return
	}

	return parseReleasesTable(tableHtml), nil

}

func (p *HtmlParser) ParseProjectReleases(body io.Reader) (rows []*ProjectVersionInfo, err error) {

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}

	var tableHtml *goquery.Selection

	doc.Find(p.ProjectTableSelector).Each(func(i int, html *goquery.Selection) {
		tableHtml = html
		return
	})

	if tableHtml == nil {
		return nil, fmt.Errorf("not found html tag by selector: <%s>", p.ProjectTableSelector)
	}

	return parseProjectTable(tableHtml), nil

}

func (p *HtmlParser) ParseProjectRelease(body io.Reader) (rows []ReleaseFileInfo, err error) {

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}

	doc.Find(p.ProjectReleaseSelector).Each(func(i int, releaseFileHtml *goquery.Selection) {
		name := strings.TrimSpace(releaseFileHtml.Text())
		url := releaseFileHtml.AttrOr("href", "")
		if strings.HasPrefix(url, "/version_file") {
			rows = append(rows, ReleaseFileInfo{
				name: name,
				url:  url,
			})
		}

	})

	return

}

func (p *HtmlParser) ParseReleaseFiles(body io.Reader) (rows []string, err error) {

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}

	doc.Find(p.ReleaseFilesSelector).Each(func(i int, html *goquery.Selection) {
		ref, ok := html.Attr("href")
		if !ok {
			return
		}
		rows = append(rows, ref)
	})

	return

}

type ReleaseFileInfo struct {
	name string
	url  string
}

type ProjectInfo struct {
	ID               string
	Group            string
	GroupID          string
	Name             string
	Url              string
	VersionsInfo     []*ProjectVersionInfo
	TestVersionsInfo []*ProjectVersionInfo
}

type ProjectVersionInfo struct {
	Url         string
	Name        string
	PublishDate time.Time
}

func parseReleasesTable(s *goquery.Selection) (rows []ProjectInfo) {

	s.Find("[group]").Each(func(i int, groupInfo *goquery.Selection) {
		groupId, _ := groupInfo.Attr("group")
		groupName := groupInfo.Find(".group-name").Text()

		log.Debugf("Group <%s> - <%s>", groupId, groupName)

		s.Find(fmt.Sprintf("[parent-group$='%s']", groupId)).Each(func(_ int, releaseRow *goquery.Selection) {

			info := ProjectInfo{
				Group:   groupName,
				GroupID: groupId,
			}

			info.Name = strings.TrimSpace(releaseRow.Find(".nameColumn").Text())
			info.Url, _ = releaseRow.Find(".nameColumn a").Attr("href")
			info.ID = strings.TrimLeft(info.Url, "/project/")

			releaseRow.Find("td").Each(func(i int, rowHtml *goquery.Selection) {

				switch i {
				case 1: // .versionColumn .actualVersionColumn
					rowHtml.Find("a").Each(func(i int, releaseHtml *goquery.Selection) {
						releaseInfo := &ProjectVersionInfo{
							Url:  releaseHtml.AttrOr("href", ""),
							Name: strings.TrimSpace(releaseHtml.Text()),
						}
						info.VersionsInfo = append(info.VersionsInfo, releaseInfo)
					})
				case 2: // .releaseDate
					rowHtml.Find("span").Each(func(i int, releaseHtml *goquery.Selection) {
						if len(info.VersionsInfo) > i {
							textDate := strings.TrimSpace(releaseHtml.Text())
							publishDate, err := parseReleaseDate(textDate)
							if err != nil {
								log.Errorf("Error parse <%s> for %s", textDate, info.VersionsInfo[i].Name)
							}
							info.VersionsInfo[i].PublishDate = publishDate
						}
					})
				case 6: // .versionColumn
					rowHtml.Find("a").Each(func(i int, releaseHtml *goquery.Selection) {
						releaseInfo := &ProjectVersionInfo{
							Url:  releaseHtml.AttrOr("href", ""),
							Name: strings.TrimSpace(releaseHtml.Text()),
						}
						info.TestVersionsInfo = append(info.TestVersionsInfo, releaseInfo)
					})
				case 7: // .publicationDate
					rowHtml.Find("span").Each(func(i int, releaseHtml *goquery.Selection) {
						if len(info.TestVersionsInfo) > i {
							textDate := strings.TrimSpace(releaseHtml.Text())
							publishDate, err := parseReleaseDate(textDate)
							if err != nil {
								log.Errorf("Error parse <%s> for %s", textDate, info.TestVersionsInfo[i].Name)
							}
							info.TestVersionsInfo[i].PublishDate = publishDate
						}
					})
				}

			})

			rows = append(rows, info)

		})

	})

	return

}

func parseProjectTable(s *goquery.Selection) (rows []*ProjectVersionInfo) {

	s.Find("tr").Each(func(i int, releaseRow *goquery.Selection) {

		if releaseRow.Parent().Is("thead") {
			return
		}

		releaseInfo := &ProjectVersionInfo{}
		versionColumn := releaseRow.Find(".versionColumn a")
		releaseInfo.Name = strings.TrimSpace(versionColumn.Text())
		releaseInfo.Url = versionColumn.AttrOr("href", "")

		textDate := strings.TrimSpace(releaseRow.Find(".dateColumn").Text())
		var err error
		if releaseInfo.PublishDate, err = parseReleaseDate(textDate); err != nil {
			log.Errorf("Error parse <%s> for %s <%s>", textDate, releaseInfo.Name, releaseInfo.Url)
		}

		rows = append(rows, releaseInfo)

	})

	return

}

func parseReleaseDate(textDate string) (time.Time, error) {

	date, err := time.Parse("02.01.06", textDate)
	return date, err

}
