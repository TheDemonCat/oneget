package unpacker

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

func TestUnpackTarGz(t *testing.T) {
	tempDir, err := getTempDir()
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	currentDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	path := filepath.Join(currentDir, "test", "fixtures","linux", "client")
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileName := file.Name()
		Extract(filepath.Join("test", "fixtures","linux", "client", fileName), tempDir)

		fileGZ := fileName[:len(fileName) - len(filepath.Ext(fileName))]
		dirName := fileGZ[:len(fileGZ) - len(filepath.Ext(fileGZ))]

		files, err := ioutil.ReadDir(filepath.Join(tempDir, dirName))
		if err != nil {
			log.Fatal(err)
		}

		assert.Equal(t, len(files), 4, "Распаковка успешно завершена")
	}
}

func TestNameAliases_8_3_16(t *testing.T) {
	distrNames := []string{
		"1c-enterprise83-client_8.3.16-1876_amd64.deb",
		"1c-enterprise83-client-nls_8.3.16-1876_amd64.deb",
		"1c-enterprise83-thin-client_8.3.16-1876_amd64.deb",
		"1c-enterprise83-thin-client-nls_8.3.16-1876_amd64.deb",
		"1c-enterprise83-common_8.3.16-1876_amd64.deb",
		"1c-enterprise83-common-nls_8.3.16-1876_amd64.deb",
		"1c-enterprise83-crs_8.3.16-1876_amd64.deb",
		"1c-enterprise83-server_8.3.16-1876_amd64.deb",
		"1c-enterprise83-server-nls_8.3.16-1876_amd64.deb",
		"1c-enterprise83-ws_8.3.16-1876_amd64.deb",
		"1c-enterprise83-ws-nls_8.3.16-1876_amd64.deb",
	}
	checkAliases(t, distrNames)
}

func TestNameAliases_8_3_18_1334(t *testing.T) {
	distrNames := [] string{
		"1c-enterprise-8.3.18.1334-client_8.3.18-1334_amd64.deb",
		"1c-enterprise-8.3.18.1334-client-nls_8.3.18-1334_amd64.deb",
		"1c-enterprise-8.3.18.1334-thin-client_8.3.18-1334_amd64.deb",
		"1c-enterprise-8.3.18.1334-thin-client-nls_8.3.18-1334_amd64.deb",
	}

	checkAliases(t, distrNames)
}

func checkAliases(t *testing.T, distrNames []string) {

	for _, distrName := range distrNames {
		result := GetAliasesDistrib(distrName)
		fmt.Println(result)
		regexp := regexp.MustCompile(`^(.*)-([\d\.\d\.\d*\.\d*]*-[\d]*).(.*)$`)
		find := regexp.ReplaceAllString(result, `$1-VERSION.$3`)
		assert.Contains(t, getExpectedName(), find)
	}
}

func getExpectedName() map[string]string {
	return map[string]string{
		"client-VERSION.deb": 			"", //	"1c-enterprise83-client_8.3.16-1876_amd64.deb"
		"client-nls-VERSION.deb": 		"", //	"1c-enterprise83-client-nls_8.3.16-1876_amd64.deb"
		"common-VERSION.deb": 			"", //	"1c-enterprise83-common_8.3.16-1876_amd64.deb"
		"crs-VERSION.deb": 				"", //	"1c-enterprise83-crs_8.3.16-1876_amd64.deb"
		"common-nls-VERSION.deb": 		"", //	"1c-enterprise83-common-nls_8.3.16-1876_amd64.deb"
		"server-VERSION.deb": 			"", //	"1c-enterprise83-server_8.3.16-1876_amd64.deb"
		"server-nls-VERSION.deb":		"", //  "1c-enterprise83-server-nls_8.3.16-1876_amd64.deb"
		"thin-client-nls-VERSION.deb": 	"", //	"1c-enterprise83-thin-client-nls_8.3.16-1876_amd64.deb"
		"thin-client-VERSION.deb": 		"", //	"1c-enterprise83-thin-client_8.3.16-1876_amd64.deb"
		"ws-nls-VERSION.deb": 			"", //	"1c-enterprise83-ws-nls_8.3.16-1876_amd64.deb"
		"ws-VERSION.deb": 				"", //	"1c-enterprise83-ws_8.3.16-1876_amd64.deb"
	}
}

func getTempDir() (string, error) {
	dir, err := ioutil.TempDir("", "oneget")
	if err != nil {
		log.Fatal(err)
	}
	return dir, nil
}