package main

import (
	"backend_article/config"
	// "backend_article/controllers"
	"backend_article/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":8080")
}
