package repository

import (
	"database/sql"
	"fmt"
	"go-books/internal/service"
	"os"
)

type BookDBRepository struct {
	db *sql.DB
}

func NewBookDBRepository(db *sql.DB) *BookDBRepository {
	return &BookDBRepository{db: db}
}

// SetupSchema loads and executes the SQL schema from a file.
func (r *BookDBRepository) SetupSchema(schemaPath string) error {
	sqlFile, err := os.ReadFile(schemaPath)
	if err != nil {
		return fmt.Errorf("failed to read schema file: %w", err)
	}

	_, err = r.db.Exec(string(sqlFile))
	if err != nil {
		return fmt.Errorf("failed to execute schema: %w", err)
	}

	return nil
}

func (r *BookDBRepository) CreateBook(book *service.Book) error {
	query := "INSERT INTO books (title, author, genre, year) VALUES (?, ?, ?, ?)"
	_, err := r.db.Exec(query, book.Title, book.Author, book.Genre, book.Year)

	if err != nil {
		return fmt.Errorf("failed to create book %w", err)
	}

	return nil
}

func (r *BookDBRepository) GetBooks() ([]service.Book, error) {
	query := "SELECT id, title, author, genre, year FROM books"
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, fmt.Errorf("failed to get books %w", err)
	}

	defer rows.Close()

	var books []service.Book

	for rows.Next() {
		var book service.Book

		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.Year)

		if err != nil {
			return nil, fmt.Errorf("failed to scan book %w", err)
		}

		books = append(books, book)
	}

	return books, nil
}

func (r *BookDBRepository) GetBookById(id int) (*service.Book, error) {
	query := "SELECT id, title, author, genre, year FROM books WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var book service.Book

	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.Year)

	if err != nil {
		return nil, fmt.Errorf("failed to get book by id %w", err)
	}

	return &book, nil
}

func (r *BookDBRepository) UpdateBook(book *service.Book) error {
	query := "UPDATE books SET title = ?, author = ?, genre = ?, year = ? WHERE id = ?"
	_, err := r.db.Exec(query, book.Title, book.Author, book.Genre, book.Year, book.ID)
	return err
}

func (r *BookDBRepository) SearchBooksByName(name string) ([]service.Book, error) {
	query := "SELECT id, title, author, genre, year FROM books WHERE title LIKE ?"
	rows, err := r.db.Query(query, "%"+name+"%")

	if err != nil {
		return nil, fmt.Errorf("failed to search books %w", err)
	}

	defer rows.Close()

	var books []service.Book

	for rows.Next() {
		var book service.Book

		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.Year)

		if err != nil {
			return nil, fmt.Errorf("failed to scan book %w", err)
		}

		books = append(books, book)
	}

	return books, nil
}

func (r *BookDBRepository) DeleteBook(id int) error {
	query := "DELETE FROM books WHERE id = ?"
	_, err := r.db.Exec(query, id)

	if err != nil {
		return fmt.Errorf("failed to delete book %w", err)
	}

	return nil
}
