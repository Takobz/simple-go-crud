package repository

import (
	"context"
	"simple-go-crud/configuration"
	"simple-go-crud/models"

	"github.com/jackc/pgx/v5"
)

/*
Things to read about:
- context package.
*/
var connectionString = configuration.GetPostgresConnectionString()

func CreateANote(note models.Note) models.Note {
	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	arguments := pgx.NamedArgs{
		"title":     note.Title,
		"content":   note.Content,
		"createdAt": note.CreatedAt,
	}

	err = conn.QueryRow(context.Background(), INSERT_NOTE_QUERY_RETURNING_ID, arguments).Scan(&note.Id)
	if err != nil {
		panic(err)
	}

	return note
}

func GetAllNotes() []models.Note {
	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	var notes []models.Note
	rows, err := conn.Query(context.Background(), SELECT_ALL_NOTES_QUERY)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var note models.Note
		err = rows.Scan(&note.Id, &note.Title, &note.Content, &note.CreatedAt)
		if err != nil {
			panic(err)
		}
		notes = append(notes, note)
	}

	return notes
}
