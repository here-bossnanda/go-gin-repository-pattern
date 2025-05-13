package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	_ "github.com/sekolahmu/boilerplate-go/docs"
	"github.com/sekolahmu/boilerplate-go/internal/config"
	"github.com/sekolahmu/boilerplate-go/internal/delivery/http"
	"github.com/sekolahmu/boilerplate-go/internal/repository"
	"github.com/sekolahmu/boilerplate-go/internal/usecase"
	"github.com/swaggo/gin-swagger"
)

// @title Boilerplate Go API
// @version 1.0
// @description This is a boilerplate Go API using clean architecture
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Load configuration
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize database connection
	dbConnStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)

	db, err := sql.Open(cfg.Database.Driver, dbConnStr)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Initialize repository
	userRepo := repository.NewUserRepository(db)

	// Initialize usecase
	userUseCase := usecase.NewUserUseCase(userRepo)

	// Initialize HTTP handler
	userHandler := http.NewUserHandler(userUseCase)

	// Initialize Gin router
	router := gin.Default()

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	v1 := router.Group("/api/v1")
	{
		userHandler.RegisterRoutes(v1)
	}

	// Start server
	serverAddr := fmt.Sprintf(":%s", cfg.Server.Port)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
