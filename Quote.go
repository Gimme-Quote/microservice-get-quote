package quote

type Quote struct {
	Author   string `json:"author"`
	Quote    string `json:"quote"`
	Category string `json:"category"`
	Year     string `json:"year"`
}
