package controllers

import (
	"fmt"
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

func (tc *TodoController) GetHello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}

func (tc *TodoController) GetTodos(c echo.Context) error {
	todos, err := tc.Service.GetAllTodos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    todos,
		"message": "Successfully",
	})
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

func (tc *TodoController) UpdateTodoByID(c echo.Context) error {
	id := c.Param("id")

	var input models.Todo
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	updatedTodo, err := tc.Service.UpdateTodoByID(id, input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, updatedTodo)
}

func (tc *TodoController) DeleteTodoByID(c echo.Context) error {
	id := c.Param("id")

	err := tc.Service.DeleteTodoByID(id)
	if err != nil {
		if err.Error() == fmt.Sprintf("todo with id %s not found", id) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Todo deleted"})
}
