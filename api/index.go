package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/lib/pq"
)

type Quote struct {
    ID      int
    Quote   string
    Author  string
    Ratings pq.Int64Array
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

    query := "SELECT id, quote, author, ratings FROM htmx_playaround.Quotes LIMIT 1"

    var quote Quote
    err = db.QueryRow(query).Scan(&quote.ID, &quote.Quote, &quote.Author, &quote.Ratings)

    if err != nil {
        log.Fatal(err)
    }

    htmlTemplate := `
<section id="quote">
  <blockquote class="text-5xl italic border-t border-slate-800 pt-16">
    %s
    <br />
    <span class="text-lg opacity-50 font-normal not-italic">
      by %s
    </span>
  </blockquote>
  <div>
    <h2 class="text-xl border-t border-slate-800 mt-16 pt-16">On a scale from 1 to 5, how cringey was that quote?</h2>
    <div class="flex justify-center mt-4 items-center space-x-2">
      <form id="rating-form" class="flex items-center space-x-2">
        <button type="submit" hx-post="/api/rate?id=%d&value=1" hx-target="#quote" class="rounded-full h-10 w-10 flex items-center justify-center text-slate-900 bg-red-400 hover:bg-red-600">1</button>
        <button type="submit" hx-post="/api/rate?id=%d&value=2" hx-target="#quote" class="rounded-full h-10 w-10 flex items-center justify-center text-slate-900 bg-orange-400 hover:bg-orange-600">2</button>
        <button type="submit" hx-post="/api/rate?id=%d&value=3" hx-target="#quote" class="rounded-full h-10 w-10 flex items-center justify-center text-slate-900 bg-yellow-400 hover:bg-yellow-600">3</button>
        <button type="submit" hx-post="/api/rate?id=%d&value=4" hx-target="#quote" class="rounded-full h-10 w-10 flex items-center justify-center text-slate-900 bg-lime-400 hover:bg-lime-600">4</button>
        <button type="submit" hx-post="/api/rate?id=%d&value=5" hx-target="#quote" class="rounded-full h-10 w-10 flex items-center justify-center text-slate-900 bg-green-400 hover:bg-green-600">5</button>
      </form>
    </div>
    <p class="pt-4 opacity-50 text-sm">The current average rating is: %v</p>
  </div>
</section>
`

    var sum int64
    for _, value := range quote.Ratings {
        sum += value
    }

    average := float64(sum) / float64(len(quote.Ratings))
    roundedAverage := strconv.FormatFloat(average, 'f', 2, 64)

    // todo: should not need .ID 5 times... :|
    fmt.Fprintf(w, htmlTemplate, quote.Quote, quote.Author, quote.ID, quote.ID, quote.ID, quote.ID, quote.ID, roundedAverage)
}
