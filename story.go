package cyoa

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"
)

// Initiate template creation when code is running.
func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}

var tpl *template.Template

// Story map contain a string Intro as key and Chapter struct as value.
type Story map[string]Chapter

// Chapter struct type to store each chapter of story.
type Chapter struct {
	Title     string   `json:"title"`
	Paragraph []string `json:"story"`
	Options   []Option `json:"options"`
}

// Option struct type to store option in a chapter.
type Option struct {
	Text string `json:"text"`
	Chap string `json:"arc"`
}

// Our default Handler Template.
var defaultHandlerTmpl = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <title>Choose Your Own Adventure</title>
    </head>
    <body>
        <h1>{{.Title}}</h1>
        {{range .Paragraph}}
            <p>{{.}}</p>
        {{end}}
        <ul>
        {{range .Options}}
            <li><a href="/{{.Chap}}">{{.Text}}</a></li>
        {{end}}
        </ul>
    </body>
</html>
`

type handler struct {
	s Story
}

func NewHandler(s Story) handler {
	return handler{s}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Check if the path is empty. If empty, set it to intro
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	// Trim the "/..." from path
	path = path[1:]

	if chap, ok := h.s[path]; ok {
		err := tpl.Execute(w, chap)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something ges wrong... I can't feel it", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found :(", http.StatusNotFound)
}

// JsonStory to parse JSON into Story map type.
func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}
