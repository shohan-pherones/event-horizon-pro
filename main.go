package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shohan-pherones/event-horizon-pro/db"
	"github.com/shohan-pherones/event-horizon-pro/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.InvokingRoutes(server)
	server.Run(":8080")
}
