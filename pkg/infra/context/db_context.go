package dbcontext

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	dbDriverName string = os.Getenv("DB_DRIVER_NAME")
	dbSourceName string = os.Getenv("DB_SOURCE_NAME")
)

type IDbContext interface {
	GetConnection() (*sql.DB, error)
}

type DbContext struct{}

func NewDbContext() IDbContext {
	return &DbContext{}
}

func (*DbContext) GetConnection() (*sql.DB, error) {
	if dbDriverName == "sqlite3" {
		if _, err := os.Stat(dbSourceName); err != nil {
			log.Fatalf("Database file '%v' not exists", dbSourceName)
			return nil, err
		}
	}

	db, err := sql.Open(dbDriverName, dbSourceName)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
		return nil, err
	}
	return db, nil
}
