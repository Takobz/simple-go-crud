package repository

import (
	"context"
	"simple-go-crud/configuration"

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

	_, err = conn.Exec(
		context.Background(),
		"INSERT INTO notes (title, content, createdAt) VALUES (?, ?, ?)",
		note.Title,
		note.Content,
		note.CreatedAt)

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
		err = rows.Scan(&note.Id, &note.Title, &note.Content, &note.CreatedAt)
		if err != nil {
			panic(err)
		}
		notes = append(notes, note)
	}

	return notes
}
