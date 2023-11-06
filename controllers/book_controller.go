package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/moshriguez/personalGoodReads/config"
	"github.com/moshriguez/personalGoodReads/models"
	"github.com/moshriguez/personalGoodReads/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var BookCollection *mongo.Collection = config.GetCollections(config.DB, "books")

func GetBooks(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var books []models.Book

	results, err := BookCollection.Find(ctx, bson.M{})
	if err != nil {
		return responses.InternalServerError(c, err.Error())
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var book models.Book
		if err := results.Decode(&book); err != nil {
			return responses.InternalServerError(c, err.Error())
		}

		books = append(books, book)
	}

	return responses.OK(c, books)
}

func AddBook(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var book models.Book

	objID := primitive.NewObjectID()

	// validate the request body
	if err := c.BodyParser(&book); err != nil {
		return responses.BadRequest(c, fmt.Sprintf("Error parsing request: %s", err.Error()))
	}

	newBook := models.Book{
		ID:         objID,
		Title:      book.Title,
		Author:     book.Author,
		Read:       book.Read,
		ReadOn:     book.ReadOn,
		LostInMove: book.LostInMove,
		Deleted:    false,
	}

	result, err := BookCollection.InsertOne(ctx, newBook)
	if err != nil {
		return responses.InternalServerError(c, err.Error())
	}

	return responses.Created(c, result)
}

func UpdateBook(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var book models.Book
	bookID := c.Params("bookID")

	objID, _ := primitive.ObjectIDFromHex(bookID)

	fmt.Printf("bookID: %s | objectID: %v", bookID, objID)

	// validate the request body
	if err := c.BodyParser(&book); err != nil {
		return responses.BadRequest(c, fmt.Sprintf("Error parsing request: %s", err.Error()))
	}

	update := bson.M{
		"title":        book.Title,
		"author":       book.Author,
		"read":         book.Read,
		"read_on":      book.ReadOn,
		"lost_in_move": book.LostInMove,
		"deleted":      book.Deleted,
	}

	result, err := BookCollection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": update})
	if result.ModifiedCount == 0 {
		return responses.NotFound(c, "could not find book to update")
	}
	if err != nil {
		return responses.InternalServerError(c, err.Error())
	}
	// get updated book details
	if result.MatchedCount != 1 {
		return responses.InternalServerError(c, "error: bookID matched > 1 book")
	}
	var updatedBook models.Book
	err = BookCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&updatedBook)
	if err != nil {
		return responses.InternalServerError(c, err.Error())
	}

	return responses.OK(c, updatedBook)
}

func DeleteBook(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	bookID := c.Params("bookID")

	objID, _ := primitive.ObjectIDFromHex(bookID)

	result, err := BookCollection.DeleteOne(ctx, bson.M{"id": objID})
	if err != nil {
		return responses.InternalServerError(c, err.Error())
	}

	if result.DeletedCount < 1 {
		return responses.NotFound(c, "Book with specified ID not found!")
	}

	return responses.OK(c, "Book successfully deleted!!!")
}

func GetBook(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var book models.Book
	bookID := c.Params("bookID")

	objID, _ := primitive.ObjectIDFromHex(bookID)

	err := BookCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&book)
	if err != nil {
		return responses.InternalServerError(c, err.Error())
	}

	return responses.OK(c, book)
}
