package downloader

import (
	"fmt"
	"github.com/xelaj/go-dry"
	"regexp"
	"sort"
	"strings"
	"time"
)

const (
	Platform83Project = "Platform83"
	EDTProject        = "DevelopmentTools10"
)
const (
	x64re     = "(?smU)(?:64-bit|64 бит).*"
	rpmre     = "(?smU)(?:RPM|ОС Linux|для Linux$|tar.bz2).*"
	debre     = "(?smU)(?:DEB|ОС Linux|для Linux$|tar.bz2).*"
	windowsre = "(?smU)(?:Windows|ОС Windows|zip).*"
	osxre     = "(?smU)(?:OS X|macOS|MacOS|ОС macOS).*"

	clientre      = "(?smU)Клиент"
	serverre      = "(?smU)(?:Cервер|Сервер)"
	thinre        = "(?smU)Тонкий клиент"
	fullwindowsre = "(?smU)Технологическая платформа"
)

/*

	Специальные фильтры для платформы и ряда других приложение

	platform:thin.mac@latest
	platform:full.win.x64@latest
	platform:win.x64@latest
	platform:server.deb.x64@latest


*/

var (
	ProjectAliases = map[string]string{
		"platform": Platform83Project,
		"edt":      EDTProject,
		"ring":     "EnterpriseLicenseTools",
		"executor": "Executor",
		"pg":       "AddCompPostgre",
	}

	shortFilters = map[string]string{
		"mac":         osxre,
		"windows":     windowsre,
		"win":         windowsre,
		"deb":         debre,
		"rpm":         rpmre,
		"x64":         x64re,
		"client":      clientre,
		"server":      serverre,
		"thin-client": thinre,
		"thin":        thinre,
		"online":      "edt.online",
		"jdk":         "edt.jdk",
		"full":        fullwindowsre,
	}

	x64Regexp = regexp.MustCompile(x64re)
)

type FileFilter interface {
	MatchString(source string) bool
}

type VersionFilter interface {
	Filter(source []*ProjectVersionInfo) (result []*ProjectVersionInfo)
}

type PlatformMatchFilter struct {
	filters      []*regexp.Regexp
	x64bitMatch  bool
	x64bitRegexp *regexp.Regexp
}

func GetProjectIDByAlias(alias string) string {

	if val, ok := ProjectAliases[alias]; ok {
		return val
	}

	return alias

}

func NewFileFilter(project string, filter string) (FileFilter, error) {

	switch project {
	case Platform83Project:
		return newPlatformMatchFilter(filter)
	case EDTProject:
		return newEdtFilter(filter)
	default:
		return nil, fmt.Errorf("unknown filter builder for project <%s>", project)
	}
}

func NewFileFilterMust(project string, filter string) FileFilter {

	newFilter, err := NewFileFilter(project, filter)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	return newFilter
}

func newPlatformMatchFilter(filter string) (*PlatformMatchFilter, error) {

	m := &PlatformMatchFilter{
		x64bitMatch:  false,
		x64bitRegexp: x64Regexp,
	}

	filters := strings.Split(filter, ".")
	filters = removeFromFilter(filters, "x32")

	if ok := dry.StringInSlice("x64", filters); ok {
		filters = removeFromFilter(filters, "x64")
		m.x64bitMatch = true
	}

	// Для Windows если стоит только фильтр по нему и другого нет, то установим для платформы скачиваем полного дистрибутива
	if ok := dry.StringInSlice("win", filters) || dry.StringInSlice("windows", filters); ok && len(filters) == 1 {
		filters = append(filters, "full")
	}

	if ok := dry.StringInSlice("mac", filters); ok {
		filters = removeFromFilter(filters, "server")
		m.x64bitMatch = false
	}

	err := m.build(filters)

	return m, err
}

func (m *PlatformMatchFilter) MatchString(source string) bool {

	if m.x64bitRegexp.MatchString(source) != m.x64bitMatch {
		return false
	}

	for _, filter := range m.filters {

		if ok := filter.MatchString(source); !ok {
			return false
		}

	}

	return true
}

func (m *PlatformMatchFilter) build(filters []string) error {

	for _, filter := range filters {

		val, ok := shortFilters[filter]

		if !ok {
			return fmt.Errorf("unknown <%s> filter", filter)
		}

		m.filters = append(m.filters, regexp.MustCompile(val))

	}

	return nil
}

