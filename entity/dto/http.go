package dto

import "github.com/gofiber/fiber/v2"

type HandlerResponse struct {
	Status  string       `json:"status"`
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    *interface{} `json:"data,omitempty"`
}

func ResponseSuccessStruct(ctx *fiber.Ctx, message string, code int, data interface{}) error {
	resp := &HandlerResponse{
		Status:  "Success",
		Code:    code,
		Message: message,
		Data:    &data,
	}

	return ctx.Status(code).JSON(resp)
}

func ResponseFailedStruct(ctx *fiber.Ctx, message string, code int) error {
	resp := &HandlerResponse{
		Status:  "Failed",
		Code:    code,
		Message: message,
	}

	return ctx.Status(code).JSON(resp)
}
