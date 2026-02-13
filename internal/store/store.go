package store

import (
	"fmt"
	"sync"

	"github.com/burger/go-ci-books/internal/models"
)

type Store struct {
	mu      sync.RWMutex
	books   map[int]models.Book
	authors map[int]models.Author
}

func New() *Store {
	s := &Store{
		books:   make(map[int]models.Book),
		authors: make(map[int]models.Author),
	}
	s.seed()
	return s
}

func (s *Store) GetAllAuthors() []models.Author {
	s.mu.RLock()
	defer s.mu.RUnlock()

	authors := make([]models.Author, 0, len(s.authors))
	for _, a := range s.authors {
		authors = append(authors, a)
	}
	return authors
}

func (s *Store) GetAuthorByID(id int) (models.Author, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	a, ok := s.authors[id]
	if !ok {
		return models.Author{}, fmt.Errorf("author not found")
	}
	return a, nil
}

func (s *Store) GetAllBooks() []models.Book {
	s.mu.RLock()
	defer s.mu.RUnlock()

	books := make([]models.Book, 0, len(s.books))
	for _, b := range s.books {
		books = append(books, b)
	}
	return books
}

func (s *Store) GetBookByID(id int) (models.Book, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	b, ok := s.books[id]
	if !ok {
		return models.Book{}, fmt.Errorf("book not found")
	}
	return b, nil
}

func (s *Store) seed() {
	s.authors = map[int]models.Author{
		1: {ID: 1, Name: "Frank Herbert"},
		2: {ID: 2, Name: "Isaac Asimov"},
		3: {ID: 3, Name: "Ursula K. Le Guin"},
		4: {ID: 4, Name: "Philip K. Dick"},
		5: {ID: 5, Name: "Arthur C. Clarke"},
	}

	s.books = map[int]models.Book{
		1:  {ID: 1, Title: "Dune", AuthorID: 1, Year: 1965},
		2:  {ID: 2, Title: "Foundation", AuthorID: 2, Year: 1951},
		3:  {ID: 3, Title: "The Left Hand of Darkness", AuthorID: 3, Year: 1969},
		4:  {ID: 4, Title: "Do Androids Dream of Electric Sheep?", AuthorID: 4, Year: 1968},
		5:  {ID: 5, Title: "2001: A Space Odyssey", AuthorID: 5, Year: 1968},
		6:  {ID: 6, Title: "Childhood's End", AuthorID: 5, Year: 1953},
		7:  {ID: 7, Title: "The Dispossessed", AuthorID: 3, Year: 1974},
		8:  {ID: 8, Title: "I, Robot", AuthorID: 2, Year: 1950},
		9:  {ID: 9, Title: "Dune Messiah", AuthorID: 1, Year: 1969},
		10: {ID: 10, Title: "A Scanner Darkly", AuthorID: 4, Year: 1977},
	}
}
