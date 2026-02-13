package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/burger/go-ci-books/internal/handlers"
	"github.com/burger/go-ci-books/internal/store"
)

func main() {
	s := store.New()
	r := setupRouter(s)

	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func setupRouter(s *store.Store) *gin.Engine {
	r := gin.Default()

	r.GET("/health", handlers.Health)

	bookHandler := &handlers.BookHandler{Store: s}
	r.GET("/books", bookHandler.ListBooks)
	r.GET("/books/:id", bookHandler.GetBook)

	authorHandler := &handlers.AuthorHandler{Store: s}
	r.GET("/authors", authorHandler.ListAuthors)
	r.GET("/authors/:id", authorHandler.GetAuthor)

	return r
}
