package configs

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

var(
	DB = InitDB(NewConfig().DataBaseURL)
	//переменная "DB" создает новую конфикурацию, получает путь к базе данных и инициализирует подключение к БД.
)

func InitDB(databaseURL string) *sql.DB {
	log.Println("Using DB URL:", databaseURL)
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
	log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
	log.Fatal(err)
	}
	
	return db
	
}