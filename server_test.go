package revealgo

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestContentHandler(t *testing.T) {
	param := ServerParam{
		Path:       "testdata/example.md",
		Theme:      "beige.css",
		Transition: "fade",
	}
	ts := httptest.NewServer(contentHandler(param, http.FileServer(http.Dir("."))))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal("unexpected", err)
	}
	if res.StatusCode != 200 {
		t.Error("server status error")
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	s := buf.String()

	match := "revealjs/css/theme/beige.css"
	r := regexp.MustCompile(match)
	if r.MatchString(s) == false {
		t.Errorf("content do not match %v\n", match)
	}

	match = `data-markdown="testdata/example.md"`
	r = regexp.MustCompile(match)
	if r.MatchString(s) == false {
		t.Errorf("content do not match %v\n", match)
	}

	match = `|| 'zoom',`
	r = regexp.MustCompile(match)
	if r.MatchString(s) == false {
		t.Errorf("content do not match %v\n", match)
	}

	r2, err := http.Get(ts.URL + "/testdata/markdown.svg")
	if err != nil {
		t.Fatal("unexpected", err)
	}
	if r2.StatusCode != 200 {
		t.Errorf("server status error\n")
	}
	if r2.Header.Get("Content-Type") != "image/svg+xml" {
		t.Errorf("content type error\n")
	}
}

func TestAssetHandler(t *testing.T) {
	ts := httptest.NewServer(assetsHandler("/assets/", http.FileServer(http.FS(revealjs))))
	defer ts.Close()

	res, err := http.Get(ts.URL + "/revealjs/js/reveal.js")
	if err != nil {
		t.Errorf("unexpected\n")
	}
	if res.StatusCode != 200 {
		t.Errorf("server status error\n")
	}
	if res.Header.Get("Content-Type") != "application/javascript" {
		t.Errorf("content type error\n")
	}
}
