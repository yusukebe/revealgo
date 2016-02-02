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
	actual := detectContentType("/css/moon.css", []byte("h1 {};"), )
	expected := "text/css"
	if actual != expected {
		t.Errorf("got %v\n want %v", actual, expected)
	}
	actual = detectContentType("/js/hello.js", []byte("console('Hello')"), )
	expected = "application/javascript"
	if actual != expected {
		t.Errorf("got %v\n want %v", actual, expected)
	}
	actual = detectContentType("/readme.txt", []byte("Hello"), )
	expected = "text/plain; charset=utf-8"
	if actual != expected {
		t.Errorf("got %v\n want %v", actual, expected)
	}
}

