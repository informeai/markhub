package markhub

import (
	// "log"
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

// go test -v -run ^TestGetApi
func TestGetApi(t *testing.T) {
	m := NewMarkHub()
	m.content = "## teste 1\n>blockotes1\n>blockotes2\n\nmade with :heart:"
	_, err := m.getApi()
	if err != nil {
		t.Errorf("TestGetApi(): got -> %v, want: nil", err)
	}
}

// go test -v -run ^TestParseString
func TestParseString(t *testing.T) {
	m := MarkHub{}
	err := m.ParseString("# h1 title\n:+1:")
	if err != nil {
		t.Errorf("TestParseString(): got -> %v, want: nil", err)
	}

}

// go test -v -run ^TestParseFile
func TestParseFile(t *testing.T) {
	m := MarkHub{}
	err := m.ParseFile("./test/tester.md")
	if err != nil {
		t.Errorf("TestParseString(): got -> %v, want: nil", err)
	}

}

// go test -v -run ^TestServe
func TestServe(t *testing.T) {
	m := MarkHub{}
	err := m.ParseFile("./test/tester.md")
	if err != nil {
		t.Errorf("TestServe(): got -> %v, want: nil", err)
	}
	err = m.Serve(":4000")
	if err != nil {
		t.Errorf("TestServe(): got -> %v, want: nil", err)
	}

}
