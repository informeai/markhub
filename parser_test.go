package markhub

import (
	"fmt"
	"io/ioutil"
	"strings"
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
	b, err := ioutil.ReadFile("./test/tester.md")
	if err != nil {
		t.Errorf("TestReadFile(): got -> %v, want: []byte ", err)
	}
	fmt.Println(string(b))

}

// go test -v -run ^TestReadString
func TestReadString(t *testing.T) {
	reader := strings.NewReader("# title 1")
	if reader.Len() == 0 {
		t.Errorf("TestReadString(): got -> %v, want: lenght > 0", reader)
	}
}
