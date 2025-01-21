package repository

type Note struct {
	Id        string `db:"id"`
	Title     string `db:"title"`
	Content   string `db:"content"`
	CreatedAt string `db:"created_at"`
}
