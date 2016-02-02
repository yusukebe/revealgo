package revealgo

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestDetectContentType(t *testing.T) {
	actual := detectContentType("/css/moon.css", []byte("h1 {};"))
	expected := "text/css"
	if actual != expected {
		t.Errorf("got %v\n want %v", actual, expected)
	}
	actual = detectContentType("/js/hello.js", []byte("console('Hello')"))
	expected = "application/javascript"
	if actual != expected {
		t.Errorf("got %v\n want %v", actual, expected)
	}
	actual = detectContentType("/readme.txt", []byte("Hello"))
	expected = "text/plain; charset=utf-8"
	if actual != expected {
		t.Errorf("got %v\n want %v", actual, expected)
	}
}

func TestRootHandler(t *testing.T) {
	param := ServerParam{
		Path:       "slide.md",
		Theme:      "beige.css",
		Transition: "fade",
	}
	ts := httptest.NewServer(&rootHandler{param: param})
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("unexpected\n")
	}
	if res.StatusCode != 200 {
		t.Errorf("server status error\n")
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	s := buf.String()

	match := "revealjs/css/theme/beige.css"
	r := regexp.MustCompile(match)
	if r.MatchString(s) == false {
		t.Errorf("content do not match %v\n", match)
	}

	match = `data-markdown="slide.md"`
	r = regexp.MustCompile(match)
	if r.MatchString(s) == false {
		t.Errorf("content do not match %v\n", match)
	}

	match = `|| 'zoom',`
	r = regexp.MustCompile(match)
	if r.MatchString(s) == false {
		t.Errorf("content do not match %v\n", match)
	}
}

func TestAssetHandler(t *testing.T) {
	ts := httptest.NewServer(&assetHandler{assetPath: "assets"})
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
