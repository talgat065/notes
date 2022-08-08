package domain

import "testing"

func TestCreateNote(t *testing.T) {
	note, err := Note{}.CreateNote("abc", "#music Joy Division - Love Will Tear Us Apart")
	if err != nil {
		t.Error(err)
	}

	if note.Text != "Joy Division - Love Will Tear Us Apart" {
		t.Errorf("Expected note text to be 'Joy Division - Love Will Tear Us Apart', got '%s'", note.Text)
	}

	if note.ID != "abc" {
		t.Errorf("Expected note ID to be '', got '%s'", note.ID)
	}

	if note.Tag.Name != "#music" {
		t.Errorf("Expected note tag name to be 'Hello', got '%s'", note.Tag.Name)
	}
}
