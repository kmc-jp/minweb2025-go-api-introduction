package main

import (
	"fmt"
	"log"
	"os"
	"part3/internal/handler"
	"part3/internal/middleware"
	"part3/internal/model"
	"part3/internal/repository"
	"part3/internal/service"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// 環境変数からDB接続情報を取得
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "user")
	dbPassword := getEnv("DB_PASSWORD", "password")
	dbName := getEnv("DB_NAME", "app_db")

	// PostgreSQL接続文字列の構築
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// データベース接続（リトライ機能付き）
	var db *gorm.DB
	var err error
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Successfully connected to database")
			break
		}
		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(time.Second * 2)
	}
	if err != nil {
		log.Fatal("failed to connect database after retries:", err)
	}

	// Migrate the schema
	if err := db.AutoMigrate(&model.Task{}, &model.Schedule{}, &model.User{}); err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	// Initialize repositories
	taskRepo := repository.NewTaskRepository(db)
	scheduleRepo := repository.NewScheduleRepository(db)
	// Initialize services
	taskService := service.NewTaskService(taskRepo)
	scheduleService := service.NewScheduleService(scheduleRepo, taskRepo)
	authService := service.NewAuthService(db)

	// Initialize handlers
	taskHandler := handler.NewTaskHandler(taskService)
	scheduleHandler := handler.NewScheduleHandler(scheduleService)
	authHandler := handler.NewAuthHandler(authService)

	// Set up Gin router
	r := gin.Default()

	// Auth routes
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	// Task routes
	r.POST("/tasks", taskHandler.CreateTask)
	r.GET("/tasks/:id", taskHandler.GetTask)
	r.PUT("/tasks/:id", taskHandler.UpdateTask)
	r.DELETE("/tasks/:id", taskHandler.DeleteTask)
	r.GET("/tasks", taskHandler.ListTasks)

	// Schedule routes (認証必須)
	authGroup := r.Group("/schedules")
	authGroup.Use(middleware.AuthMiddleware())
	{
		authGroup.POST("/", scheduleHandler.CreateSchedule)
		authGroup.GET("/:id", scheduleHandler.GetSchedule)
		authGroup.GET("/tasks/:taskId/schedules", scheduleHandler.GetSchedulesByTask)
		authGroup.PUT("/:id", scheduleHandler.UpdateSchedule)
		authGroup.DELETE("/:id", scheduleHandler.DeleteSchedule)
		authGroup.GET("/", scheduleHandler.ListSchedules)
	}

	// Start the server
	log.Println("Starting server on :8080")
	r.Run(":8080")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
