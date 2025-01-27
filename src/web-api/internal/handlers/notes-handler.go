package handlers

import (
	"net/http"
	"simple-go-crud/dtos"
	"github.com/gin-gonic/gin"
	"simple-go-crud/mappers"
)

func CreateNote(ginContext *gin.Context) {
	var createNoteRequest dtos.CreateNoteRequest
	err := ginContext.BindJSON(&createNoteRequest)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	note := mappers.CreateNoteRequestToNoteDto(createNoteRequest)
	ginContext.JSON(http.StatusOK, gin.H{
		"message": "Create a note",
	})
}

func GetAllNotes(ginContext *gin.Context) {
	ginContext.JSON(http.StatusOK, gin.H{
		"message": "Get all notes",
	})
}
