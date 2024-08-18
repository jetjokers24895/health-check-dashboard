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

type HandlerApi struct {
	_service services.IService
}

func NewHandlerApi(db *gorm.DB) *HandlerApi {
	return &HandlerApi{
		_service: services.NewServices(db, services.NewCronJobManager()),
	}
}

func (h *HandlerApi) NewService(c *fiber.Ctx) error {
	input := new(dtos.Service)
	if err := c.BodyParser(input); err != nil {
		c.JSON(dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "",
			Data:    nil,
		})
	}

	if input.URL == "" || input.Name == "" {
		return c.JSON(dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "URL or Name must be specified",
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

func (h *HandlerApi) UpdateService(c *fiber.Ctx) error {
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

func (h *HandlerApi) DeleteService(c *fiber.Ctx) error {
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

func (h *HandlerApi) GetServices(c *fiber.Ctx) error {
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
