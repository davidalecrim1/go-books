package main

import (
	"database/sql"
	"go-books/internal/cli"
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

	sqlFile, err := os.ReadFile("./sql/schema.sql")

	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(string(sqlFile))

	if err != nil {
		log.Fatal(err)
	}

	log.Println("sql script for schema executed successfully!")

	bookService := service.NewBookService(db)
	bookHandler := web.NewBookHandler(bookService)

	// cli init
	if len(os.Args) > 1 && (os.Args[1] == "search" || os.Args[1] == "simulate") {
		bookCli := cli.NewBookCli(bookService)
		bookCli.Run()
		return
	}

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
