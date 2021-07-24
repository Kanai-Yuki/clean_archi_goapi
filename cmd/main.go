package main

import (
	"log"
	"net/http"

	"github.com/Kanai-Yuki/clean_archi_goapi/internal/router"
)

func main() {
	if err := http.ListenAndServe(":8080", router.CreateRouter()); err != nil {
		log.Fatalln(err)
	}
}
