package handler

import "net/http"

type HealthHandler struct{}

func (h *HealthHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok\n"))
}