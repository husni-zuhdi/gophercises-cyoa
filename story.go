package cyoa

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
