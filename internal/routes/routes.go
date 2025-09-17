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

	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(200)
	})

	v1 := r.Group("/api/v1")
	{
		v1.POST("/contact", handlers.Contact)
	}

	return r
}