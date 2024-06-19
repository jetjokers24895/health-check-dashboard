package main

import (
	"log"
	"net/http"

	"app/db"
	"app/dtos"

	"app/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Create a new engine
	engine := html.New("./views", ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	cursor := db.ConnectDB()
	_handler := handlers.NewHandler(cursor)

	app.Get("/", _handler.Home)

	app.Post("/api/services", _handler.NewService)
	app.Put("/api/services/:id", _handler.UpdateService)
	app.Delete("/api/services/:id", _handler.DeleteService)
	app.Get("/api/services", _handler.GetServices)

	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.JSON(dtos.Response{
			Status:  http.StatusOK,
			Message: "I'm ok, bro!",
			Data:    nil,
		})
	})
	log.Fatal(app.Listen(":3000"))
}
