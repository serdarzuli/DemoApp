package models

// Book represents a book in the library
// @Description Book information
type Book struct {
	ID       int    `json:"id" example:"1"`
	Title    string `json:"title" example:"The Go Programming Language"`
	Author   string `json:"author" example:"Alan A. A. Donovan"`
	ISBN     string `json:"isbn" example:"978-0134190440"`
	Quantity int    `json:"quantity" example:"5"`
} 