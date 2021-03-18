package downloader

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"
)

func TestNewDownloader(t *testing.T) {
	startDate, err := time.Parse("02.01.2006", "20.01.2020")
	if err != nil {
		t.Error(err)
	}

	conf := Config{
		Login:     "user",
		Password:  "user",
		StartDate: startDate,
	}
	New(conf)
}

func TestLoginIncorrect(t *testing.T) {
	conf := Config{
		Login:     "user",
		Password:  "user",
		StartDate: time.Now(),
	}
	downldr := New(conf)
	_, err := downldr.Get()

	if !(strings.Contains(err.Error(), "Incorrect login or password") ||
		strings.Contains(err.Error(), "Too many failed attempts")) {
		t.Error("Test bad login :(")
	}
}

func TestGetPlatform_8_3_18_1334_linux(t *testing.T) {
	nicks := make(map[string]bool, 0)
	nicks["platform83"] = true

	dir, err := ioutil.TempDir("", "oneget")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	conf := Config{
		Login:         "user",
		Password:      "user",
		Nicks:         nicks,
		BasePath:      dir,
		VersionFilter: "8.3.18.1334",
		DistribFilter: "deb64.tar.gz$",
	}

	files := GetPlatform(t, conf)
	if len(files) != 1 {
		t.Errorf("files must be 1")
	}
}

func TestGetPlatform_8_3_18_1334_windows(t *testing.T) {
	nicks := make(map[string]bool, 0)
	nicks["platform83"] = true

	dir, err := ioutil.TempDir("", "oneget")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	conf := Config{
		Login:         "user",
		Password:      "user",
		Nicks:         nicks,
		BasePath:      dir,
		VersionFilter: "8.3.18.1334",
		DistribFilter: "windows64",
	}
	files := GetPlatform(t, conf)
	if len(files) != 1 {
		t.Errorf("files must be 1")
	}
}

func GetPlatform(t *testing.T, conf Config) []os.FileInfo {

	handler := getHandler()

	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	releasesURL_bak := releasesURL
	releasesURL = ts.URL + "/releases"

	loginURL_bak := loginURL
	loginURL = ts.URL + "/login"

	defer func() { releasesURL = releasesURL_bak; loginURL = loginURL_bak }()

	downldr := New(conf)
	files, err := downldr.Get()
	if err != nil {
		t.Error(err)
	}
	return files
}

func getHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/login/rest/public/ticket/get" {
			fmt.Fprint(w, "{\"ticket\": \"Hello\"}")
		} else if r.RequestURI == "/login/ticket/auth?token=Hello" {
			fmt.Fprintln(w, "<a href=\"/project/Platform80\"/>")
			fmt.Fprintln(w, "<a href=\"/project/Platform81\"/>")
			fmt.Fprintln(w, "<a href=\"/project/Platform83\"/>")
		} else if strings.Contains(r.RequestURI, "project/Platform83") {
			fmt.Fprintln(w, "<table id=\"versionsTable\" class=\"customTable table-hover\">")
			fmt.Fprintln(w,
				"<tr>",
				"<td class=\"versionColumn\"><a href=\"/version_files?nick=Platform83&ver=8.3.18.1334\"/></td>",
				"<td class=\"dateColumn\">27.04.17</td>",
				"</tr>")
			fmt.Fprintln(w,
				"<tr>",
				"<td class=\"versionColumn\"><a href=\"/version_files?nick=Platform83&ver=8.3.17.1851\"/></td>",
				"<td class=\"dateColumn\">01.09.16</td>",
				"</tr>")
			fmt.Fprintln(w,
				"<tr>",
				"<td class=\"versionColumn\"><a href=\"/version_files?nick=Platform83&ver=8.3.16.1814\"/></td>",
				"<td class=\"dateColumn\">29.12.15</td>",
				"</tr>")
			fmt.Fprintln(w, "</table>")
		} else if r.URL.Path == "/releases/version_files" {
			query, err := url.ParseQuery(r.URL.RawQuery)
			if err != nil {
				log.Fatal(err.Error())
			}

			nick := query.Get("nick")
			ver := query.Get("ver")
			ver = strings.Replace(ver, ".", "_", -1)

			fmt.Fprintf(w, "<a href=\"/version_file?%s&path=%s\\%s\\Readme.txt\"/>\n",
				r.URL.RawQuery, nick, ver)
			fmt.Fprintf(w, "<a href=\"/version_file?%s&path=%s\\%s\\setuptc_%s.rar\"/>\n",
				r.URL.RawQuery, nick, ver, ver)
			fmt.Fprintf(w, "<a href=\"/version_file?%s&path=%s\\%s\\thin.client_%s.deb32.tar.gz\"/>\n",
				r.URL.RawQuery, nick, ver, ver)
			fmt.Fprintf(w, "<a href=\"/version_file?%s&path=%s\\%s\\thin.client_%s.deb64.tar.gz\"/>\n",
				r.URL.RawQuery, nick, ver, ver)
			fmt.Fprintf(w, "<a href=\"/version_file?%s&path=%s\\%s\\windows64_%s.rar\"/>\n",
				r.URL.RawQuery, nick, ver, ver)
			fmt.Fprintf(w, "<a href=\"/version_file?%s&path=%s\\%s\\windows_%s.rar\"/>\n",
				r.URL.RawQuery, nick, ver, ver)
		} else if strings.HasSuffix(r.RequestURI, ".exe") {
			fmt.Fprintf(w, "<a href=\"%s/public/file/get/id\"/>", releasesURL)
		} else if strings.HasSuffix(r.RequestURI, ".rar") {
			fmt.Fprintf(w, "<a href=\"%s/public/file/get/id\"/>", releasesURL)
		} else if strings.HasSuffix(r.RequestURI, ".gz") {
			fmt.Fprintf(w, "<a href=\"%s/public/file/get/id\"/>", releasesURL)
		} else if strings.HasSuffix(r.RequestURI, ".txt") {
			fmt.Fprintln(w, "File received!")
		} else if strings.Contains(r.RequestURI, "/public/file/get/id") {
			fmt.Fprintln(w, "Distribution received!")
		} else {
			println("debug")
		}
	}
}
