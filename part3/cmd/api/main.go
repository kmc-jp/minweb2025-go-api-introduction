package main

import (
	"part3/internal/handler"
	"part3/internal/middleware"
	"part3/internal/model"
	"part3/internal/repository"
	"part3/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize GORM with SQLite
	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Task{}, &model.Schedule{}, &model.User{})

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
	r.POST("/tasks", taskHandler.CreateTask)
	r.GET("/tasks/:id", taskHandler.GetTask)
	r.PUT("/tasks/:id", taskHandler.UpdateTask)
	r.DELETE("/tasks/:id", taskHandler.DeleteTask)
	r.GET("/tasks", taskHandler.ListTasks)

	// Start the server
	r.Run(":8080")
}
