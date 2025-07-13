package bootstrap

import (
	"log"
	"os"

	"github.com/StevieAdrian/Fyn-API/auth-service/config"
	"github.com/StevieAdrian/Fyn-API/auth-service/internal/auth/delivery/http"
	"github.com/StevieAdrian/Fyn-API/auth-service/internal/auth/repository/gorm"
	"github.com/StevieAdrian/Fyn-API/auth-service/internal/auth/usecase"
	"github.com/StevieAdrian/Fyn-API/auth-service/pkg/token"
	"github.com/StevieAdrian/Fyn-API/auth-service/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type App struct {
	engine *gin.Engine
}

func InitializeApp() *App {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()
	db := config.GetDB()
	db.AutoMigrate(&gorm.UserModel{})

	jwtSecret := os.Getenv("JWT_SECRET")
	token.SetJWTKey(jwtSecret)

	userRepo := gorm.NewUserRepository(db)
	userUC := usecase.NewUserUsecase(userRepo)
	handler := http.NewHandler(userUC)

	engine := gin.Default()
	routes.SetupRoutes(engine, handler)

	return &App{engine: engine}
}

func (a *App) Run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server is running on port:", port)
	return a.engine.Run(":" + port)
}
