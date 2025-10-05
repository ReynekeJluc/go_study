package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	database "github.com/ReynekeJluc/go_study.git/db"
	models "github.com/ReynekeJluc/go_study.git/struct/models"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	search := strings.TrimSpace(r.URL.Query().Get("search"))

	var rows *sql.Rows
	var err error

	if search != "" {
		searchPattern := search + "%"
		rows, err = database.DB.Query(`
			SELECT 
				book_id, book_name, book_total_quantity, book_pages_amount,
				book_desc, book_price, book_cover, book_super_cover,
				book_publisher, book_year_release, book_isbn
			FROM Book
			WHERE book_name IS NOT NULL
				AND trim(book_name) LIKE ?
		`, searchPattern)
	} else {
		rows, err = database.DB.Query(`
			SELECT 
				book_id, book_name, book_total_quantity, book_pages_amount,
				book_desc, book_price, book_cover, book_super_cover,
				book_publisher, book_year_release, book_isbn
			FROM Book
		`)
	}


	if err != nil {
		http.Error(w, "ошибка при чтении базы: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var b models.Book
		err := rows.Scan(
			&b.Id, &b.Name, &b.TotalQuantity, &b.PagesAmount,
			&b.Desc, &b.Price, &b.Cover, &b.SuperCover,
			&b.Publisher, &b.YearRelease, &b.ISBN,
		)
		if err != nil {
			http.Error(w, "ошибка при сканировании строки: "+err.Error(), http.StatusInternalServerError)
			return
		}
		books = append(books, b)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "ошибка при чтении строк: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(books)
}
