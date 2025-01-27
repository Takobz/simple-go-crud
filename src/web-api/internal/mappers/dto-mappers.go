package mappers

import (
	"simple-go-crud/dtos"
	"simple-go-crud/models"
)

func CreateNoteRequestToNote(dto dtos.CreateNoteRequest)  models.Note {
	return models.Note{
		Title: dto.Title,
		Content: dto.Content,
	}
}