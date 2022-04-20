package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/sirodoht/luhmn/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	database.Connect()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RedirectSlashes)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			panic(err)
		}
		t.Execute(w, nil)
	})

	r.Get("/docs/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			panic(err)
		}
		t.Execute(w, nil)
	})

	fmt.Println("Listening on http://127.0.0.1:8000/")
	http.ListenAndServe(":8000", r)
}
