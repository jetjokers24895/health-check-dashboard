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
		_service: services.NewServices(db),
	}
}

func (h *Handler) NewService(c *fiber.Ctx) error {
	input := new(dtos.Service)
	if err := c.BodyParser(input); err != nil {
		c.JSON(dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "",
			Data:    nil,
		})
	}

	if input.Command == "" || input.Name == "" {
		return c.JSON(dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Command or Name must be specified",
			Data:    nil,
		})
	}

	if err := h._service.AddService(input); err != nil {
		log.Println(err)
		return c.JSON(dtos.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(dtos.Response{Status: http.StatusOK, Data: nil, Message: "OK"})
}

func (h *Handler) UpdateService(c *fiber.Ctx) error {
	input := new(dtos.Service)

	if err := c.ParamsParser(input); err != nil {
		c.JSON(dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "",
			Data:    nil,
		})
	}

	if err := c.BodyParser(input); err != nil {
		c.JSON(dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "",
			Data:    nil,
		})
	}

	if input.ID == 0 {
		return c.JSON(dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    nil,
		})
	}

	if err := h._service.UpdateService(input); err != nil {
		log.Println(err)
		return c.JSON(dtos.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(dtos.Response{Status: http.StatusOK, Data: nil, Message: "OK"})
}

func (h *Handler) DeleteService(c *fiber.Ctx) error {
	input := new(dtos.Service)

	if err := c.ParamsParser(input); err != nil {
		c.JSON(dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "",
			Data:    nil,
		})
	}

	if input.ID == 0 {
		return c.JSON(dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    nil,
		})
	}

	if err := h._service.DeleteService(input.ID); err != nil {
		log.Println(err)
		return c.JSON(dtos.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(dtos.Response{Status: http.StatusOK, Data: nil, Message: "OK"})
}

func (h *Handler) GetServices(c *fiber.Ctx) error {
	rawStatus := c.Query("status")
	var status = 0
	var err error
	if rawStatus != "" {
		status, err = strconv.Atoi(rawStatus)
	}
	if err != nil {
		c.JSON(dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Status must be a number 1 or 2",
			Data:    nil,
		})
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

	return c.JSON(dtos.Response{Status: http.StatusOK, Data: rs, Message: "OK"})
}

func (h *Handler) Home(c *fiber.Ctx) error {
	rawStatus := c.Query("status")
	var status = 0
	var err error
	if rawStatus != "" {
		status, err = strconv.Atoi(rawStatus)
	}
	if err != nil || (status != 1 && status != 2) {
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
