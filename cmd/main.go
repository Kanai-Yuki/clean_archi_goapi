package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/Kanai-Yuki/clean_archi_goapi/internal/router"
)

func main() {
	var db *sql.DB

	db, err := sql.Open("postgres", "host=postgres port=5432 user=user password=password dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	if err := http.ListenAndServe(":8080", router.CreateRouter()); err != nil {
		log.Fatalln(err)
	}
}
