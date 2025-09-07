package models

type BookAuthor struct {
	BookId   int `db:"book_id"`
	AuthorId int `db:"author_id"`
}