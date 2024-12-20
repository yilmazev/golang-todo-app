package services

import (
	"context"
	"golang-todo-app/internal/domain/models"
	"golang-todo-app/internal/repository"
)

type TodoService struct {
	Repo *repository.TodoRepository
}

func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{Repo: repo}
}

func (ts *TodoService) GetAllTodos() ([]models.Todo, error) {
	return ts.Repo.GetTodos(context.Background())
}

func (ts *TodoService) CreateTodo(input models.Todo) (models.Todo, error) {
	return ts.Repo.CreateTodo(context.Background(), input)
}

func (ts *TodoService) UpdateTodoByID(id string, input models.Todo) (models.Todo, error) {
	return ts.Repo.UpdateTodoByID(context.Background(), id, input)
}

func (ts *TodoService) DeleteTodoByID(id string) error {
	return ts.Repo.DeleteTodoByID(context.Background(), id)
}
