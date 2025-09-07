package main

import (
	"log"

	database "github.com/ReynekeJluc/go_study.git/db"
	models "github.com/ReynekeJluc/go_study.git/struct/models"
)

func main() {
	var books []models.Book
	rows, err := database.DB.Query(`
			SELECT 
				book_id,
				book_name,
				book_total_quantity,
				book_pages_amount,
				book_desc,
				book_price,
				book_cover,
				book_super_cover,
				book_publisher,
				book_year_release,
				book_isbn
			FROM Book
	`)
	if err != nil {
		log.Fatal("Ошибка:", err)
	}
	defer rows.Close()
	
	for rows.Next() {
		var b models.Book
		err := rows.Scan(
			&b.Id,
			&b.Name,
			&b.TotalQuantity,
			&b.PagesAmount,
			&b.Desc,
			&b.Price,
			&b.Cover,
			&b.SuperCover,
			&b.Publisher,
			&b.YearRelease,
			&b.ISBN,
		)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, b)
	}

	log.Println(books)
}
