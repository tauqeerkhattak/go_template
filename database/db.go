package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"template/config"
)

func InitDB() *sql.DB {
	dbDriver := config.Env("DB_DRIVER", "")
	dbSource := config.Env("DB_SOURCE", "")

	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Failed to connect DB:", err)
	}

	log.Println("Database connected and migrated")
	return db
}
