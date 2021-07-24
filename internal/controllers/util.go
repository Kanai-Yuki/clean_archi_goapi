package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

type RouterFunc func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

func ResponseHandler(w http.ResponseWriter, r *http.Request, routerFunc RouterFunc) {
	body, status, err := routerFunc(w, r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(status)
		return
	}

	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
