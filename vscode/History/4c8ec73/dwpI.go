package dynamicstory
Type story Chapter[]
type Chapter struct {
	Title      string    `json:"title"`
	Paragraphs []string  `json:"story"`
	Options    []Options `json:"options"`
}
type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
