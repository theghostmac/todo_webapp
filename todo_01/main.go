package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"fullstack-go/todo_webapp/models"
	"fullstack-go/todo_webapp/database"
	"gorm.io/gorm"
)

func message(c *fiber.Ctx) error {
	return c.SendString("Building a ToDo Web App with Golang, GORM, Fiber and ReactJS!")
}

func initDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=fu11st4ck dbname=todowebapp port=54321"
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database...")
	} else {
		fmt.Println("Database connected successfully!")
		database.DBConn.AutoMigrate(models.ToDo{})
		fmt.Println("Migrated DB...")
	}
}

func setupRoutes(app *fiber.App) {
	app.Get("/todos", models.GetTodos)
}

func main() {
	app := fiber.New()
	initDatabase()
	app.Get("/", message)
	setupRoutes(app)
	app.Listen(":8000")
}