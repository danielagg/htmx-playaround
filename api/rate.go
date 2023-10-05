package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type RateQuote struct {
    QuoteId int
    Rating  int
}

func RateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed.", http.StatusMethodNotAllowed)
		return
	}

	quoteId := r.URL.Query().Get("id")
	rating := r.URL.Query().Get("value")

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

	query := fmt.Sprintf("UPDATE htmx_playaround.Quotes SET ratings = array_append(ratings, %s) WHERE id = %s", rating, quoteId)

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	// todo: refetch quote after updating the rating
    fmt.Fprintf(w, "<div>OK</div>")
}
