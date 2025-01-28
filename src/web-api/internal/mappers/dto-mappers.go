package mappers

import (
	"simple-go-crud/dtos"
	"simple-go-crud/models"
	"time"
)

func CreateNoteRequestToNote(dto dtos.CreateNoteRequest) models.Note {
	return models.Note{
		Title:     dto.Title,
		Content:   dto.Content,
		CreatedAt: time.Now(),
	}
}

func NoteToCreateNoteResponse(note models.Note) dtos.CreateNoteResponse {
	return dtos.CreateNoteResponse{
		Id:        note.Id,
		Title:     note.Title,
		Content:   note.Content,
		CreatedAt: note.CreatedAt.Format(time.RFC3339),
	}
}