package handler

import (
	"html/template"
	"net/http"
)

type PageHandler struct {
	indexTmpl *template.Template
}

func NewPageHandler(indexTmpl *template.Template) *PageHandler {
	return &PageHandler { indexTmpl: indexTmpl}
}

func (h *PageHandler) Index(w http.ResponseWriter , r *http.Request) {
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	_ = h.indexTmpl.Execute(w,map[string]any{
		"Title" : "んなー",
	})
}