package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rpraveenkumar1203/Golang/tree/main/REST_API/db"
	"github.com/rpraveenkumar1203/Golang/tree/main/REST_API/routes"
)

func main() {
	//initialiae get post and satart serbvert\

	db.InitDB()
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
