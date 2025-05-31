package handler

import (
	"log"

	"github.com/abiyyu03/auth-service/constant/message"
	"github.com/abiyyu03/auth-service/entity/dto"
	"github.com/abiyyu03/auth-service/service"
	"github.com/gofiber/fiber/v2"
)

type UserHandlerInterface interface {
	Find(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
}

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) UserHandlerInterface {
	return &UserHandler{
		service: service,
	}
}

func (u *UserHandler) Find(ctx *fiber.Ctx) error {
	user, err := u.service.Find()

	if err != nil {
		log.Fatal("ERROR", err.Error())
		return dto.ResponseFailedStruct(ctx, err.Error(), fiber.StatusInternalServerError)
	}

	return dto.ResponseSuccessStruct(ctx, message.SuccessFetched, fiber.StatusOK, user)
}

func (u *UserHandler) FindById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	user, err := u.service.FindById(id)

	if user == nil {
		return dto.ResponseFailedStruct(ctx, err.Error(), fiber.StatusNotFound)
	}

	if err != nil {
		return dto.ResponseFailedStruct(ctx, err.Error(), fiber.StatusInternalServerError)
	}

	return dto.ResponseSuccessStruct(ctx, message.SuccessFetched, fiber.StatusOK, user)
}

func (u *UserHandler) Register(ctx *fiber.Ctx) error {
	var userCreate *dto.UserCreate

	if err := ctx.BodyParser(&userCreate); err != nil {
		return dto.ResponseFailedStruct(ctx, message.ErrBadRequest.Error(), fiber.StatusBadRequest)
	}

	user, err := u.service.Register(userCreate)

	if err != nil {
		return dto.ResponseFailedStruct(ctx, err.Error(), fiber.StatusInternalServerError)
	}

	return dto.ResponseSuccessStruct(ctx, message.SuccessCreated, fiber.StatusCreated, user)
}
