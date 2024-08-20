package service

import (
	"database/sql"
	"strconv"
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
	db *sql.DB
}

func NewBookService(db *sql.DB) *BookService {
	return &BookService{db: db}
}

func (s *BookService) CreateBook(book *Book) error {
	query := "INSERT INTO books (title, author, genre, year) VALUES (?, ?, ?, ?)"

	_, err := s.db.Exec(query, book.Title, book.Author, book.Genre, book.Year)

	if err != nil {
		return err
	}

	return nil
}

func (s *BookService) GetBooks() ([]Book, error) {
	query := "SELECT id, title, author, genre, year FROM books"

	rows, err := s.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []Book

	for rows.Next() {
		var book Book

		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.Year)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (s *BookService) GetBook(id int) (*Book, error) {
	query := "SELECT id, title, author, genre, year FROM books WHERE id = ?"

	row := s.db.QueryRow(query, id)

	var book Book

	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.Year)

	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (s *BookService) UpdateBook(book *Book) error {
	query := "UPDATE books SET title = ?, author = ?, genre = ?, year = ? WHERE id = ?"
	_, err := s.db.Exec(query, book.Title, book.Author, book.Genre, book.Year, book.ID)

	return err
}

func (s *BookService) DeleteBook(id int) error {
	query := "DELETE FROM books WHERE id = ?"
	_, err := s.db.Exec(query, id)

	return err
}
