package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rpraveenkumar1203/Golang/tree/main/REST_API/middleware"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authentication := server.Group("/")

	authentication.Use(middleware.Authenticate)

	authentication.POST("/events", createEvent)
	authentication.PUT("/events/:id", updateEvent)
	authentication.DELETE("/events/:id", deleteEvent)
	authentication.POST("/events/:id/registeration", register)
	authentication.DELETE("/events/:id/registeration", deregister)

	server.POST("/signup/", signUp)
	server.POST("/login/", login)
}
