package main

import (
	"golang-todo-app/config"
	"golang-todo-app/internal/http"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnv()

	db := config.ConnectDB()
	defer db.Close()

	e := echo.New()
	http.RegisterRoutes(e, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Starting server on " + port)
	e.Logger.Fatal(e.Start(":" + port))
}
