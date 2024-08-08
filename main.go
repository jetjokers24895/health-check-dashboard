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
	_handlerApi := handlers.NewHandlerApi(cursor)
	_handler := handlers.NewHandler(cursor)

	app.Get("/", _handler.Home)
	app.Get("/new-service", _handler.NewService)
	app.Static("/static", "./views/static")

	app.Post("/api/services", _handlerApi.NewService)
	app.Put("/api/services/:id", _handlerApi.UpdateService)
	app.Delete("/api/services/:id", _handlerApi.DeleteService)
	app.Get("/api/services", _handlerApi.GetServices)

	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.JSON(dtos.Response{
			Status:  http.StatusOK,
			Message: "I'm ok, bro!",
			Data:    nil,
		})
	})
	log.Fatal(app.Listen(":3000"))
}
