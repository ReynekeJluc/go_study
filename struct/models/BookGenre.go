package models

type BookGenre struct {
	BookId  int `db:"book_id"`
	GenreId int `db:"author_id"`
}