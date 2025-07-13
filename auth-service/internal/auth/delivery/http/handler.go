package http

import (
	"context"
	"net/http"
	"time"

	"github.com/StevieAdrian/Fyn-API/auth-service/internal/auth/dto"
	"github.com/StevieAdrian/Fyn-API/auth-service/internal/auth/usecase"
	"github.com/StevieAdrian/Fyn-API/auth-service/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	UserUC usecase.UserUsecase
}

func NewHandler(userUC usecase.UserUsecase) *Handler {
	return &Handler{
		UserUC: userUC,
	}
}

var validate = validator.New()

func (h *Handler) Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.SignupRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		if err := validate.Struct(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		err := h.UserUC.Signup(c.Request.Context(), req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	}
}

func (h *Handler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.LoginRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, token, refresh, err := h.UserUC.Login(c.Request.Context(), req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user":          user,
			"token":         token,
			"refresh_token": refresh,
		})
	}
}

func (h *Handler) GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		claimsAny, exists := c.Get("claims")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		claims, ok := claimsAny.(*token.Claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid claims"})
			return
		}

		if claims.Role != "ADMIN" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		_, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		users, err := h.UserUC.GetAll(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func (h *Handler) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// implement nanti kalau udah siap detail-nya
		c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet"})
	}
}
