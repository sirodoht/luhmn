package main

import (
	"html/template"
	"net/http"
	"fmt"
	"log"

    _ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// database
	_, err := sqlx.Connect("postgres", "user=luhmn dbname=luhmn sslmode=disable")
    if err != nil {
        log.Fatalln(err)
    }

	// router
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

	fmt.Println("Listening on http://127.0.0.1:8000/")
	http.ListenAndServe(":8000", r)
}
