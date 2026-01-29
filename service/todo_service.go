package service

import (
	"errors"

	"github.com/alex6712/learning-golang/models"
	"github.com/alex6712/learning-golang/repository"

	"gorm.io/gorm"
)

// TodoService содержит бизнес-логику
type TodoService struct {
	repo *repository.TodoRepository
}

// NewTodoService создает новый сервис
func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

// CreateTodo создает новую задачу
func (s *TodoService) CreateTodo(data models.TodoCreate) (*models.Todo, error) {
	todo := &models.Todo{
		Title:       data.Title,
		Description: data.Description,
		Completed:   false,
	}

	if err := s.repo.Create(todo); err != nil {
		return nil, err
	}

	return todo, nil
}

// GetAllTodos возвращает список задач с пагинацией
func (s *TodoService) GetAllTodos(skip, limit int) ([]models.Todo, int64, error) {
	todos, err := s.repo.GetAll(skip, limit)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.Count()
	if err != nil {
		return nil, 0, err
	}

	return todos, total, nil
}

// GetTodoByID возвращает задачу по ID
func (s *TodoService) GetTodoByID(id uint) (*models.Todo, error) {
	todo, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("задача не найдена")
		}
		return nil, err
	}
	return todo, nil
}

// UpdateTodo обновляет задачу
func (s *TodoService) UpdateTodo(id uint, data models.TodoUpdate) (*models.Todo, error) {
	todo, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("задача не найдена")
		}
		return nil, err
	}

	// Обновляем только переданные поля
	if data.Title != nil {
		todo.Title = *data.Title
	}
	if data.Description != nil {
		todo.Description = *data.Description
	}
	if data.Completed != nil {
		todo.Completed = *data.Completed
	}

	if err := s.repo.Update(todo); err != nil {
		return nil, err
	}

	return todo, nil
}

// DeleteTodo удаляет задачу
func (s *TodoService) DeleteTodo(id uint) error {
	// Проверяем существование
	_, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("задача не найдена")
		}
		return err
	}

	return s.repo.Delete(id)
}
