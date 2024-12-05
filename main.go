package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/synapsis-challenge/db"
	"github.com/synapsis-challenge/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDB()

	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":8080")
}
