package repository

import (
	"context"
	"fmt"
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
	rows, err := r.DB.Query(ctx, "SELECT id, title, description, completed FROM todos ORDER BY id DESC")
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

func (r *TodoRepository) UpdateTodoByID(ctx context.Context, id string, input models.Todo) (models.Todo, error) {
	query := `
		UPDATE todos
		SET title = COALESCE($1, title),
		    description = COALESCE($2, description),
		    completed = COALESCE($3, completed)
		WHERE id = $4
		RETURNING id, title, description, completed
	`

	var updatedTodo models.Todo
	err := r.DB.QueryRow(ctx, query, input.Title, input.Description, input.Completed, id).
		Scan(&updatedTodo.ID, &updatedTodo.Title, &updatedTodo.Description, &updatedTodo.Completed)

	return updatedTodo, err
}

func (r *TodoRepository) DeleteTodoByID(ctx context.Context, id string) error {
	commandTag, err := r.DB.Exec(ctx, "DELETE FROM todos WHERE id=$1", id)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("todo with id %s not found", id)
	}

	return nil
}
