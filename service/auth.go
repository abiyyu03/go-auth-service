package service

import (
	"log"

	"github.com/abiyyu03/auth-service/entity/dto"
	"github.com/abiyyu03/auth-service/repository"
	"github.com/abiyyu03/auth-service/service/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthServiceInterface interface {
	Login(ctx *fiber.Ctx, auth *dto.AuthRequest) (resp *dto.AuthResponse, err error)
}

type AuthService struct {
	Repo *repository.UserRepo
}

func NewAuthService(repo *repository.UserRepo) AuthServiceInterface {
	return &AuthService{
		Repo: repo,
	}
}

func (service *AuthService) Login(ctx *fiber.Ctx, auth *dto.AuthRequest) (resp *dto.AuthResponse, err error) {
	user, err := service.Repo.FetchLogin(auth.Email)

	if user == nil {
		err = fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
		return nil, err
	}
	if err != nil {
		err = fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
		return nil, err
	}

	token, err := utils.CreateJWT(user.Email, user.Fullname, user.Role.ID)

	log.Println("Token failed to verified:", token)
	if err != nil {
		err = fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
		return nil, err
	}
	decryptToken, err := utils.VerifyJWT(token)

	if err != nil {
		err = fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
		return nil, err
	}

	resp = &dto.AuthResponse{
		UserData: dto.UserResponse{
			ID:       user.ID,
			Email:    user.Email,
			Fullname: user.Fullname,
			RoleID:   user.Role.ID,
			RoleName: user.Role.RoleName,
		},
		TokenData: dto.TokenResponse{
			AccessToken:    token,
			AccessExpires:  decryptToken.ExpiresAt.Unix(),
			RefreshToken:   token,
			RefreshExpires: decryptToken.ExpiresAt.Unix(),
		},
	}

	return
}
