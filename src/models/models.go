package models

type Note struct {
	Id        int    `db:"id"`
	Title     string `db:"title"`
	Content   string `db:"content"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
