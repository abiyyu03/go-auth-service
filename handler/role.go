package handler

import (
	"strconv"

	"github.com/abiyyu03/auth-service/entity/dto"
	"github.com/abiyyu03/auth-service/entity/model"
	"github.com/abiyyu03/auth-service/service"
	"github.com/gofiber/fiber/v2"
)

type RoleHandlerInterface interface {
	Find(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}

type RoleHandler struct {
	service *service.RoleService
}

func NewRoleHandler(service *service.RoleService) RoleHandlerInterface {
	return &RoleHandler{
		service: service,
	}
}

func (r *RoleHandler) Find(ctx *fiber.Ctx) error {
	roles, err := r.service.Find()

	if err != nil {
		return dto.ResponseFailedStruct(ctx, err.Error(), fiber.StatusInternalServerError)
	}

	// if len(roles) == 0 {
	// 	return dto.ResponseSuccessStruct(ctx, "No roles found", fiber.StatusNotFound, nil)
	// }

	return dto.ResponseSuccessStruct(ctx, "Roles fetched successfully", fiber.StatusOK, roles)
}

func (r *RoleHandler) FindById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	castedId, _ := strconv.Atoi(id)

	role, err := r.service.FindById(castedId)

	if role == nil {
		return dto.ResponseFailedStruct(ctx, err.Error(), fiber.StatusNotFound)
	}

	if err != nil {
		return dto.ResponseFailedStruct(ctx, err.Error(), fiber.StatusInternalServerError)
	}

	return dto.ResponseSuccessStruct(ctx, "Role fetched successfully", fiber.StatusOK, role)
}

func (r *RoleHandler) Create(ctx *fiber.Ctx) error {
	var roleCreate *model.Role

	if err := ctx.BodyParser(&roleCreate); err != nil {
		return dto.ResponseFailedStruct(ctx, "Invalid request body", fiber.StatusBadRequest)
	}

	err := r.service.Create(roleCreate)

	if err != nil {
		return dto.ResponseFailedStruct(ctx, err.Error(), fiber.StatusInternalServerError)
	}

	return dto.ResponseSuccessStruct(ctx, "Role created successfully", fiber.StatusCreated, nil)
}

func (r *RoleHandler) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	castedId, _ := strconv.Atoi(id)

	var roleUpdate *model.Role

	if err := ctx.BodyParser(&roleUpdate); err != nil {
		return dto.ResponseFailedStruct(ctx, "Invalid request body", fiber.StatusBadRequest)
	}

	err := r.service.Update(roleUpdate, castedId)

	if err != nil {
		return dto.ResponseFailedStruct(ctx, err.Error(), fiber.StatusInternalServerError)
	}

	return dto.ResponseSuccessStruct(ctx, "Role updated successfully", fiber.StatusOK, nil)
}
