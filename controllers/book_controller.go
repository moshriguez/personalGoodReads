package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/moshriguez/personalGoodReads/config"
	"github.com/moshriguez/personalGoodReads/db"
	"github.com/moshriguez/personalGoodReads/models"
	"github.com/moshriguez/personalGoodReads/responses"
	"go.mongodb.org/mongo-driver/mongo"
)

var BookCollection *mongo.Collection = config.GetCollections(config.DB, "books")

func GetBooks(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	books, err := db.GetBooks(ctx)
	if err != nil {
		return responses.InternalServerError(c, err.Error())
	}

	return responses.OK(c, books)
}

func AddBook(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var book models.Book

	// validate the request body
	if err := c.BodyParser(&book); err != nil {
		return responses.BadRequest(c, fmt.Sprintf("Error parsing request: %s", err.Error()))
	}

	addedBook, err := db.AddBook(ctx, &book)
	if err != nil {
		return responses.InternalServerError(c, err.Error())
	}

	return responses.Created(c, addedBook)
}

func UpdateBook(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var book models.Book
	bookID := c.Params("bookID")

	// validate the request body
	if err := c.BodyParser(&book); err != nil {
		return responses.BadRequest(c, fmt.Sprintf("Error parsing request: %s", err.Error()))
	}

	updatedBook, err := db.UpdateBook(ctx, bookID, &book)
	if err != nil {
		return responses.InternalServerError(c, err.Error())
	}

	return responses.OK(c, updatedBook)
}

func DeleteBook(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	bookID := c.Params("bookID")

	err := db.DeleteBook(ctx, bookID)
	if err != nil {
		return responses.InternalServerError(c, err.Error())
	}

	return responses.OK(c, "Book successfully deleted!!!")
}

func GetBook(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	bookID := c.Params("bookID")

	book, err := db.GetBook(ctx, bookID)
	if err != nil {
		return responses.InternalServerError(c, err.Error())
	}

	return responses.OK(c, book)
}
