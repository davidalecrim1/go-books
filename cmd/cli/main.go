package main

import (
	"database/sql"
	"go-books/internal/cli"
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

	if len(os.Args) > 1 && (os.Args[1] == "search" || os.Args[1] == "simulate") {
		bookCli := cli.NewBookCli(bookService)
		bookCli.Run()
		return
	}
}
