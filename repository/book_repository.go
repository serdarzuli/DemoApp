package repository

import (
	"errors"
	"library/interfaces"
	"library/models"
	"sync"
)

type bookRepository struct {
	books map[int]*models.Book
	mutex sync.RWMutex
}

// NewBookRepository creates a new instance of BookRepository
func NewBookRepository() interfaces.BookRepository {
	return &bookRepository{
		books: make(map[int]*models.Book),
	}
}

// Create adds a new book to the repository
func (r *bookRepository) Create(book *models.Book) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.books[book.ID]; exists {
		return errors.New("book already exists")
	}

	r.books[book.ID] = book
	return nil
}

// Read retrieves a book by ID
func (r *bookRepository) Read(id int) (*models.Book, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	book, exists := r.books[id]
	if !exists {
		return nil, errors.New("book not found")
	}

	return book, nil
}

// Update modifies an existing book
func (r *bookRepository) Update(book *models.Book) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.books[book.ID]; !exists {
		return errors.New("book not found")
	}

	r.books[book.ID] = book
	return nil
}

// Delete removes a book from the repository
func (r *bookRepository) Delete(id int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.books[id]; !exists {
		return errors.New("book not found")
	}

	delete(r.books, id)
	return nil
}

// List returns all books in the repository
func (r *bookRepository) List() ([]*models.Book, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	books := make([]*models.Book, 0, len(r.books))
	for _, book := range r.books {
		books = append(books, book)
	}
	return books, nil
} 