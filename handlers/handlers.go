package handlers

import (
	"app/constants"
	"app/services"
	"net/http"
	"strconv"

	"app/dtos"

	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler struct {
	_service services.IService
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		_service: services.NewServices(db, nil),
	}
}

func (h *Handler) Home(c *fiber.Ctx) error {
	rawStatus := c.Query("status")
	var status = 0
	var err error
	if rawStatus != "" {
		status, err = strconv.Atoi(rawStatus)
	}
	if err != nil || (status != 1 && status != 2 && status != 0) {
		return c.Render("error", fiber.Map{
			"Title": "Status must be a number 1 or 2",
		}, "layouts/main")
	}

	rs, err := h._service.GetServices(constants.Status(status))
	if err != nil {
		log.Println(err)
		return c.JSON(dtos.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.Render("index", fiber.Map{
		"services": rs,
	}, "layouts/main")
}

func (h *Handler) NewService(c *fiber.Ctx) error {

	return c.Render("newService", fiber.Map{}, "layouts/main")
}
