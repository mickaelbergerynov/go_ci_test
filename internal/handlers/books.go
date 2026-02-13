package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/burger/go-ci-books/internal/store"
)

type BookHandler struct {
	Store *store.Store
}

func (h *BookHandler) ListBooks(c *gin.Context) {
	c.JSON(http.StatusOK, h.Store.GetAllBooks())
}

func (h *BookHandler) GetBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	book, err := h.Store.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}
