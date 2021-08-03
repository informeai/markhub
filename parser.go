package markhub

import (
	"encoding/json"
	"errors"
	"fmt"
	css "github.com/informeai/markhub/css"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	//address api the github markdown.
	API      string = "https://api.github.com/markdown"
	templCss string = css.TemplCss
)

//MarkHub is struct base for manipulate and parse,
//text the files markdown the github.
type MarkHub struct {
	content string
	html    string
}

//jsonBody is struct for marshall text.
type jsonBody struct {
	Text string `json:"text"`
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
	m.content = string(b)
	return nil
}

//readString read content of string text.
func (m *MarkHub) readString(text string) error {
	if len(text) == 0 {
		return errors.New("empty text is not allowed.")
	}
	m.content = text
	return nil
}

//getApi is responsable of send text markdown for api github,
//and return the html.
func (m *MarkHub) getApi() (string, error) {
	body := jsonBody{Text: m.content}
	textJson, err := json.Marshal(&body)
	if err != nil {
		return "", err
	}
	reader := strings.NewReader(string(textJson))
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, API, reader)
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

//ParseString is responsible for parsing the markdown,
//text using the github api.
func (m *MarkHub) ParseString(text string) error {
	err := m.readString(text)
	if err != nil {
		return err
	}
	t, err := m.getApi()
	if err != nil {
		return err
	}
	template := fmt.Sprintf("<html><head><style>%v</style></head><body>%v</body></html>", templCss, t)
	m.html = template
	return nil
}

//ParseFile is responsible for parsing the markdown,
//file using the github api.
func (m *MarkHub) ParseFile(filePath string) error {
	err := m.readFile(filePath)
	if err != nil {
		return err
	}
	t, err := m.getApi()
	if err != nil {
		return err
	}

	template := fmt.Sprintf("<html><head><style>%v</style></head><body>%v</body></html>", templCss, t)
	m.html = template
	return nil
}

//HTML is return html parsed from github api.
func (m *MarkHub) HTML() string {
	return m.html
}

//Serve is render the html at the given address.
func (m *MarkHub) Serve(address string) error {
	serveHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, m.html)
	}
	http.HandleFunc("/", serveHandler)
	fmt.Println("Serving markdown...")
	err := http.ListenAndServe(address, nil)
	if err != nil {
		return err
	}
	return nil
}
