package service

import (
	"library/interfaces"
	"library/models"
	"errors"
)

type bookService struct {
	repo interfaces.BookRepository
}

func NewBookService(repo interfaces.BookRepository) interfaces.BookService {
	return &bookService{
		repo: repo,
	}
}

func (s *bookService) CreateBook(book *models.Book) error {
	// Add business logic validation here
	if book.Quantity < 0 {
		return errors.New("quantity cannot be negative")
	}
	if book.Title == "" {
		return errors.New("title cannot be empty")
	}
	return s.repo.Create(book)
}

func (s *bookService) GetBook(id int) (*models.Book, error) {
	return s.repo.Read(id)
}

func (s *bookService) UpdateBook(book *models.Book) error {
	// Add business logic validation here
	if book.Quantity < 0 {
		return errors.New("quantity cannot be negative")
	}
	return s.repo.Update(book)
}

func (s *bookService) DeleteBook(id int) error {
	return s.repo.Delete(id)
}

func (s *bookService) ListBooks() ([]*models.Book, error) {
	return s.repo.List()
} 