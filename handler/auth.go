package handler

import (
	"github.com/abiyyu03/auth-service/entity/dto"
	"github.com/abiyyu03/auth-service/service"
	"github.com/gofiber/fiber/v2"
)

type AuthHandlerInterface interface {
	Login(ctx *fiber.Ctx) error
	RefreshToken(ctx *fiber.Ctx) error
}

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) AuthHandlerInterface {
	return &AuthHandler{
		service: service,
	}
}

func (a *AuthHandler) Login(ctx *fiber.Ctx) error {
	var authRequest *dto.AuthRequest

	if err := ctx.BodyParser(&authRequest); err != nil {
		return dto.ResponseFailedStruct(ctx, err.Error(), fiber.ErrBadRequest.Code)
	}

	userLogin, err := a.service.Login(ctx, authRequest)

	if err != nil {
		return dto.ResponseFailedStruct(ctx, err.Error(), fiber.ErrBadRequest.Code)
	}

	return dto.ResponseSuccessStruct(ctx, "Login successfuly", 200, userLogin)
}

func (a *AuthHandler) RefreshToken(ctx *fiber.Ctx) error {
	var refreshToken string

	if err := ctx.BodyParser(&refreshToken); err != nil {
		return dto.ResponseFailedStruct(ctx, err.Error(), fiber.ErrBadRequest.Code)
	}

	if refreshToken == "" {
		return dto.ResponseFailedStruct(ctx, "Refresh token is required", fiber.ErrUnauthorized.Code)
	}

	newAccessToken, err := a.service.RefreshTokenRequest(ctx, refreshToken)

	if err != nil {
		return dto.ResponseFailedStruct(ctx, err.Error(), fiber.ErrUnauthorized.Code)
	}

	return dto.ResponseSuccessStruct(ctx, "Token refreshed successfully", 200, newAccessToken)
}
