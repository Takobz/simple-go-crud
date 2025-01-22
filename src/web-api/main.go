package main

import (
	"fmt"
	"simple-go-crud/repository"
	"time"
)

func main() {
	note := repository.Note{
		Title:     "My first note",
		Content:   "This is the content of my first note",
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	note = repository.CreateANote(note)
	fmt.Println(note)

	notes := repository.GetAllNotes()
	fmt.Println(notes)
}
