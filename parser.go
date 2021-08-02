package markhub

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	startChars string = "#>*-"
)

//MarkHub is struct base for manipulate and parse,
//text the files markdown.
type MarkHub struct {
	content []byte
	html    string
}

//NewMarkHub return of instance the MarkHub
func NewMarkHub() *MarkHub {
	return &MarkHub{}
}

//readFile read content of file.
func (m *MarkHub) readFile(filePath string) error {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	m.content = b
	return nil
}

//readString read content of string text.
func (m *MarkHub) readString(text string) error {
	if len(text) == 0 {
		return errors.New("empty text is not allowed.")
	}
	m.content = []byte(text)
	return nil
}

//splitContent return slice of strings separeted by lines.
func (m *MarkHub) splitContent() []string {
	return strings.Split(string(m.content), "\n")
}

//verifyChars is responsible for checking
// the characters at the beginning of the line.
func (m *MarkHub) verifyChars(line string) bool {
	switch string(line[0]) {
	case "#":
		fmt.Println("#")
		return true
	case ">":
		fmt.Println(">")
		return true
	case "*":
		fmt.Println("*")
		return true
	case "-":
		fmt.Println("-")
		return true
	}
	return false
}

//parseTitles is responsible for parsing tags h.
func (m *MarkHub) parseTitles(line string) {
	start := strings.Split(line, " ")[0]
	end := strings.Split(line, " ")[1:]
	inner := strings.Join(end, " ")
	var verifyHash bool = false
	var text string = line
	for _, v := range start {
		if strings.Compare(string(v), "#") == 0 {
			verifyHash = true
			continue
		} else {
			verifyHash = false
			break
		}
	}
	if verifyHash && len(start) < 7 {
		lenght := len(start)
		text = fmt.Sprintf("<h%v>%v</h%v>", lenght, inner, lenght)
	}
	m.html += text
}
