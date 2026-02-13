package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/burger/go-ci-books/internal/store"
)

type AuthorHandler struct {
	Store *store.Store
}

func (h *AuthorHandler) ListAuthors(c *gin.Context) {
	c.JSON(http.StatusOK, h.Store.GetAllAuthors())
}

func (h *AuthorHandler) GetAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	author, err := h.Store.GetAuthorByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "author not found"})
		return
	}

	c.JSON(http.StatusOK, author)
}
