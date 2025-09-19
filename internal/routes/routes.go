package routes

import (
	"github.com/Habeebamoo/contact-go/internal/handlers"
	"github.com/Habeebamoo/contact-go/internal/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes() *gin.Engine {
	r := gin.Default()

	//middlewares
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "X-API-KEY"},
	}))
	r.Use(middlewares.RequireAPIKey())

	v1 := r.Group("/api/v1")
	{
		v1.POST("/contact", handlers.Contact)
	}

	return r
}