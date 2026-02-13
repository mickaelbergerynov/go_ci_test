package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/burger/go-ci-books/internal/models"
	"github.com/burger/go-ci-books/internal/store"
)

func setupBooksRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	s := store.New()
	h := &BookHandler{Store: s}

	r := gin.New()
	r.GET("/books", h.ListBooks)
	r.GET("/books/:id", h.GetBook)
	return r
}

func TestListBooks(t *testing.T) {
	r := setupBooksRouter()

	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var books []models.Book
	if err := json.Unmarshal(w.Body.Bytes(), &books); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}
	if len(books) != 10 {
		t.Errorf("expected 10 books, got %d", len(books))
	}
}

func TestGetBook(t *testing.T) {
	r := setupBooksRouter()

	req := httptest.NewRequest(http.MethodGet, "/books/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var book models.Book
	if err := json.Unmarshal(w.Body.Bytes(), &book); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}
	if book.Title != "Dune" {
		t.Errorf("expected title 'Dune', got %q", book.Title)
	}
}

func TestGetBook_NotFound(t *testing.T) {
	r := setupBooksRouter()

	req := httptest.NewRequest(http.MethodGet, "/books/999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", w.Code)
	}
}

func TestGetBook_InvalidID(t *testing.T) {
	r := setupBooksRouter()

	req := httptest.NewRequest(http.MethodGet, "/books/abc", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}
