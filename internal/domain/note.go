package domain

import "strings"

type Note struct {
	ID        string `json:"id"`
	Tag       Tag    `json:"tag"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

type NoteRepository interface {
	Save(note Note) error
}

func (n Note) CreateNote(id string, text string) (Note, error) {
	text = strings.TrimSpace(text)
	if !strings.HasPrefix(text, "#") {
		return Note{}, ErrWrongPrefix
	}

	tagName, msg := "", ""
	for i, v := range strings.SplitN(text, " ", 2) {
		if i == 0 {
			tagName = v
		} else if i == 1 {
			msg = v
		}
	}

	newNote := Note{
		ID: id,
		Tag: Tag{
			Name: tagName,
		},
		Text: msg,
	}

	return newNote, nil
}

type Tag struct {
	Name string `json:"name"`
}
