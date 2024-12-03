package main

import (
	"github.com/gin-gonic/gin"
	"github.com/synapsis-challenge/db"
	"github.com/synapsis-challenge/routes"
)

func main() {

	db.InitDB()

	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":8080")
}
