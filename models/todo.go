package models

import (
	"time"

	"gorm.io/gorm"
)

// Todo модель задачи (аналог SQLAlchemy модели)
type Todo struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"not null;size:200" json:"title" binding:"required"`
	Description string         `gorm:"type:text" json:"description"`
	Completed   bool           `gorm:"default:false" json:"completed"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// TodoCreate схема для создания задачи
type TodoCreate struct {
	Title       string `json:"title" binding:"required,min=1,max=200"`
	Description string `json:"description"`
}

// TodoUpdate схема для обновления задачи
type TodoUpdate struct {
	Title       *string `json:"title" binding:"omitempty,min=1,max=200"`
	Description *string `json:"description"`
	Completed   *bool   `json:"completed"`
}

// TodoResponse схема ответа (если нужна кастомизация)
type TodoResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ToResponse конвертирует Todo в TodoResponse
func (t *Todo) ToResponse() TodoResponse {
	return TodoResponse{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Completed:   t.Completed,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}
