package markhub

import (
	"log"
	// "strings"
	"testing"
)

// go test -v -run ^TestNewMarkHub
func TestNewMarkHub(t *testing.T) {
	m := NewMarkHub()
	if m == nil {
		t.Errorf("TestNewMarkHub(): got -> nil, want: MarkHub{}")
	}
}

// go test -v -run ^TestReadFile
func TestReadFile(t *testing.T) {
	m := NewMarkHub()
	err := m.readFile("./test/tester.md")
	if err != nil {
		t.Errorf("TestReadFile(): got -> %v, want: nil ", err)
	}
}

// go test -v -run ^TestReadString
func TestReadString(t *testing.T) {
	m := NewMarkHub()
	err := m.readString("# title 1")
	if err != nil {
		t.Errorf("TestReadString(): got -> %v, want: nil", err)
	}
}

//go test -v -run ^TestSplitContent
func TestSplitContent(t *testing.T) {
	m := NewMarkHub()
	m.content = []byte("# title1\n# title 2\n\nnew paragraph.")
	sliceOfStrings := m.splitContent()
	if sliceOfStrings == nil {
		t.Errorf("TestSplitContent(): got -> %v, want: []string", sliceOfStrings)
	}
	log.Println(sliceOfStrings)
}

//go test -v -run ^TestVerifyChars
func TestVerifyChars(t *testing.T) {
	m := NewMarkHub()
	b := m.verifyChars("# title 1")
	if b == false {
		t.Errorf("TestSplitContent(): got -> %v, want: true", b)
	}
}

//go test -v -run ^TestParseTitles
func TestParseTitles(t *testing.T) {
	m := NewMarkHub()
	m.parseTitles("###### titles 3 ")
	if len(m.html) == 0 {
		t.Errorf("TestParseTitles(): got -> %v , want: > 0", len(m.html))
	}
	log.Println(m.html)
}
