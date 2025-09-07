package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	database "github.com/ReynekeJluc/go_study.git/db"
	mux "github.com/gorilla/mux"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
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

	var exists int
	err = database.DB.QueryRow("SELECT COUNT(*) FROM Book WHERE book_id = ?", id).Scan(&exists)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "ошибка проверки книги"})
		log.Println("Ошибка проверки существования книги:", err)
		return
	}
	if exists == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "книга не найдена"})
		return
	}

	_, err = database.DB.Exec(`
		DELETE FROM Book 
		WHERE book_id = ?
		`, 
		id,
	)	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "ошибка при удаление книги",
		})
		log.Println("ошибка обновления в БД:", err)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "успешно удалено",
	})
}