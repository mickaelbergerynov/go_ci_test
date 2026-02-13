package store

import (
	"testing"
)

func TestGetAllBooks(t *testing.T) {
	s := New()
	books := s.GetAllBooks()
	if len(books) != 10 {
		t.Errorf("expected 10 books, got %d", len(books))
	}
}

func TestGetBookByID(t *testing.T) {
	s := New()

	book, err := s.GetBookByID(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if book.Title != "Dune" {
		t.Errorf("expected title 'Dune', got %q", book.Title)
	}
}

func TestGetBookByID_NotFound(t *testing.T) {
	s := New()

	_, err := s.GetBookByID(999)
	if err == nil {
		t.Fatal("expected error for non-existent book, got nil")
	}
}

func TestGetAllAuthors(t *testing.T) {
	s := New()
	authors := s.GetAllAuthors()
	if len(authors) != 5 {
		t.Errorf("expected 5 authors, got %d", len(authors))
	}
}

func TestGetAuthorByID(t *testing.T) {
	s := New()

	author, err := s.GetAuthorByID(2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if author.Name != "Isaac Asimov" {
		t.Errorf("expected name 'Isaac Asimov', got %q", author.Name)
	}
}

func TestGetAuthorByID_NotFound(t *testing.T) {
	s := New()

	_, err := s.GetAuthorByID(999)
	if err == nil {
		t.Fatal("expected error for non-existent author, got nil")
	}
}
