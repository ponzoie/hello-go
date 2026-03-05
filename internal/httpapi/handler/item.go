package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Item struct {
	ID int `json:"id"`
	Name string `json:"Name"`
}

type ItemHandler struct {
	items map[int]Item
}

func NewItemHandler() *ItemHandler {
	return &ItemHandler{
		items:map[int]Item{
			1:{ID:1,Name:"pen"},
		},
	}
}

func (h *ItemHandler) Get(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r,"id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id",http.StatusBadRequest)
		return 
	}
	
	item,ok := h.items[id]
	if !ok {
		http.Error(w,"not found",http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type","application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(item)
}