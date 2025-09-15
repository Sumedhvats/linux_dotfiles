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
	    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FCF6FC;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #797;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: underline;
        color: #555;
      }
      a:active,
      a:hover {
        color: #222;
      }
      p {
        text-indent: 1em;
      }
    </style>
</head>
<body>
<section>
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
	path := r.URL.Path
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]
	if chapter, ok := h.s[path]; ok {
		tpl := template.Must(template.New("").Parse(defaultHandlerTemp))
		err := tpl.Execute(w, chapter)
		if err != nil {
			panic(err)
		}

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
