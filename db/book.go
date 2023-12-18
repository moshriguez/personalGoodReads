package db

import (
	"context"
	"fmt"

	"github.com/moshriguez/personalGoodReads/config"
	"github.com/moshriguez/personalGoodReads/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var BookCollection *mongo.Collection = config.GetCollections(config.DB, "books")

func GetBooks(ctx context.Context) (*[]models.Book, error) {
	var books []models.Book

	results, err := BookCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var book models.Book
		if err := results.Decode(&book); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return &books, nil
}

func AddBook(ctx context.Context, book *models.Book) (*models.Book, error) {
	objID := primitive.NewObjectID()

	newBook := models.Book{
		ID:         objID,
		Title:      book.Title,
		Author:     book.Author,
		Read:       book.Read,
		ReadOn:     book.ReadOn,
		LostInMove: book.LostInMove,
		Deleted:    false,
	}

	_, err := BookCollection.InsertOne(ctx, newBook)
	if err != nil {
		return nil, err
	}

	return &newBook, nil
}

func UpdateBook(ctx context.Context, bookID string, book *models.Book) (*models.Book, error) {
	objID, _ := primitive.ObjectIDFromHex(bookID)

	update := buildUpdateObj(book)

	result, err := BookCollection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": update})
	if result.ModifiedCount == 0 {
		return nil, fmt.Errorf("not found error: could not find book to update")
	}
	if err != nil {
		return nil, err
	}
	// get updated book details
	if result.MatchedCount != 1 {
		return nil, fmt.Errorf("error: bookID matched > 1 book")
	}
	var updatedBook models.Book
	err = BookCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&updatedBook)
	if err != nil {
		return nil, err
	}

	return &updatedBook, nil
}

func buildUpdateObj(book *models.Book) bson.D {
	var updateObj bson.D

	if book.Title != "" {
		updateObj = append(updateObj, bson.E{Key: "title", Value: book.Title})
	}
	if book.Author != "" {
		updateObj = append(updateObj, bson.E{Key: "author", Value: book.Author})
	}
	if book.Read {
		updateObj = append(updateObj, bson.E{Key: "read", Value: book.Read})
	}
	if !book.ReadOn.IsZero() {
		updateObj = append(updateObj, bson.E{Key: "read_on", Value: book.ReadOn})
	}
	if book.LostInMove {
		updateObj = append(updateObj, bson.E{Key: "lost_in_move", Value: book.LostInMove})
	}
	if book.Deleted {
		updateObj = append(updateObj, bson.E{Key: "deleted", Value: book.Deleted})
	}

	return updateObj
}

func DeleteBook(ctx context.Context, bookID string) error {
	objID, _ := primitive.ObjectIDFromHex(bookID)

	result, err := BookCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	if result.DeletedCount < 1 {
		return fmt.Errorf("not found error: book with specified ID not found")
	}

	return nil
}

func GetBook(ctx context.Context, bookID string) (*models.Book, error) {
	var book models.Book

	objID, _ := primitive.ObjectIDFromHex(bookID)

	err := BookCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&book)
	if err != nil {
		return nil, err
	}

	return &book, nil
}
