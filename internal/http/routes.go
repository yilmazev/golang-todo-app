package http

import (
	"golang-todo-app/internal/domain/services"
	"golang-todo-app/internal/http/controllers"
	"golang-todo-app/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, db *pgxpool.Pool) {
	todoRepo := repository.NewTodoRepository(db)
	todoService := services.NewTodoService(todoRepo)
	todoController := controllers.NewTodoController(todoService)

	e.GET("/", todoController.GetHello)
	e.GET("/todos", todoController.GetTodos)
	e.POST("/create-todo", todoController.CreateTodo)
	e.DELETE("/todo/:id", todoController.DeleteTodoByID)
	e.PATCH("/todo/:id", todoController.UpdateTodoByID)
}
