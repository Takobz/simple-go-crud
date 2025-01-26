package main

import (
	"simple-go-crud/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/note", handlers.CreateNote)
	router.GET("/note", handlers.GetAllNotes)
	router.Run(":8080")
}
