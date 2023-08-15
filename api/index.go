package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type Quote struct {
    ID      int
    Quote   string
    Author  string
}
 
func Handler(w http.ResponseWriter, r *http.Request) {
	host := os.Getenv("POSTGRES_HOST")
    port := 5432
    user := os.Getenv("POSTGRES_USER")
    password := os.Getenv("POSTGRES_PASSWORD")
    dbname := os.Getenv("POSTGRES_DB")

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
	host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

	query := "SELECT id, quote, author FROM htmx_playaround.Quotes LIMIT 1"

    var quote Quote
    err = db.QueryRow(query).Scan(&quote.ID, &quote.Quote, &quote.Author)

    if err != nil {
        log.Fatal(err)
    }

	fmt.Fprintf(w, "<blockquote class=\"text-4xl italic\">%s<br /><span class=\"text-lg font-normal not-italic\">by %s\n</span></blockquote>", quote.Quote, quote.Author)
}
