package dtos

type CreateNoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
