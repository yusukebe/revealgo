package revealgo

import (
	"testing"
)

func TestAddExtention(t *testing.T) {
	actual := addExtention("black", "css")
	expected := "black.css"
	if actual != expected {
		t.Errorf("got %v\n want %v", actual, expected)
	}
	actual = addExtention("white.css", "css")
	expected = "white.css"
	if actual != expected {
		t.Errorf("got %v\n want %v", actual, expected)
	}
}

func TestDetectContentType(t *testing.T) {
	scenarios := []struct {
		path     string
		expected string
	}{
		{path: "css/moon.css", expected: "text/css"},
		{path: "js/hello.js", expected: "application/javascript"},
		{path: "testdata/markdown.svg", expected: "image/svg+xml"},
		{path: "readme.txt", expected: ""},
	}

	for _, s := range scenarios {
		actual := detectContentType(s.path)
		if actual != s.expected {
			t.Errorf("got %v\n want %v", actual, s.expected)
		}
	}
}
