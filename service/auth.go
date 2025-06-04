package service

import (
	"github.com/abiyyu03/auth-service/entity/dto"
	"github.com/abiyyu03/auth-service/entity/model"
	"github.com/abiyyu03/auth-service/repository"
	"github.com/abiyyu03/auth-service/service/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthServiceInterface interface {
	Login(ctx *fiber.Ctx, auth *dto.AuthRequest) (resp *dto.AuthResponse, err error)
}

type AuthService struct {
	Repo          *repository.UserRepo
	AuthTokenRepo *repository.AuthTokenRepo
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

	//create access token
	accessToken, err := utils.CreateJWT("access", user.Email, user.Fullname, user.Role.ID)

	if err != nil {
		err = fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
		return nil, err
	}
	decryptAccess, err := utils.VerifyJWT(accessToken)

	if err != nil {
		err = fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
		return nil, err
	}

	// create refresh token
	refreshToken, err := utils.CreateJWT("refresh", user.Email, user.Fullname, user.Role.ID)
	decryptRefresh, err := utils.VerifyJWT(refreshToken)

	if err != nil {
		err = fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
		return nil, err
	}

	// store auth token in database
	tokenData := &model.AuthToken{
		UserID:         user.ID,
		AccessToken:    accessToken,
		AccessExpires:  decryptAccess.ExpiresAt.String(),
		RefreshToken:   refreshToken,
		RefreshExpires: decryptRefresh.ExpiresAt.String(),
		Revoked:        false,
		Expired:        false,
		IPAddress:      ctx.IP(),
		Device:         ctx.Get("User-Agent"),
	}

	err = service.AuthTokenRepo.CreateAuthToken(tokenData)

	if err != nil {
		err = fiber.NewError(fiber.StatusInternalServerError, "Failed to create auth token")
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
			AccessToken:    accessToken,
			AccessExpires:  decryptAccess.ExpiresAt.Unix(),
			RefreshToken:   refreshToken,
			RefreshExpires: decryptRefresh.ExpiresAt.Unix(),
		},
	}

	return
}
