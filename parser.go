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
	if string(m.content[len(m.content)-1]) == "\n" {
		m.content = m.content[:len(m.content)-1]
	}
	return strings.Split(string(m.content), "\n")
}

//verifyChars is responsible for checking
// the characters at the beginning of the line.
func (m *MarkHub) verifyChars() bool {
	var ok bool = false
	var blockQuotes []string
	var isBlockQuote bool = false
	lines := m.splitContent()
	for _, l := range lines {
		fmt.Println(l)
		switch string(l[0]) {
		case "#":
			if isBlockQuote == true && len(blockQuotes) > 0 {
				ok = m.parseBlockQuotes(blockQuotes)
				isBlockQuote = false

			}
			ok = m.parseTitles(l)
		case ">":
			blockQuotes = append(blockQuotes, l)
			isBlockQuote = true

		case "*":
			fmt.Println("*")
		case "-":
			fmt.Println("-")
		default:
			if isBlockQuote == true && len(blockQuotes) > 0 {
				ok = m.parseBlockQuotes(blockQuotes)
				isBlockQuote = false

			}
			continue

		}

	}

	return ok
}

//parseTitles is responsible for parsing tags h.
func (m *MarkHub) parseTitles(line string) bool {
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
	return verifyHash
}

//parseBlockQuotes is responsible for parsing tags blockquotes.
func (m *MarkHub) parseBlockQuotes(lines []string) bool {
	var verify bool = false
	var innerQuotes string = ""
	var end string
	for _, l := range lines {
		if string(l[1]) == " " {
			end = l[1:]
		} else {
			end = fmt.Sprintf(" %v", l[1:])
		}
		if l[len(l)-2:] == "  " {
			innerQuotes += fmt.Sprintf("%v<br>", end)
		} else {
			innerQuotes += end
		}
		verify = true
	}
	m.html += fmt.Sprintf("<blockquotes><p>%v</p></blockquotes>", innerQuotes)
	return verify
}
