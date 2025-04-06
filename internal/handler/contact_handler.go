package handler

import (
	"my-contacts/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAllContacts(c *gin.Context) {
	contacts, err := repository.FindAllContacts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, contacts)
}
