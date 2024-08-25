package service

import (
	"fmt"
	"strconv"
	"time"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Year   int    `json:"year"`
}

func (b *Book) GetFullBook() string {
	return b.Title + " by " + b.Author + " (" + strconv.Itoa(b.Year) + ")"
}

type BookService struct {
	repo BookRepository
}

type BookRepository interface {
	CreateBook(book *Book) error
	GetBooks() ([]Book, error)
	GetBookById(id int) (*Book, error)
	UpdateBook(book *Book) error
	DeleteBook(id int) error
	SearchBooksByName(name string) ([]Book, error)
}

func NewBookService(repo BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) CreateBook(book *Book) error {
	return s.repo.CreateBook(book)
}

func (s *BookService) GetBooks() ([]Book, error) {
	return s.repo.GetBooks()
}

func (s *BookService) GetBookById(id int) (*Book, error) {
	return s.repo.GetBookById(id)
}

func (s *BookService) UpdateBook(book *Book) error {
	return s.repo.UpdateBook(book)
}

func (s *BookService) DeleteBook(id int) error {
	return s.repo.DeleteBook(id)
}

func (s *BookService) SimulateReading(bookId int, duration time.Duration, results chan<- string) {
	book, err := s.GetBookById(bookId)

	if err != nil || book == nil {
		results <- fmt.Sprintf("book id %v not found", bookId)
		return
	}

	time.Sleep(duration)
	results <- fmt.Sprintf("finished reading %v", book.GetFullBook())
}

func (s *BookService) SimulateMultipleReading(bookIds []int, duration time.Duration) []string {
	results := make(chan string, len(bookIds))
	defer close(results)

	for _, bookId := range bookIds {
		go func(id int) {
			s.SimulateReading(id, duration, results)
		}(bookId)
	}

	var responses []string

	for range bookIds {
		responses = append(responses, <-results)
	}

	return responses
}

func (s *BookService) SearchBooksByName(name string) ([]Book, error) {
	return s.repo.SearchBooksByName(name)
}
