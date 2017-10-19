package hello

import (
	"encoding/json"
	"net/http"
)

type Quote struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

var quotes = []*Quote{
	{
		Author: "Paula Scher",
		Quote:  "Less is more and more is more. It’s the middle that’s not a good place.",
	},
	{
		Author: "Howard Aiken",
		Quote:  "Don’t worry about people stealing an idea. If it’s original, you will have to ram it down their throats.",
	},
	{
		Author: "Jack Kerouac",
		Quote:  "Great things are not accomplished by those who yield to trends and fads and popular opinion.",
	},
	{
		Author: "Frank Chimero",
		Quote:  "To make pearls, you've got to eat dirt.",
	},
	{
		Author: "Cassie McDaniel",
		Quote:  "Designers must be stewards of design rather than dictators.",
	},
}

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(quotes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
