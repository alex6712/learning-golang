package routes

import (
	"github.com/alex6712/learning-golang/handlers"
	"github.com/alex6712/learning-golang/models"
	"github.com/alex6712/learning-golang/repository"
	"github.com/alex6712/learning-golang/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes регистрирует все роуты приложения
func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	// Миграция моделей (автосоздание таблиц)
	db.AutoMigrate(&models.Todo{})

	// Инициализация слоев
	todoRepo := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handlers.NewTodoHandler(todoService)

	// API группа
	api := router.Group("/api/v1")
	{
		// Роуты для задач
		todos := api.Group("/todos")
		{
			todos.POST("", todoHandler.CreateTodo)
			todos.GET("", todoHandler.GetAllTodos)
			todos.GET("/:id", todoHandler.GetTodoByID)
			todos.PUT("/:id", todoHandler.UpdateTodo)
			todos.DELETE("/:id", todoHandler.DeleteTodo)
		}

		// Health check endpoint
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "API работает",
			})
		})
	}
}
