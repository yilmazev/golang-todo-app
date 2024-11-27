package controllers

import (
	"golang-todo-app/internal/domain/models"
	"golang-todo-app/internal/domain/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TodoController struct {
	Service *services.TodoService
}

func NewTodoController(service *services.TodoService) *TodoController {
	return &TodoController{Service: service}
}

func (tc *TodoController) GetTodos(c echo.Context) error {
	todos, err := tc.Service.GetAllTodos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, todos)
}

func (tc *TodoController) CreateTodo(c echo.Context) error {
	var input models.Todo
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	todo, err := tc.Service.CreateTodo(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, todo)
}
