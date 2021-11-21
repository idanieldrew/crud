package models

type Product struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"Body"`
}
