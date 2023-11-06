package main

import (
	// "encoding/json"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/moshriguez/personalGoodReads/config"
	"github.com/moshriguez/personalGoodReads/models"
	"github.com/moshriguez/personalGoodReads/routes"
)

var library []models.Book

func main() {

	app := fiber.New()
	_ = config.ConnectToDB(context.Background())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	})

	routes.BookRoutes(app)

	app.Listen(":3000")

	// // init router
	// router := mux.NewRouter()

	// // endpoints
	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Welcome to my library.")
	// }).Methods("GET")
	// if err := http.ListenAndServe(":3000", router); err != nil {
	// 	log.Fatal(err)
	// }
}
