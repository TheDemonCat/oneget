package downloader

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

func TestMatchFilter_Match(t *testing.T) {

	type args struct {
		filters string
		project string
	}

	type want struct {
		source string
		want   bool
	}

	tests := []struct {
		filters string
		project string
		source  string
		want    bool
	}{
		{
			"server.x64.win",
			Platform83Project,
			"Cервер 1С:Предприятия (64-bit) для Windows",
			true,
		}, {
			"server.x64.win",
			Platform83Project,
			`Тонкий клиент 1С:Предприятие (64-bit) для Windows
Тонкий клиент 1С:Предприятия для Windows
Тонкий клиент 1С:Предприятие (64-bit) для Windows
Тонкий клиент 1С:Предприятия для DEB-based Linux-систем
Тонкий клиент 1С:Предприятия для RPM-based Linux-систем
Тонкий клиент 1С:Предприятия (64-bit) для DEB-based Linux-систем
Тонкий клиент 1С:Предприятия (64-bit) для RPM-based Linux-систем
Тонкий клиент 1С:Предприятия для macOS
Технологическая платформа 1С:Предприятия для Windows
Технологическая платформа 1С:Предприятия (64-bit) для Windows
Cервер 1С:Предприятия для DEB-based Linux-систем
Клиент 1С:Предприятия для macOS
Cервер 1С:Предприятия для RPM-based Linux-систем
Cервер 1С:Предприятия (64-bit) для DEB-based Linux-систем
Cервер 1С:Предприятия (64-bit) для RPM-based Linux-систем
Cервер 1С:Предприятия для Windows
`,
			false,
		}, {
			"server.x32.win",
			Platform83Project,
			"Cервер 1С:Предприятия для Windows",
			true,
		}, {
			"server.x32.win",
			Platform83Project,
			`Тонкий клиент 1С:Предприятие (64-bit) для Windows
Тонкий клиент 1С:Предприятия для Windows
Тонкий клиент 1С:Предприятие (64-bit) для Windows
Тонкий клиент 1С:Предприятия для DEB-based Linux-систем
Тонкий клиент 1С:Предприятия для RPM-based Linux-систем
Тонкий клиент 1С:Предприятия (64-bit) для DEB-based Linux-систем
Тонкий клиент 1С:Предприятия (64-bit) для RPM-based Linux-систем
Тонкий клиент 1С:Предприятия для macOS
Технологическая платформа 1С:Предприятия для Windows
Технологическая платформа 1С:Предприятия (64-bit) для Windows
Cервер 1С:Предприятия для DEB-based Linux-систем
Клиент 1С:Предприятия для macOS
Cервер 1С:Предприятия для RPM-based Linux-систем
Cервер 1С:Предприятия (64-bit) для DEB-based Linux-систем
Cервер 1С:Предприятия (64-bit) для RPM-based Linux-систем
`,
			false,
		}, {
			"server.win",
			Platform83Project,
			"Cервер 1С:Предприятия для Windows",
			true,
		}, {
			"server.win",
			Platform83Project,
			`Тонкий клиент 1С:Предприятие (64-bit) для Windows
Тонкий клиент 1С:Предприятия для Windows
Тонкий клиент 1С:Предприятие (64-bit) для Windows
Тонкий клиент 1С:Предприятия для DEB-based Linux-систем
Тонкий клиент 1С:Предприятия для RPM-based Linux-систем
Тонкий клиент 1С:Предприятия (64-bit) для DEB-based Linux-систем
Тонкий клиент 1С:Предприятия (64-bit) для RPM-based Linux-систем
Тонкий клиент 1С:Предприятия для macOS
Технологическая платформа 1С:Предприятия для Windows
Технологическая платформа 1С:Предприятия (64-bit) для Windows
Cервер 1С:Предприятия для DEB-based Linux-систем
Клиент 1С:Предприятия для macOS
Cервер 1С:Предприятия для RPM-based Linux-систем
Cервер 1С:Предприятия (64-bit) для DEB-based Linux-систем
Cервер 1С:Предприятия (64-bit) для RPM-based Linux-систем
`,
			false,
		}, {
			"thin.win",
			Platform83Project,
			"Тонкий клиент 1С:Предприятия для Windows",
			true,
		}, {
			"thin.win",
			Platform83Project,
			`Тонкий клиент 1С:Предприятие (64-bit) для Windows
Тонкий клиент 1С:Предприятие (64-bit) для Windows
Тонкий клиент 1С:Предприятия для DEB-based Linux-систем
Тонкий клиент 1С:Предприятия для RPM-based Linux-систем
Тонкий клиент 1С:Предприятия (64-bit) для DEB-based Linux-систем
Тонкий клиент 1С:Предприятия (64-bit) для RPM-based Linux-систем
Тонкий клиент 1С:Предприятия для macOS
Технологическая платформа 1С:Предприятия для Windows
Технологическая платформа 1С:Предприятия (64-bit) для Windows
Cервер 1С:Предприятия для DEB-based Linux-систем
Клиент 1С:Предприятия для macOS
Cервер 1С:Предприятия для RPM-based Linux-систем
Cервер 1С:Предприятия (64-bit) для DEB-based Linux-систем
Cервер 1С:Предприятия (64-bit) для RPM-based Linux-систем
Cервер 1С:Предприятия для Windows
`,
			false,
		}, {
			"win",
			Platform83Project,
			`Технологическая платформа 1С:Предприятия для Windows`,
			true,
		}, {
			"win.x64",
			Platform83Project,
			`Технологическая платформа 1С:Предприятия (64-bit) для Windows`,
			true,
		}, {
			"win",
			Platform83Project,
			`Тонкий клиент 1С:Предприятие (64-bit) для Windows
Тонкий клиент 1С:Предприятие (64-bit) для Windows
Тонкий клиент 1С:Предприятия для DEB-based Linux-систем
Тонкий клиент 1С:Предприятия для RPM-based Linux-систем
Тонкий клиент 1С:Предприятия (64-bit) для DEB-based Linux-систем
Тонкий клиент 1С:Предприятия (64-bit) для RPM-based Linux-систем
Тонкий клиент 1С:Предприятия для macOS
Тонкий клиент 1С:Предприятия для Windows
Cервер 1С:Предприятия для Windows
Технологическая платформа 1С:Предприятия (64-bit) для Windows
Cервер 1С:Предприятия для DEB-based Linux-систем
Клиент 1С:Предприятия для macOS
Cервер 1С:Предприятия для RPM-based Linux-систем
Cервер 1С:Предприятия (64-bit) для DEB-based Linux-систем
Cервер 1С:Предприятия (64-bit) для RPM-based Linux-систем
`,
			false,
		}, {
			"win.full",
			Platform83Project,
			`Технологическая платформа 1С:Предприятия для Windows`,
			true,
		}, {
			"win.full.x64",
			Platform83Project,
			`Технологическая платформа 1С:Предприятия (64-bit) для Windows`,
			true,
		}, {
			"win.full",
			Platform83Project,
			`Тонкий клиент 1С:Предприятие (64-bit) для Windows
Тонкий клиент 1С:Предприятие (64-bit) для Windows
Тонкий клиент 1С:Предприятия для DEB-based Linux-систем
Тонкий клиент 1С:Предприятия для RPM-based Linux-систем
Тонкий клиент 1С:Предприятия (64-bit) для DEB-based Linux-систем
Тонкий клиент 1С:Предприятия (64-bit) для RPM-based Linux-систем
Тонкий клиент 1С:Предприятия для macOS
Технологическая платформа 1С:Предприятия (64-bit) для Windows
Cервер 1С:Предприятия для DEB-based Linux-систем
Клиент 1С:Предприятия для macOS
Cервер 1С:Предприятия для RPM-based Linux-систем
Cервер 1С:Предприятия (64-bit) для DEB-based Linux-систем
Cервер 1С:Предприятия (64-bit) для RPM-based Linux-систем
`,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.filters, func(t *testing.T) {
			m, err := NewFileFilter(tt.project, tt.filters)

			if err != nil {
				t.Fatal(err)
			}

			cases := strings.Split(tt.source, "\n")
			for _, testCase := range cases {
				if got := m.MatchString(testCase); got != tt.want {
					t.Errorf("Match() = %v, want %v, source: %s", got, tt.want, testCase)
				}
			}

		})
	}
}