func removeFromFilter(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

type EdtMatchFilter struct {
	filters        []*regexp.Regexp
	matchOffline   bool
	matchOfflineRe []*regexp.Regexp
	distrRe        []*regexp.Regexp
}

func newEdtFilter(filter string) (*EdtMatchFilter, error) {
	m := &EdtMatchFilter{
		matchOffline: true,
		matchOfflineRe: []*regexp.Regexp{
			regexp.MustCompile(".*оффлайн установки.*"),
			regexp.MustCompile(".*1c_edt_distr_offline.*"),
		},
		distrRe: []*regexp.Regexp{
			regexp.MustCompile(".*1C:EDT*"),
			regexp.MustCompile(".*1c_edt_distr.*"),
		},
	}

	filters := strings.Split(filter, ".")
	filters = removeFromFilter(filters, "x32")
	filters = removeFromFilter(filters, "x64")
	if ok := dry.StringInSlice("online", filters); ok {
		m.matchOffline = false
		filters = removeFromFilter(filters, "online")
	}

	if ok := dry.StringInSlice("jdk", filters); ok {
		m.filters = append(m.filters, regexp.MustCompile(".*JDK.*"))
		filters = removeFromFilter(filters, "jdk")
	}

	err := m.build(filters)
	if err != nil {
		return nil, err
	}

	return m, err
}

func (m *EdtMatchFilter) build(filters []string) error {

	for _, filter := range filters {

		val, ok := shortFilters[filter]

		if !ok {
			return fmt.Errorf("unknown <%s> filter", filter)
		}

		m.filters = append(m.filters, regexp.MustCompile(val))

	}

	return nil
}

func (m *EdtMatchFilter) isDistrFile(source string) bool {
	for _, filter := range m.distrRe {

		if ok := filter.MatchString(source); ok {
			return true
		}

	}

	return false
}

func (m *EdtMatchFilter) isOfflineDistrFile(source string) bool {
	for _, filter := range m.matchOfflineRe {

		if ok := filter.MatchString(source); ok {
			return true
		}

	}

	return false
}

func (m *EdtMatchFilter) MatchString(source string) bool {

	for _, filter := range m.filters {

		if ok := filter.MatchString(source); !ok {
			return false
		}

	}

	if m.isDistrFile(source) {

		offlineDistr := m.isOfflineDistrFile(source)

		if m.matchOffline && offlineDistr {
			return true
		}

		if m.matchOffline && !offlineDistr {
			return false
		}

	}

	return true
}

func NewVersionFilter(_ string, filter string) (VersionFilter, error) {

	switch {
	case strings.HasPrefix(filter, "from:"):
		return newVersionDateFilter(filter)
	case strings.HasPrefix(filter, "from-v:"):
		return newVersionFromFilter(filter)
	case strings.HasPrefix(filter, "latest"):
		return newLatestVersionFilter(filter)
	default:
		return newVersionFilter(filter)
	}
}

func NewVersionFilterMust(project string, filter string) VersionFilter {

	newFilter, err := NewVersionFilter(project, filter)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	return newFilter
}

type versionFilter struct {
	filter *regexp.Regexp
}

func (m *versionFilter) Filter(source []*ProjectVersionInfo) (result []*ProjectVersionInfo) {

	eachVersion(source, func(versionInfo *ProjectVersionInfo) {
		if m.filter.MatchString(versionInfo.Name) {
			result = append(result, versionInfo)
		}
	})

	return
}

func eachVersion(source []*ProjectVersionInfo, fn func(versionInfo *ProjectVersionInfo)) {
	for _, info := range source {
		fn(info)
	}
}

func newVersionFilter(filter string) (*versionFilter, error) {

	return &versionFilter{
		regexp.MustCompile(filter),
	}, nil

}

type VersionDateFilter struct {
	date time.Time
}

func (m *VersionDateFilter) Filter(source []*ProjectVersionInfo) (result []*ProjectVersionInfo) {

	eachVersion(source, func(versionInfo *ProjectVersionInfo) {
		if versionInfo.PublishDate.After(m.date) {
			result = append(result, versionInfo)
		}
	})
	return
}

func newVersionDateFilter(filter string) (*VersionDateFilter, error) {

	values := strings.SplitN(filter, ":", 2)

	if len(values) == 1 {
		return nil, fmt.Errorf("error date filter: no contain date value")
	}

	data, err := time.Parse("02.01.06", values[1])
	if err != nil {
		return nil, fmt.Errorf("error parse date: %s", err.Error())
	}

	return &VersionDateFilter{
		data,
	}, nil

}

type VersionFromFilter struct {
	version string
}

func (m *VersionFromFilter) Filter(source []*ProjectVersionInfo) (result []*ProjectVersionInfo) {

	eachVersion(source, func(versionInfo *ProjectVersionInfo) {
		if compareVersion(versionInfo.Name, m.version) >= 0 {
			result = append(result, versionInfo)
		}
	})
	return
}

func newVersionFromFilter(filter string) (*VersionFromFilter, error) {

	values := strings.SplitN(filter, ":", 2)

	if len(values) == 1 {
		return nil, fmt.Errorf("error from-v filter: no contain version value")
	}

	// TODO Сделать проверку по формату версии

	return &VersionFromFilter{
		values[1],
	}, nil

}

type LatestVersionFilter struct {
	filter *regexp.Regexp
}

func (m *LatestVersionFilter) Filter(source []*ProjectVersionInfo) (latest []*ProjectVersionInfo) {

	if m.filter != nil {
		var result []*ProjectVersionInfo
		eachVersion(source, func(versionInfo *ProjectVersionInfo) {
			if m.filter.MatchString(versionInfo.Name) {
				result = append(result, versionInfo)
			}
		})
		source = result
	}

	if len(source) == 0 {
		return
	}

	sort.Slice(source, func(i, j int) bool {
		return compareVersion(source[i].Name, source[j].Name) > 0
	})

	latest = append(latest, source[0])

	return
}

func compareVersion(v1, v2 string) int {

	compV1 := strings.Split(v1, ".")
	compV2 := strings.Split(v2, ".")

	maxLen := len(compV1)

	if len(compV1) < len(compV2) {
		maxLen = len(compV2)
	}

	for i := len(compV1); i < maxLen; i++ {
		compV1 = append(compV1, "0")
	}
	for i := len(compV2); i < maxLen; i++ {
		compV2 = append(compV2, "0")
	}

	for i := 0; i < maxLen; i++ {

		if dry.StringToInt(compV1[i]) > dry.StringToInt(compV2[i]) {
			return 1
		}

		if dry.StringToInt(compV1[i]) < dry.StringToInt(compV2[i]) {
			return -1
		}
	}

	return 0
}

func newLatestVersionFilter(filter string) (*LatestVersionFilter, error) {

	var filterRe *regexp.Regexp

	values := strings.SplitN(filter, ":", 2)

	if len(values) == 2 {
		var err error
		filterRe, err = regexp.Compile(values[1])

		if err != nil {
			return nil, fmt.Errorf("error latest filter: %s", err.Error())
		}
	}

	return &LatestVersionFilter{
		filterRe,
	}, nil

}
