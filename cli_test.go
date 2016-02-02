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
