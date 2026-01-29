package repository

import (
	"github.com/alex6712/learning-golang/models"

	"gorm.io/gorm"
)

// TodoRepository работает с базой данных (аналог SQLAlchemy session)
type TodoRepository struct {
	db *gorm.DB
}

// NewTodoRepository создает новый репозиторий
func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

// Create создает новую задачу
func (r *TodoRepository) Create(todo *models.Todo) error {
	return r.db.Create(todo).Error
}

// GetAll возвращает все задачи
func (r *TodoRepository) GetAll(skip, limit int) ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Offset(skip).Limit(limit).Order("created_at DESC").Find(&todos).Error
	return todos, err
}

// GetByID возвращает задачу по ID
func (r *TodoRepository) GetByID(id uint) (*models.Todo, error) {
	var todo models.Todo
	err := r.db.First(&todo, id).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

// Update обновляет задачу
func (r *TodoRepository) Update(todo *models.Todo) error {
	return r.db.Save(todo).Error
}

// Delete удаляет задачу (soft delete)
func (r *TodoRepository) Delete(id uint) error {
	return r.db.Delete(&models.Todo{}, id).Error
}

// Count возвращает общее количество задач
func (r *TodoRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&models.Todo{}).Count(&count).Error
	return count, err
}
