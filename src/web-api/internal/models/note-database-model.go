package models

import (
	"time"
)

type Note struct {
	Id        string    `db:"id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}
