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

func setupAuthorsRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	s := store.New()
	h := &AuthorHandler{Store: s}

	r := gin.New()
	r.GET("/authors", h.ListAuthors)
	r.GET("/authors/:id", h.GetAuthor)
	return r
}

func TestListAuthors(t *testing.T) {
	r := setupAuthorsRouter()

	req := httptest.NewRequest(http.MethodGet, "/authors", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var authors []models.Author
	if err := json.Unmarshal(w.Body.Bytes(), &authors); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}
	if len(authors) != 5 {
		t.Errorf("expected 5 authors, got %d", len(authors))
	}
}

func TestGetAuthor(t *testing.T) {
	r := setupAuthorsRouter()

	req := httptest.NewRequest(http.MethodGet, "/authors/2", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var author models.Author
	if err := json.Unmarshal(w.Body.Bytes(), &author); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}
	if author.Name != "Isaac Asimov" {
		t.Errorf("expected name 'Isaac Asimov', got %q", author.Name)
	}
}

func TestGetAuthor_NotFound(t *testing.T) {
	r := setupAuthorsRouter()

	req := httptest.NewRequest(http.MethodGet, "/authors/999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", w.Code)
	}
}

func TestGetAuthor_InvalidID(t *testing.T) {
	r := setupAuthorsRouter()

	req := httptest.NewRequest(http.MethodGet, "/authors/abc", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}
