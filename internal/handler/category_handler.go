package handler

import (
	"my-contacts/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAllCategories(c *gin.Context) {
	categories, err := service.FindAllCategories()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, categories)
}
