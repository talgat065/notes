package repositories

import (
	"fmt"
	"notes/internal/domain"
)

type NotesMemoryRepository struct {
	notes []domain.Note
}

func NewNotesMemoryRepository() NotesMemoryRepository {
	return NotesMemoryRepository{notes: []domain.Note{}}
}

func (r NotesMemoryRepository) Save(note domain.Note) error {
	r.notes = append(r.notes, note)
	fmt.Println("note saved:", r.notes)
	return nil
}
