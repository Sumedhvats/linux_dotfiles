package dynamicstory

import (
	"encoding/json"
	"io"
	"net/http"
	"text/template"
)

var defaultHandlerTemp = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Choose your own adventure</title>
</head>
<body>
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
        <p>{{.}}</p>
    {{end}}

    <ul>
        {{range .Options}}
            <li><a href="/{{.Arc}}">{{.Text}}</a></li>
        {{end}}
    </ul>
</body>
</html>`

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path:=r.URL.Path
	if path==""||path=="/"{
		path="/intro"
	}
	path=path[1:]
	if Chapter,ok:=h.s[path]; 
	tpl := template.Must(template.New("").Parse(defaultHandlerTemp))
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}
}
func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

type Story map[string]Chapter
type Chapter struct {
	Title      string    `json:"title"`
	Paragraphs []string  `json:"story"`
	Options    []Options `json:"options"`
}
type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
