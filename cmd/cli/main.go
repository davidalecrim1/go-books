package main

import (
	"database/sql"
	"go-books/internal/cli"
	"go-books/internal/repository"
	"go-books/internal/service"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	bookRepository := repository.NewBookDBRepository(db)
	bookRepository.SetupSchema("./sql/schema.sql")
	bookService := service.NewBookService(bookRepository)

	if len(os.Args) > 1 && (os.Args[1] == "search" || os.Args[1] == "simulate") {
		bookCli := cli.NewBookCli(bookService)
		bookCli.Run()
		return
	}
}
