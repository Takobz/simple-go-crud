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

	//TODO: return dto not db model
	note := mappers.CreateNoteRequestToNote(createNoteRequest)
	createdNote := repository.CreateANote(note)
	ginContext.JSON(http.StatusOK, createdNote)
}

func GetAllNotes(ginContext *gin.Context) {
	notes := repository.GetAllNotes()
	ginContext.JSON(http.StatusOK, notes)
}
