package handler

import (
	"database/sql"
	"encoding/json"
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

	var rating RateQuote
	err := json.NewDecoder(r.Body).Decode(&rating)
	if err != nil {
		http.Error(w, "Could not parse the JSON payload.", http.StatusBadRequest)
		return
	}

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

	query := fmt.Sprintf("UPDATE htmx_playaround.Quotes SET ratings = array_append(ratings, %d) WHERE id = %d", rating.Rating, rating.QuoteId)

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	// todo: main GET API to return quoteId
	// todo: refetch quote after updating the rating


    fmt.Fprintf(w, "<div>OK</div>")
}
