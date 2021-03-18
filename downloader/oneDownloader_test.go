package downloader

import (
	"reflect"
	"regexp"
	"testing"
)

func Test_filterReleaseFiles(t *testing.T) {
	type args struct {
		list    []ReleaseFileInfo
		filters []FileFilter
	}
	tests := []struct {
		name             string
		args             args
		wantFilteredList int
	}{
		{
			"Platform83Project with addin_8_3_18_1363",
			args{
				list: []ReleaseFileInfo{
					{
						"Тонкий клиент 1С:Предприятия для DEB-based Linux-систем",
						"/version_file?nick=Platform83&ver=8.3.18.1363&path=Platform%5C8_3_18_1363%5Cthin.client_8_3_18_1363.deb32.tar.gz",
					},
					{
						"Тонкий клиент 1С:Предприятия для RPM-based Linux-систем",
						"/version_file?nick=Platform83&ver=8.3.18.1363&path=Platform%5C8_3_18_1363%5Cthin.client_8_3_18_1363.rpm32.tar.gz",
					}, {
						"Технология внешних компонент",
						"/version_file?nick=Platform83&ver=8.3.18.1363&path=Platform%5C8_3_18_1363%5Caddin_8_3_18_1363.zip",
					},
				},
				filters: []FileFilter{
					regexp.MustCompile(".*deb32.tar.gz"),
					regexp.MustCompile("Технология внешних компонент"),
					NewFileFilterMust(Platform83Project, "thin.deb"),
				},
			},
			2,
		}, {
			"DevelopmentTools10 with Bellsoft JDK",
			args{
				list: []ReleaseFileInfo{
					{
						"Дистрибутив 1C:EDT для ОС Windows 64 бит",
						"/version_file?nick=DevelopmentTools10&ver=2020.6.2&path=DevelopmentTools%5C2020_6_2%5C1c_edt_distr_2020.6.2_8_windows_x86_64.zip",
					},
					{
						"Дистрибутив для оффлайн установки 1C:EDT для ОС Linux 64 бит",
						"/version_file?nick=DevelopmentTools10&ver=2020.6.2&path=DevelopmentTools%5C2020_6_2%5C1c_edt_distr_offline_2020.6.2_8_linux_x86_64.tar.gz",
					}, {
						"Bellsoft JDK Full (64-bit) для Windows",
						"/version_file?nick=DevelopmentTools10&ver=2020.6.2&path=DevelopmentTools%5C2020_6%5Cbellsoft_jdk11.0.9_12_windows_amd64_full.msi",
					}, {
						"Дистрибутив для оффлайн установки 1C:EDT для ОС Windows 64 бит",
						"/version_file?nick=DevelopmentTools10&ver=2020.6.2&path=DevelopmentTools%5C2020_6_2%5C1c_edt_distr_offline_2020.6.2_8_windows_x86_64.zip",
					},
				},
				filters: []FileFilter{
					NewFileFilterMust("DevelopmentTools10", "deb"),
				},
			},
			1,
		}, {
			"DevelopmentTools10 with Bellsoft JDK",
			args{
				list: []ReleaseFileInfo{
					{
						"Дистрибутив для оффлайн установки 1C:EDT для ОС Linux 64 бит",
						"/version_file?nick=DevelopmentTools10&ver=2020.6.2&path=DevelopmentTools%5C2020_6_2%5C1c_edt_distr_offline_2020.6.2_8_linux_x86_64.tar.gz",
					},
					{
						"Дистрибутив 1C:EDT для ОС Linux 64 бит",
						"/version_file?nick=DevelopmentTools10&ver=2020.6.2&path=DevelopmentTools%5C2020_6_2%5C1c_edt_distr_2020.6.2_8_linux_x86_64.tar.gz",
					}, {
						"Bellsoft JDK Full (64-bit) для DEB-based Linux-систем",
						"/version_file?nick=DevelopmentTools10&ver=2020.6.2&path=DevelopmentTools%5C2020_6%5Cbellsoft_jdk11.0.9_12_linux_amd64_full.deb",
					}, {
						"Bellsoft JDK Full (64-bit) для RPM-based Linux-систем",
						"/version_file?nick=DevelopmentTools10&ver=2020.6.2&path=DevelopmentTools%5C2020_6%5Cbellsoft_jdk11.0.9_12_linux_amd64_full.rpm",
					},
				},
				filters: []FileFilter{
					NewFileFilterMust("DevelopmentTools10", "deb"),
				},
			},
			2,
		}, {
			"DevelopmentTools10 only Bellsoft JDK",
			args{
				list: []ReleaseFileInfo{
					{
						"Дистрибутив для оффлайн установки 1C:EDT для ОС Linux 64 бит",
						"/version_file?nick=DevelopmentTools10&ver=2020.6.2&path=DevelopmentTools%5C2020_6_2%5C1c_edt_distr_offline_2020.6.2_8_linux_x86_64.tar.gz",
					},
					{
						"Дистрибутив 1C:EDT для ОС Linux 64 бит",
						"/version_file?nick=DevelopmentTools10&ver=2020.6.2&path=DevelopmentTools%5C2020_6_2%5C1c_edt_distr_2020.6.2_8_linux_x86_64.tar.gz",
					}, {
						"Bellsoft JDK Full (64-bit) для DEB-based Linux-систем",
						"/version_file?nick=DevelopmentTools10&ver=2020.6.2&path=DevelopmentTools%5C2020_6%5Cbellsoft_jdk11.0.9_12_linux_amd64_full.deb",
					}, {
						"Bellsoft JDK Full (64-bit) для RPM-based Linux-систем",
						"/version_file?nick=DevelopmentTools10&ver=2020.6.2&path=DevelopmentTools%5C2020_6%5Cbellsoft_jdk11.0.9_12_linux_amd64_full.rpm",
					},
				},
				filters: []FileFilter{
					NewFileFilterMust("DevelopmentTools10", "deb.jdk.x64"),
				},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFilteredList := filterReleaseFiles(tt.args.list, tt.args.filters); !reflect.DeepEqual(len(gotFilteredList), tt.wantFilteredList) {
				t.Errorf("filterReleaseFiles() = %v, want %v", len(gotFilteredList), tt.wantFilteredList)
			}
		})
	}
}
