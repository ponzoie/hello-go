package httpapi

import (
	"net/http"

	"example.com/hello-go/internal/httpapi/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handlers struct {
	Page *handler.PageHandler
	Health *handler.HealthHandler
	Item *handler.ItemHandler
}

func NewRouter(h Handlers) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", h.Page.Index)
	r.Get("/health",h.Health.Get)
	r.Get("/items.{id}", h.Item.Get)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not found", http.StatusNotFound)
	})
	return r
}

