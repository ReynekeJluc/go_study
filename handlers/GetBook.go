package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	database "github.com/ReynekeJluc/go_study.git/db"
	models "github.com/ReynekeJluc/go_study.git/struct/models"
	mux "github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["BookId"]
	id, err := strconv.Atoi(idStr)
	// fmt.Println(idStr)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "id должно быть числом",
		})
		return
	}
	
	var book models.Book
	err = database.DB.QueryRow(`
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
		WHERE book_id = ?
	`, id).Scan(
		&book.Id,
		&book.Name,
		&book.TotalQuantity,
		&book.PagesAmount,
		&book.Desc,
		&book.Price,
		&book.Cover,
		&book.SuperCover,
		&book.Publisher,
		&book.YearRelease,
		&book.ISBN,
	)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "книга не найдена",
		})
		return
	}

	// log.Println(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(book)
}