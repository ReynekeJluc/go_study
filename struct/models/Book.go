package models

import "database/sql"

type Book struct {
	Id            int            `db:"book_id"`
	Name          string         `db:"book_name"`
	TotalQuantity int            `db:"book_total_quantity"`
	PagesAmount   int            `db:"book_pages_amount"`
	Desc          sql.NullString `db:"book_desc"`
	Price         float64        `db:"book_price"`
	Cover         int            `db:"book_cover"`
	SuperCover    int            `db:"book_super_cover"`  // 0/1
	Publisher     sql.NullString `db:"book_publisher"`    // 0/1
	YearRelease   int            `db:"book_year_release"`
	ISBN          string         `db:"book_isbn"`

	Authors       []Author
	Genres        []Genre
}
