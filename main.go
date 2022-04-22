package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/sirodoht/luhmn/document"
	_ "github.com/lib/pq"
)

type Document struct {
	ID        int    `db:"id"`
	CreatedAt string `db:"created_at"`
}

func main() {
	databaseUrl := os.Getenv("DATABASE_URL")
	db, err := sqlx.Connect("postgres", databaseUrl)
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			panic(err)
		}
		t.Execute(w, nil)
	})

	r.Post("/docs", func(w http.ResponseWriter, r *http.Request) {
		type ReqBody struct {
			Title string
			Body  string
		}
		decoder := json.NewDecoder(r.Body)
		var rb ReqBody
		err := decoder.Decode(&rb)
		if err != nil {
			panic(err)
		}

		now := time.Now()
		_, err = db.NamedExec(`INSERT INTO documents (title, body, created_at, updated_at) VALUES (:title, :body, :created_at, :updated_at)`,
			map[string]interface{}{
				"title":      rb.Title,
				"body":       rb.Body,
				"created_at": now,
				"updated_at": now,
			})
		if err != nil {
			panic(err)
		}
	})

	r.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
		docs := []Document{}
		db.Select(&docs, "SELECT * FROM documents ORDER BY id ASC")
		fmt.Println(docs)
		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			panic(err)
		}
		t.Execute(w, nil)
	})

	fmt.Println("Listening on http://127.0.0.1:8000/")
	http.ListenAndServe(":8000", r)
}
