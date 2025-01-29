package handlers

import (
	"net/http"
	"simple-go-crud/dtos"
	"simple-go-crud/helpers"
	"simple-go-crud/mappers"
	"simple-go-crud/models"
	"simple-go-crud/repository"

	"github.com/gin-gonic/gin"
)

func CreateNote(ginContext *gin.Context) {
	var createNoteRequest dtos.CreateNoteRequest
	err := ginContext.BindJSON(&createNoteRequest)
	if err != nil {
		badResponse := helpers.CreateErrorResponse(models.CREATE_NOTE_ERROR)
		ginContext.JSON(http.StatusBadRequest, badResponse)
		return
	}

	note := mappers.CreateNoteRequestToNote(createNoteRequest)
	createdNote := repository.CreateANote(note)
	ginContext.JSON(http.StatusOK, mappers.NoteToCreateNoteResponse(createdNote))
}

func GetAllNotes(ginContext *gin.Context) {
	notes := repository.GetAllNotes()
	var response []dtos.GetNoteResponse
	for _, note := range notes {
		response = append(response, mappers.NoteToGetNoteResponse(note))
	}

	ginContext.JSON(http.StatusOK, response)
}

func GetNoteById(ginContext *gin.Context) {
	noteId := ginContext.Query("noteId")
	if noteId == "" {
		badResponse := helpers.CreateErrorResponse(models.MISSING_NOTE_ID)
		ginContext.JSON(http.StatusBadRequest, badResponse)
		return
	}
	note := repository.GetNoteById(noteId)
	ginContext.JSON(http.StatusCreated, mappers.NoteToGetNoteResponse(note))
}
