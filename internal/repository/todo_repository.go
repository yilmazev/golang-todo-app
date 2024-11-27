package repository

import (
	"context"
	"golang-todo-app/internal/domain/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TodoRepository struct {
	DB *pgxpool.Pool
}

func NewTodoRepository(db *pgxpool.Pool) *TodoRepository {
	return &TodoRepository{DB: db}
}

func (r *TodoRepository) GetTodos(ctx context.Context) ([]models.Todo, error) {
	rows, err := r.DB.Query(ctx, "SELECT id, title, description, completed FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *TodoRepository) CreateTodo(ctx context.Context, todo models.Todo) (models.Todo, error) {
	err := r.DB.QueryRow(
		ctx,
		"INSERT INTO todos (title, description, completed) VALUES ($1, $2, $3) RETURNING id",
		todo.Title, todo.Description, todo.Completed,
	).Scan(&todo.ID)
	return todo, err
}
