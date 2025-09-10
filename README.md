О проекте:  
CRUD-проект на Go с использованием Gorilla Mux и SQLite3.

Структура:  
Схема бд лежит в ../db/scheme_db.svg  
Запросы к бд в ../db/queries/...  
Подключение к бд в ../db/db.go  
Модели и структуры в ../struct/...  
Роутинг и маршрутизация в ../routes/router.go
CRUD эндпоинты в ../handlers/...

Запуск проекта:  
git clone https://github.com/ReynekeJluc/go_study.git  
cd go_study  
go mod tidy  
go run main.go  
REST запросы лежат в ../http/test.http
