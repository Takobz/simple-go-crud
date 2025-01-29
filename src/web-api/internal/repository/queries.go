package repository

const (
	INSERT_NOTE_QUERY_RETURNING_ID = "INSERT INTO notes (title, content, created_at) VALUES (@title, @content, @createdAt) RETURNING id"
	SELECT_ALL_NOTES_QUERY = "SELECT * FROM notes"
	SELECT_NOTE_BY_ID_QUERY = "SELECT * FROM notes WHERE id = $1"
)