package main

import (
	"database/sql"
	"go-books/internal/service"
	"go-books/internal/web"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	bookService := service.NewBookService(db)
	bookHandler := web.NewBookHandler(bookService)

	router := http.NewServeMux()
	router.HandleFunc("POST /books", bookHandler.CreateBook)
	router.HandleFunc("GET /books", bookHandler.GetBooks)
	router.HandleFunc("GET /books/{id}", bookHandler.GetBookById)
	router.HandleFunc("PUT /books/{id}", bookHandler.UpdateBook)
	router.HandleFunc("DELETE /books/{id}", bookHandler.DeleteBook)

	http.ListenAndServe(":8080", router)
}
