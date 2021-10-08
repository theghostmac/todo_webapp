package models

import  (
	"fullstack-go/todo_webapp/database"
	"github.com/gofiber/fiber/v2"
)

type ToDo struct {
	ID uint `gorm:"primarykey" json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func GetTodos(c *fiber.Ctx) error {
	db := database.DBConn
	var todos []ToDo
	db.Find(&todos)
	return c.JSON(&todos)
}