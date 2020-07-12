package cyoa

import (
	"encoding/json"
	"io"
)

type Story map[string]Chapter

func JsonStory(r io.Reader) (story Story, err error) {
	d := json.NewDecoder(r)

	if err = d.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}

type Chapter struct {
	Title   string   `json:"title"`
	Paragraphs   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Chapter  string `json:"arc"`
}