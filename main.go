package main

import (
	"log"

	"example.com/rest/db"
	"example.com/rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.InitDb(); err != nil {
		log.Fatalf("db init failed: %v", err)
	}
	defer db.Pool.Close()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
