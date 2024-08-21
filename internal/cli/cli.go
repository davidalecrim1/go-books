package cli

import (
	"fmt"
	"go-books/internal/service"
	"os"
	"strconv"
	"time"
)

type BookCli struct {
	service *service.BookService
}

func NewBookCli(service *service.BookService) *BookCli {
	return &BookCli{
		service: service,
	}
}

func (cli *BookCli) Run() {
	if len(os.Args) < 2 {
		fmt.Println(" Usage: books <command> [arguments]")
	}

	command := os.Args[1]

	switch command {
	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Usage: books search <query>")
			return
		}
		bookName := os.Args[2]
		cli.searchBooks(bookName)

	case "simulate":
		if len(os.Args) < 3 {
			fmt.Println("Usage: books simulate <book_id> <book_id> <book_id> ...")
			return
		}
		bookIds := os.Args[2:]
		cli.simulateReading(bookIds)
	}
}

func (cli *BookCli) searchBooks(bookName string) {
	books, err := cli.service.SearchBooksByName(bookName)

	if err != nil {
		fmt.Println("Error searching for books: ", err)
		return
	}

	if len(books) == 0 {
		fmt.Println("No books found")
		return
	}

	fmt.Printf("%v books found\n", len(books))

	for _, book := range books {
		fmt.Println(book.GetFullBook())
	}
}

func (cli *BookCli) simulateReading(bookIdsStr []string) {
	var bookIds []int

	for _, bookId := range bookIdsStr {
		bookId, err := strconv.Atoi(bookId)
		if err != nil {
			fmt.Println("invalid book id: ", bookId)
			continue
		}
		bookIds = append(bookIds, bookId)
	}

	results := cli.service.SimulateMultipleReading(bookIds, time.Second*2)

	for _, result := range results {
		fmt.Println(result)
	}

	fmt.Println("simulated reading finished!")
}
