package database

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() {
	databaseUrl := os.Getenv("DATABASE_URL")
	_, err := sqlx.Connect("postgres", databaseUrl)
	if err != nil {
		log.Fatalln(err)
	}
}
