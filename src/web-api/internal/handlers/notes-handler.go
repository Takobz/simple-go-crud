package handlers

import (
	"fmt"
	"net/http"
	"simple-go-crud/dtos"
	"simple-go-crud/mappers"

	"github.com/gin-gonic/gin"
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

	note := mappers.CreateNoteRequestToNote(createNoteRequest)
	output := fmt.Sprintf("%s", note)
	ginContext.JSON(http.StatusOK, gin.H{
		"message": output,
	})
}

func GetAllNotes(ginContext *gin.Context) {
	ginContext.JSON(http.StatusOK, gin.H{
		"message": "Get all notes",
	})
}
