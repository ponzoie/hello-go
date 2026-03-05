package main

import (
	"html/template"
	"log"
	"net/http"

	"example.com/hello-go/internal/httpapi"
	"example.com/hello-go/internal/httpapi/handler"
	"example.com/hello-go/web"
)

func mustTemplate(name string) *template.Template {
	t, err := template.ParseFS(web.Templates, name)
	if err != nil {
		panic(err)
	}
	return t
}

func main() {
	indexTmpl := mustTemplate("templates/index.html")

	page := handler.NewPageHandler(indexTmpl)
	health := &handler.HealthHandler{}
	item := handler.NewItemHandler()

	router := httpapi.NewRouter(httpapi.Handlers{
		Page:   page,
		Health: health,
		Item:   item,
	})

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}