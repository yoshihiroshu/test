package model

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type DBContext struct {
	DB *sql.DB
	// TODO redisなど追加
}

func New() *DBContext {

	db, err := GetDBConnection()
	if err != nil {
		log.Fatalln(err)
	}

	return &DBContext{DB: db}
}

func GetDBConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("Fatal get Connection with DB, err=%s\n", err.Error())
		return nil, err
	}
	return db, nil
}
