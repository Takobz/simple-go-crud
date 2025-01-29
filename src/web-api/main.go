package main

import (
	"simple-go-crud/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/note", handlers.CreateNote)
	router.GET("/notes", handlers.GetAllNotes)
	router.GET("/note", handlers.GetNoteById)
	router.Run(":8080")
}
