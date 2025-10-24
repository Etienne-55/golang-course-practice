package routes

import (
	"go_backend/middleware"

	"github.com/gin-gonic/gin"
)


func RegisterRoutes(server *gin.Engine ) {
	server.POST("/events", middleware.Authenticate, CreateEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.GET("/events", GetEvents)
	authenticated.GET("/events/:id", GetEvent)
	authenticated.PUT("/events/:id", UpdateEvent)
	authenticated.DELETE("/events/:id", DeleteEvent)
	authenticated.POST("/events/:id/register", registerForEvents)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

}

