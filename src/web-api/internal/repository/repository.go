package repository

import (
	"context"
	"simple-go-crud/configuration"
	"time"

	"github.com/jackc/pgx/v5"
)

/*
Things to read about:
- context package.
*/
var connectionString = configuration.GetPostgresConnectionString()

func CreateANote(note Note) Note {
	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	query := "INSERT INTO notes (title, content, created_at) VALUES (@title, @content, @createdAt)"
	arguments := pgx.NamedArgs{
		"title":     note.Title,
		"content":   note.Content,
		"createdAt": note.CreatedAt,
	}

	_, err = conn.Exec(context.Background(), query, arguments)
	if err != nil {
		panic(err)
	}

	return note
}

func GetAllNotes() []Note {
	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	var notes []Note
	rows, err := conn.Query(context.Background(), "SELECT * FROM notes")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var note Note
		var createdAt time.Time
		err = rows.Scan(&note.Id, &note.Title, &note.Content, &createdAt)
		if err != nil {
			panic(err)
		}
		note.CreatedAt = createdAt.Format("2006-01-02 15:04:05")
		notes = append(notes, note)
	}

	return notes
}
