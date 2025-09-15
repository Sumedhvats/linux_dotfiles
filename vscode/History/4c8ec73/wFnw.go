package dynamicstory

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []struct {
	} `json:"options"`
}
type Options struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`

}