package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"./book"
	"./database"
)

func handleHome(c *fiber.Ctx) error {
	return c.SendString("Home page reached!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", handleHome)

	app.Get("/api/v1/books", book.GetBooks)
	app.Get("/api/v1/books/:id", book.GetBook)
	app.Post("/api/v1/books/:id", book.AddBook)
	app.Delete("/api/v1/books/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()
	

	setupRoutes(app)
	

	fmt.Println("Listening on port 3001...")
	fmt.Println(app.Listen(":3001"))
}