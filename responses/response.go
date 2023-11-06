package responses

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type BookResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}

func ErrorResponse(c *fiber.Ctx, httpCode int, data interface{}) error {
	return c.Status(httpCode).JSON(
		BookResponse{
			Status:  httpCode,
			Message: "error",
			Data:    &fiber.Map{"data": data},
		},
	)
}

func BadRequest(c *fiber.Ctx, data interface{}) error {
	return ErrorResponse(c, http.StatusBadRequest, data)
}

func InternalServerError(c *fiber.Ctx, data interface{}) error {
	return ErrorResponse(c, http.StatusInternalServerError, data)
}

func NotFound(c *fiber.Ctx, data interface{}) error {
	return ErrorResponse(c, http.StatusNotFound, data)
}

func SuccessResponse(c *fiber.Ctx, httpCode int, data interface{}) error {
	return c.Status(httpCode).JSON(
		BookResponse{
			Status:  httpCode,
			Message: "success",
			Data:    &fiber.Map{"data": data},
		},
	)
}

func OK(c *fiber.Ctx, data interface{}) error {
	return SuccessResponse(c, http.StatusOK, data)
}

func Created(c *fiber.Ctx, data interface{}) error {
	return SuccessResponse(c, http.StatusCreated, data)
}
