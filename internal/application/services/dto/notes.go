package dto

type CreateNoteRequest struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}
