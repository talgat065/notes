package services

import (
	"notes/internal/application/services/dto"
	"notes/internal/domain"
)

type NotesService struct {
	notes domain.NoteRepository
}

func NewNotesService(noteRepository domain.NoteRepository) NotesService {
	return NotesService{
		notes: noteRepository,
	}
}

func (s *NotesService) CreateNote(request dto.CreateNoteRequest) (string, error) {
	note, err := domain.Note{}.CreateNote(request.ID, request.Text)
	if err != nil {
		return "", err
	}

	return request.ID, s.notes.Save(note)
}
