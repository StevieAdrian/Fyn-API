package routes

import (
	"github.com/StevieAdrian/Fyn-API/auth-service/internal/auth/delivery/http"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, h *http.Handler) {
	router.POST("/signup", h.Signup())
	router.POST("/login", h.Login())

	protected := router.Group("/")
	protected.Use(http.Authenticate())
	{
		protected.GET("/users", h.GetUsers())
		protected.GET("/user/:id", h.GetUser())
	}
}
