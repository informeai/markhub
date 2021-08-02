package markhub

import (
	"errors"
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
//if there are key characters for the markdown
//at the beginning of the line.
func (m *MarkHub) verifyChars(line string) bool {

	return true
}
