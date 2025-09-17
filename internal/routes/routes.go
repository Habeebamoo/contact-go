package routes

import (
	"time"

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
		AllowMethods: []string{"POST", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "X-API-KEY"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))
	r.Use(middlewares.RequireAPIKey())

	r.POST("/api/v1/contact", handlers.Contact)

	return r
}