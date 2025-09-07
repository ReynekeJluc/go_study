package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	database "github.com/ReynekeJluc/go_study.git/db"
	models "github.com/ReynekeJluc/go_study.git/struct/models"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "некорректный json",
		})
		return
	}

	res, err := database.DB.Exec(`
		INSERT INTO Book (
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
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`,
		book.Name,
		book.TotalQuantity,
		book.PagesAmount,
		book.Desc,
		book.Price,
		book.Cover,
		book.SuperCover,
		book.Publisher,
		book.YearRelease,
		book.ISBN,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "ошибка при создании книги",
		})
		log.Println("ошибка вставки в БД:", err)
		return
	}

	id, _ := res.LastInsertId()
	book.Id = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}