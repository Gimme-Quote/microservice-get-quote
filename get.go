package quote

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	category := r.URL.Query().Get("category")
	log.Debugf(ctx, "Category: %v", category)
	q := datastore.NewQuery("Quote")
	if len(category) != 0 {
		q = q.Filter("Category =", category)
	}
	// var quotes []Quote
	var quotes = make([]Quote, 0)
	_, err := q.GetAll(ctx, &quotes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(quotes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
