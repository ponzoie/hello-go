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

	r.Get("/", func(w http.ResponseWriter,r *http.Request){
		w.Header().Set("Content-Type","text/html; charset=utf-8")
		w.Write([]byte(`<!doctype html>
		<html lang="ja">
		<head><meta charset="utf-8"><title>hello-go</title></head>
		<body>
			<h1>んなー</h1>
			<ul>
				<li><a href="/health">/health</a></li>
				<li><a href="/items/1">/items/1</a></li>
			</ul>
		</body>
		</html>`))
	})

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
