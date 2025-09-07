package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	database "github.com/ReynekeJluc/go_study.git/db"
	models "github.com/ReynekeJluc/go_study.git/struct/models"
	mux "github.com/gorilla/mux"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["BookId"]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "id необходимо число",
		})
		return
	}

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "некорректный json",
		})
		return
	}


	res, err := database.DB.Exec(`
		UPDATE Book 
		SET
			book_name = ?,
			book_total_quantity = ?,
			book_pages_amount = ?,
			book_desc = ?,
			book_price = ?,
			book_cover = ?,
			book_super_cover = ?,
			book_publisher = ?,
			book_year_release = ?,
			book_isbn = ?
		WHERE book_id = ?
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
		id,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "ошибка при обновление книги",
		})
		log.Println("ошибка обновления в БД:", err)
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "книга с таким id не найдена",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}