func TestLatestVersionFilter_Filter(t *testing.T) {

	tests := []struct {
		name       string
		filter     *regexp.Regexp
		versions   []*ProjectVersionInfo
		wantLatest []*ProjectVersionInfo
	}{
		{
			"simple",
			regexp.MustCompile("8.3.16"),
			[]*ProjectVersionInfo{
				{
					Name: "8.3.16.1324",
				}, {
					Name: "8.3.16.965",
				}, {
					Name: "8.3.17.1324",
				},
			},
			[]*ProjectVersionInfo{
				{
					Name: "8.3.16.1324",
				},
			},
		}, {
			"no filter",
			nil,
			[]*ProjectVersionInfo{
				{
					Name: "8.3.16.1324",
				}, {
					Name: "8.3.16.965",
				}, {
					Name: "8.3.17.1589",
				},
			},
			[]*ProjectVersionInfo{
				{
					Name: "8.3.17.1589",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &LatestVersionFilter{
				filter: tt.filter,
			}
			if gotLatest := m.Filter(tt.versions); !reflect.DeepEqual(gotLatest, tt.wantLatest) {
				t.Errorf("Filter() = %v, want %v", gotLatest[0], tt.wantLatest[0])
			}
		})
	}
}

func Test_compareVersion(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		v1   string
		v2   string
		want int
	}{
		{"8.3.10.1877", "8.3.9.2016", 1},
		{"8.3.10.1877", "", 1},
		{"08.03.010.01877", "8.3.9.2016", 1},
		{"8.3.9.2016", "8.3.10.1877", -1},
		{"8.3", "8.3.9.2016", -1},
		{"08.03.09.0002016", "8.3.9.2016", 0},
		// two values
		{"1.2", "1.1", 1},
		{"1.10", "1.9", 1},
		{"1.10.1", "1.10", 1},
		{"1.2", "", 1},
		{"1.1", "1.2", -1},
		{"1.5", "1.5.0", 0},
		// edt format
		{"2020.2", "2020.1", 1},
		{"2020.2.1", "2020.2.0", 1},
		{"2020.3", "1.16.0.363", 1},
		{"2020.6", "2021.1", -1},
		{"2020.6.2", "2020.6.3", -1},
		{"2020.2.0", "2020.2", 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s vs %s", tt.v1, tt.v2), func(t *testing.T) {
			if got := compareVersion(tt.v1, tt.v2); got != tt.want {
				t.Errorf("compareVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersionFromFilter_Filter(t *testing.T) {

	tests := []struct {
		name       string
		filter     string
		versions   []string
		wantResult []string
	}{
		{"no results", "8.3.18", []string{"8.3.16.1564"}, []string{}},
		{"all results", "8.3.16", []string{"8.3.16.1564", "8.3.16.965"}, []string{"8.3.16.1564", "8.3.16.965"}},
		{"only 8.3", "8.3", []string{"8.3.16.1564", "8.3.16.965", "8.2.8"}, []string{"8.3.16.1564", "8.3.16.965"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			m := &VersionFromFilter{
				version: tt.filter,
			}

			var versions = []*ProjectVersionInfo{}
			for _, version := range tt.versions {
				versions = append(versions, &ProjectVersionInfo{Name: version})
			}

			var wantResult = []*ProjectVersionInfo{}
			for _, version := range tt.wantResult {
				wantResult = append(wantResult, &ProjectVersionInfo{Name: version})
			}

			if gotResult := m.Filter(versions); !(len(gotResult) == 0 && len(wantResult) == 0) && !reflect.DeepEqual(gotResult, wantResult) {
				t.Errorf("Filter() = %v, want %v", gotResult, wantResult)
			}
		})
	}
}
