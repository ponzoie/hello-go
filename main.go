package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var items = map[int]Item{
	1: {ID: 1, Name: "pen"},
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// るーと定義
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok\n"))
	})

	// json kaesu
	r.Get("/items/{id}", getItem)

	// 404
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not found", http.StatusNotFound)
	})

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getItem(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	item, ok := items[id]
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(item)
}

/*
{
	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request) {
		w.Write([]byte("ok\n"))
	})
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080",nil))
}

*/
