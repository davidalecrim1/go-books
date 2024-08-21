package web

import (
	"encoding/json"
	"go-books/internal/service"
	"log"
	"net/http"
	"strconv"
	"time"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetBooks()

	if err != nil {
		log.Println(err)
		http.Error(w, "failed to get books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// CreateBook with PUT /books/{id}
func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book service.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.service.CreateBook(&book)

	if err != nil {
		http.Error(w, "failed to create book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetBookById with GET /books/{id}
func (h *BookHandler) GetBookById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	if idStr == "" {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "failed to get book", http.StatusInternalServerError)
		return
	}

	book, err := h.service.GetBookById(id)

	if err != nil {
		http.Error(w, "failed to get book", http.StatusInternalServerError)
		return
	}

	if book == nil {
		http.Error(w, "book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// UpdateBook with POST /books/{id}
func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	if idStr == "" {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "failed to update book", http.StatusInternalServerError)
		return
	}

	var book service.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}
	book.ID = id

	err = h.service.UpdateBook(&book)

	if err != nil {
		http.Error(w, "failed to update book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteBook with DELETE /books/{id}
func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	if idStr == "" {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "failed to delete book", http.StatusInternalServerError)
		return
	}

	err = h.service.DeleteBook(id)

	if err != nil {
		http.Error(w, "failed to delete book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

type SimulationRequest struct {
	IDs []int `json:"id"`
}

type SimulationResponse struct {
	Messages []string `json:"messages"`
}

// SimulateMultipleReading with POST /books/simulations/read
func (h *BookHandler) SimulateMultipleReading(w http.ResponseWriter, r *http.Request) {
	var request SimulationRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {

		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	results := h.service.SimulateMultipleReading(request.IDs, time.Second*2)
	response := SimulationResponse{
		Messages: results,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
