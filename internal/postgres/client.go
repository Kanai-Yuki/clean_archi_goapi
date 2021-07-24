package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Client struct {
	DB *sql.DB
}

var cli *Client

func New() (*Client, error) {
	db, err := sql.Open("postgres", "host=postgres port=5432 user=user password=password dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	// 接続確認
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return &Client{DB: db}, nil
}
