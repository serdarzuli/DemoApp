package interfaces

import "library/models"

type BookRepository interface {
	Create(*models.Book) error
	Read(id int) (*models.Book, error)
	Update(*models.Book) error
	Delete(id int) error
	List() ([]*models.Book, error)
} 