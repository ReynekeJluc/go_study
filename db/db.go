package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	godotenv "github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

// глобалка дял подключения
var DB *sql.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("ошибка загрузки env переменных: %s", err)
	}
}

// ConnectDB открывает соединение с SQLite и возвращает объект *sql.DB
func ConnectDB() (*sql.DB, error) {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "database.db"
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть базу: %w", err)
	}

	// проверка внешних ключей
	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return nil, fmt.Errorf("не удалось включить foreign_keys: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ошибка соединения с базой: %w", err)
	}

	DB = db
	log.Println("подключение к бд успешно")

	return db, nil
}
