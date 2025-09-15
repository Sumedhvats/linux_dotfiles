package dynamicstory

import (
	"encoding/json"
	"io"
)

var defaultHandlerTemp=`<!DOCTYPE html>
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
    <ul>
        {{range .Options}}
    <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
    </ul>
    {{end}}
</body>
</html>
`
func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
return nil,err	
	}
return story	,nil
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
