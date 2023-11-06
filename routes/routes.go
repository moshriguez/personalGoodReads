package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/moshriguez/personalGoodReads/controllers"
)

func BookRoutes(app *fiber.App) {
	// all routes interacting with book model
	app.Post("/books", controllers.AddBook)
	app.Get("/books/:bookID", controllers.GetBook)
	app.Put("/books/:bookID", controllers.UpdateBook)
	app.Delete("/books/:bookID", controllers.DeleteBook)
	app.Get("/books", controllers.GetBooks)
}
