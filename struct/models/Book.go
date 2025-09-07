package models

type Book struct {
	Id            int     `db:"book_id" json:"book_id"`
	Name          string  `db:"book_name" json:"book_name"`
	TotalQuantity int     `db:"book_total_quantity" json:"book_total_quantity"`
	PagesAmount   int     `db:"book_pages_amount" json:"book_pages_amount"`
	Desc          *string `db:"book_desc" json:"book_desc"`
	Price         float64 `db:"book_price" json:"book_price"`
	Cover         bool    `db:"book_cover" json:"book_cover"`
	SuperCover    bool    `db:"book_super_cover" json:"book_super_cover"` // 0/1
	Publisher     *string `db:"book_publisher" json:"book_publisher"`     // 0/1
	YearRelease   int     `db:"book_year_release" json:"book_year_release"`
	ISBN          string  `db:"book_isbn" json:"book_isbn"`

	// Authors []Author
	// Genres  []Genre
}
