package cyoa

import (
	"encoding/json"
	"io"
)

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

// JsonStory to parse JSON into Story map type.
func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}
