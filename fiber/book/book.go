package book

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"../database"
)

type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	return c.JSON(book)
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	return c.JSON(books)
}

func AddBook(c *fiber.Ctx) error {
	db := database.DBConn
	var book Book
	book.Title = "1984"
	book.Author = "George Orwell"
	book.Rating = 5
	db.Create(&book)
	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
        return c.Status(500).SendString("No Book Found with ID")
	}
	db.Delete(&book)
	return c.SendString("Book Successfully deleted")
}