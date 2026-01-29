package handlers

import (
	"net/http"
	"strconv"

	"github.com/alex6712/learning-golang/models"
	"github.com/alex6712/learning-golang/service"

	"github.com/gin-gonic/gin"
)

// TodoHandler обрабатывает HTTP запросы (аналог FastAPI роутов)
type TodoHandler struct {
	service *service.TodoService
}

// NewTodoHandler создает новый хэндлер
func NewTodoHandler(service *service.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

// CreateTodo создает новую задачу
// @Summary Создать задачу
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body models.TodoCreate true "Данные задачи"
// @Success 201 {object} models.Todo
// @Router /todos [post]
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var data models.TodoCreate

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := h.service.CreateTodo(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

// GetAllTodos возвращает список всех задач
// @Summary Получить все задачи
// @Tags todos
// @Produce json
// @Param skip query int false "Пропустить записей" default(0)
// @Param limit query int false "Лимит записей" default(100)
// @Success 200 {object} map[string]interface{}
// @Router /todos [get]
func (h *TodoHandler) GetAllTodos(c *gin.Context) {
	skip, _ := strconv.Atoi(c.DefaultQuery("skip", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))

	todos, total, err := h.service.GetAllTodos(skip, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": todos,
		"total": total,
		"skip":  skip,
		"limit": limit,
	})
}

// GetTodoByID возвращает задачу по ID
// @Summary Получить задачу по ID
// @Tags todos
// @Produce json
// @Param id path int true "ID задачи"
// @Success 200 {object} models.Todo
// @Router /todos/{id} [get]
func (h *TodoHandler) GetTodoByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	todo, err := h.service.GetTodoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// UpdateTodo обновляет задачу
// @Summary Обновить задачу
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "ID задачи"
// @Param todo body models.TodoUpdate true "Обновленные данные"
// @Success 200 {object} models.Todo
// @Router /todos/{id} [put]
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	var data models.TodoUpdate
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := h.service.UpdateTodo(uint(id), data)
	if err != nil {
		if err.Error() == "задача не найдена" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// DeleteTodo удаляет задачу
// @Summary Удалить задачу
// @Tags todos
// @Param id path int true "ID задачи"
// @Success 204
// @Router /todos/{id} [delete]
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	if err := h.service.DeleteTodo(uint(id)); err != nil {
		if err.Error() == "задача не найдена" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
