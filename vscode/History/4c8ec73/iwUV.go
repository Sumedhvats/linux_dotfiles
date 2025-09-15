package dynamicstory

import (
	"encoding/json"
	"io"
)

func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}

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
