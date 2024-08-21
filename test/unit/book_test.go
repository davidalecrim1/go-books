package service_test

import (
	"go-books/internal/service"
	"testing"
)

func TestGetFullBook(t *testing.T) {

	t.Run("Get full book", func(t *testing.T) {
		book := service.Book{
			ID:     1,
			Title:  "The Great Gatsby",
			Author: "F. Scott Fitzgerald",
			Year:   1925,
		}

		got := book.GetFullBook()
		want := "The Great Gatsby by F. Scott Fitzgerald (1925)"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
