package handler

import (
	"fmt"
	"net/http"
)


func RateHandler(w http.ResponseWriter, r *http.Request) {
	// todo: this can only accept POST
	// todo: body: { quoteId, rating }
	// todo: main GET API to return quoteId
	// todo: refetch quote after updating the rating
	fmt.Println("%s", r.Body)

    fmt.Fprintf(w, "<div>OK</div>")
}
