package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

var defaultHandlerTemplate = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose your own adventure</title>
  </head>
  <body>
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
    <p>{{.}}</p>
    {{end}}

    <ul>
      {{range .Options}}
      <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
      {{end}}
    </ul>
  </body>
</html>
`

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("").Parse(defaultHandlerTemplate))
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		log.Fatal(err)
	}
}

type Story map[string]Chapter

func JsonStory(r io.Reader) (story Story, err error) {
	d := json.NewDecoder(r)

	if err = d.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
