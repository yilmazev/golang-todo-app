package main

import (
	"golang-todo-app/config"
	"golang-todo-app/internal/http"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnv()
	db := config.ConnectDB()
	defer db.Close()

	e := echo.New()
	http.RegisterRoutes(e, db)

	log.Println("Starting server on port 8080...")
	e.Logger.Fatal(e.Start(":8080"))
}
