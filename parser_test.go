package markhub

import (
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

// go test -v -run ^TestHTML
func TestHTML(t *testing.T) {
	m := MarkHub{}
	err := m.ParseString("# title 1")
	if err != nil {
		t.Errorf("TestHTML(): got -> %v, want: nil", err)
	}
	html := m.HTML()
	if len(html) == 0 {
		t.Errorf("TestHTML(): got -> %v, want: length > 0", len(html))
	}
}

// go test -v -run ^TestServe
func TestServe(t *testing.T) {
	m := MarkHub{}
	err := m.ParseFile("./README.md")
	if err != nil {
		t.Errorf("TestServe(): got -> %v, want: nil", err)
	}
	err = m.Serve(":4000")
	if err != nil {
		t.Errorf("TestServe(): got -> %v, want: nil", err)
	}

}

// go test -v -run ^ReadFile -bench=.
func BenchmarkReadFile(b *testing.B) {
	m := MarkHub{}
	for n := 0; n < b.N; n++ {
		_ = m.readFile("./test/tester.md")
	}
}

// go test -v -run ^ReadString -bench=.
func BenchmarkReadString(b *testing.B) {
	m := MarkHub{}
	for n := 0; n < b.N; n++ {
		_ = m.readString("# title 1")
	}
}

// go test -v -run ^ParseFile -bench=.
func BenchmarkParseFile(b *testing.B) {
	m := MarkHub{}
	for n := 0; n < b.N; n++ {
		_ = m.ParseFile("./test/tester.md")
	}
}

// go test -v -run ^ParseString -bench=.
func BenchmarkParseString(b *testing.B) {
	m := MarkHub{}
	for n := 0; n < b.N; n++ {
		_ = m.ParseString("> blockquote")
	}
}

// go test -v -run ^GetApi -bench=.
func BenchmarkGetApi(b *testing.B) {
	m := MarkHub{}
	m.content = "# title"
	for n := 0; n < b.N; n++ {
		_, _ = m.getApi()
	}
}
