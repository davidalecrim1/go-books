package main

import (
	"database/sql"
	"go-books/internal/repository"
	"go-books/internal/service"
	"go-books/internal/web"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// database init
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	bookRepository := repository.NewBookDBRepository(db)
	bookRepository.SetupSchema("./sql/schema.sql")
	bookService := service.NewBookService(bookRepository)
	bookHandler := web.NewBookHandler(bookService)

	//web http init
	router := http.NewServeMux()
	router.HandleFunc("POST /books", bookHandler.CreateBook)
	router.HandleFunc("POST /books/simulations/read", bookHandler.SimulateMultipleReading)
	router.HandleFunc("GET /books", bookHandler.GetBooks)
	router.HandleFunc("GET /books/{id}", bookHandler.GetBookById)
	router.HandleFunc("PUT /books/{id}", bookHandler.UpdateBook)
	router.HandleFunc("DELETE /books/{id}", bookHandler.DeleteBook)

	//server init
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("server is running on port %v", port)
	err = http.ListenAndServe(":"+port, router)

	if err != nil {
		log.Fatal(err)
	}
}
