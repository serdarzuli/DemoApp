package interfaces

import "library/models"

type BookService interface {
	CreateBook(*models.Book) error
	GetBook(id int) (*models.Book, error)
	UpdateBook(*models.Book) error
	DeleteBook(id int) error
	ListBooks() ([]*models.Book, error)
} 