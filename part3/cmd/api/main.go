package main

import (
	"part3/internal/handler"
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
	db.AutoMigrate(&model.Task{}, &model.Schedule{})

	// Initialize repositories
	taskRepo := repository.NewTaskRepository(db)
	scheduleRepo := repository.NewScheduleRepository(db)

	// Initialize services
	taskService := service.NewTaskService(taskRepo)
	scheduleService := service.NewScheduleService(scheduleRepo, taskRepo)

	// Initialize handlers
	taskHandler := handler.NewTaskHandler(taskService)
	scheduleHandler := handler.NewScheduleHandler(scheduleService)

	// Set up Gin router
	r := gin.Default()

	// Task routes
	r.POST("/tasks", taskHandler.CreateTask)
	r.GET("/tasks/:id", taskHandler.GetTask)
	r.PUT("/tasks/:id", taskHandler.UpdateTask)
	r.DELETE("/tasks/:id", taskHandler.DeleteTask)
	r.GET("/tasks", taskHandler.ListTasks)

	// Schedule routes
	r.POST("/schedules", scheduleHandler.CreateSchedule)
	r.GET("/schedules/:id", scheduleHandler.GetSchedule)
	r.GET("/tasks/:taskId/schedules", scheduleHandler.GetSchedulesByTask)
	r.PUT("/schedules/:id", scheduleHandler.UpdateSchedule)
	r.DELETE("/schedules/:id", scheduleHandler.DeleteSchedule)
	r.GET("/schedules", scheduleHandler.ListSchedules)

	// Start the server
	r.Run(":8080")
}
