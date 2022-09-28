package book

import (
	"errors"
	"fmt"
)

type fileRepository struct {
	// FindAll() ([]Book, error)
	// FindById(ID int) (Book, error)
	// Create(book Book) (Book, error)
}

func NewFileRepository() *fileRepository {
	return &fileRepository{}
}

func (r *fileRepository) FindAll() ([]Book, error) {
	var books []Book
	fmt.Println("Find All")
	return books, errors.New("Dummy")
}

func (r *fileRepository) FindById(ID int) (Book, error) {
	var book Book
	fmt.Println("Find By ID")
	return book, nil
}

func (r *fileRepository) Create(book Book) (Book, error) {
	fmt.Println("CREATE")
	return book, nil
}
