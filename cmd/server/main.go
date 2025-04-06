package main

import (
	"log"
	"my-contacts/internal/handler"
	"my-contacts/internal/infra"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.InitDB()
	defer infra.DB.Close()

	router := gin.Default()

	router.GET("/ping", handler.Ping)
	router.GET("/categories", handler.FindAllCategories)
	router.GET("/contacts", handler.FindAllContacts)

	log.Println("🚀 Server running at http://localhost:8080")
	router.Run(":8080")
}
