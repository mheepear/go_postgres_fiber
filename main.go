package main

import (
	"fmt"
	"todolist/database"
	"todolist/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func helloWord(c *fiber.Ctx) error {
	return c.SendString("Hello world")
}

func initDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=1234 dbname=goTodo port=5433"
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connected!")
	database.DBConn.AutoMigrate(&models.Todo{})
	fmt.Println("Migrated DB")
}

func setupRoutes(app *fiber.App) {
	app.Get("/todos", models.GetTodos)
	app.Post("/todos", models.CreateTodo)

}

func main() {
	app := fiber.New()
	initDatabase()
	app.Get("/", helloWord)
	setupRoutes(app)
	app.Listen(":8000")
}
