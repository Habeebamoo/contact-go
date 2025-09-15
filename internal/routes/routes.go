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
		AllowOrigins: []string{"https://habeebamoo.netlify.app", "https://acitglobal.vercel.app"},
		AllowMethods: []string{"POST"},
		AllowHeaders: []string{"Content-Type", "X-API-KEY", "Origin"},
	}))
	r.Use(middlewares.RequireAPIKey())

	r.POST("/api/v1/contact", handlers.Contact)

	return r
}