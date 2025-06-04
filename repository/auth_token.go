package repository

import (
	"github.com/abiyyu03/auth-service/entity/model"
	"gorm.io/gorm"
)

type AuthTokenRepoInterface interface {
	CreateAuthToken(data *model.AuthToken) error
	GetAuthTokenByUserID(userID string) (*model.AuthToken, error)
	DeleteAuthTokenByUserID(userID string) error
}

type AuthTokenRepo struct {
	DB *gorm.DB
}

func NewAuthTokenRepo(db *gorm.DB) AuthTokenRepoInterface {
	return &AuthTokenRepo{
		DB: db,
	}
}

func (repo *AuthTokenRepo) CreateAuthToken(data *model.AuthToken) error {
	err := repo.DB.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *AuthTokenRepo) GetAuthTokenByUserID(userID string) (*model.AuthToken, error) {
	var authToken model.AuthToken
	err := repo.DB.Where("user_id = ?", userID).First(&authToken).Error
	if err != nil {
		return nil, err
	}
	return &authToken, nil
}

func (repo *AuthTokenRepo) DeleteAuthTokenByUserID(userID string) error {
	var authToken model.AuthToken
	err := repo.DB.Where("user_id = ?", userID).Delete(&authToken).Error
	if err != nil {
		return err
	}
	return nil
}
