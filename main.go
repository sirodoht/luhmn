package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/sirodoht/luhmn/document"
	_ "github.com/lib/pq"
)

func main() {
	databaseUrl := os.Getenv("DATABASE_URL")
	db, err := sqlx.Connect("postgres", databaseUrl)
	if err != nil {
		panic(err)
	}

	store := document.NewSQLStore(db)
	api := document.NewAPI(store)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			panic(err)
		}
		t.Execute(w, nil)
	})

	r.Post("/docs", api.InsertHandler)
	r.Get("/docs", api.GetAllHandler)

	fmt.Println("Listening on http://127.0.0.1:8000/")
	http.ListenAndServe(":8000", r)
}
