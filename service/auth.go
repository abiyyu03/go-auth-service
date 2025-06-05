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

	checkPassword, _ := utils.CheckHashPassword(user.Password, auth.Password)
	if checkPassword == false {
		err = fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
		return nil, err
	}

	//create access token
	accessToken, err := utils.CreateJWT("access", user.Email, user.Fullname, user.Role.ID)

	if err != nil {
		err = fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
		return nil, err
	}
	decryptAccess, err := utils.ParseJWT(accessToken)

	if err != nil {
		err = fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
		return nil, err
	}

	// create refresh token
	refreshToken, err := utils.CreateJWT("refresh", user.Email, user.Fullname, user.Role.ID)
	decryptRefresh, err := utils.ParseJWT(refreshToken)

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

func (service *AuthService) RefreshTokenRequest(ctx *fiber.Ctx, refreshToken string) (resp *dto.TokenResponse, err error) {
	claims, err := utils.ParseJWT(refreshToken)
	if err != nil {
		err = fiber.NewError(fiber.StatusUnauthorized, "Invalid refresh token")
		return nil, err
	}

	// Create a new access token
	newAccessToken, err := utils.CreateJWT("access", claims.Email, claims.Fullname, int(claims.RoleID))
	if err != nil {
		err = fiber.NewError(fiber.StatusInternalServerError, "Failed to create new access token")
		return nil, err
	}

	decryptAccess, err := utils.ParseJWT(newAccessToken)
	if err != nil {
		err = fiber.NewError(fiber.StatusInternalServerError, "Failed to parse new access token")
		return nil, err
	}

	newRefreshToken, err := utils.CreateJWT("refresh", claims.Email, claims.Fullname, int(claims.RoleID))
	if err != nil {
		err = fiber.NewError(fiber.StatusInternalServerError, "Failed to create new refresh token")
		return nil, err
	}
	decryptRefresh, err := utils.ParseJWT(newRefreshToken)
	if err != nil {
		err = fiber.NewError(fiber.StatusInternalServerError, "Failed to parse new refresh token")
		return nil, err
	}

	// Get the old auth token from the database
	oldAuthToken, err := service.AuthTokenRepo.GetAuthTokenByUserID(refreshToken)
	if err != nil {
		if err.Error() == "record not found" {
			err = fiber.NewError(fiber.StatusUnauthorized, "Refresh token not found")
			return nil, err
		}
		err = fiber.NewError(fiber.StatusInternalServerError, "Failed to get old auth token")
		return nil, err
	}

	// remove old auth token
	err = service.AuthTokenRepo.DeleteAuthTokenByUserID(refreshToken)
	if err != nil {
		err = fiber.NewError(fiber.StatusInternalServerError, "Failed to delete old auth token")
		return nil, err
	}
	// store new auth token in database
	tokenData := &model.AuthToken{
		UserID:         oldAuthToken.UserID,
		AccessToken:    newAccessToken,
		AccessExpires:  decryptAccess.ExpiresAt.String(),
		RefreshToken:   newRefreshToken,
		RefreshExpires: decryptRefresh.ExpiresAt.String(),
		Revoked:        false,
		Expired:        false,
		IPAddress:      ctx.IP(),
		Device:         ctx.Get("User-Agent"),
	}
	err = service.AuthTokenRepo.CreateAuthToken(tokenData)
	if err != nil {
		err = fiber.NewError(fiber.StatusInternalServerError, "Failed to create new auth token")
		return nil, err
	}

	resp = &dto.TokenResponse{
		AccessToken:    newAccessToken,
		AccessExpires:  decryptAccess.ExpiresAt.Unix(),
		RefreshToken:   newRefreshToken,
		RefreshExpires: decryptRefresh.ExpiresAt.Unix(),
	}

	return
}
