package repository

import (
	"github.com/abiyyu03/auth-service/entity/model"
	"gorm.io/gorm"
)

type AuthTokenRepoInterface interface {
	CreateAuthToken(data *model.AuthToken) error
	GetAuthTokenByUserID(token string) (*model.AuthToken, error)
	DeleteAuthTokenByUserID(token string) error
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

func (repo *AuthTokenRepo) GetAuthTokenByUserID(token string) (authToken *model.AuthToken, err error) {
	err = repo.DB.Where("refresh_token = ?", token).First(&authToken).Error
	if err != nil {
		return nil, err
	}
	return
}

func (repo *AuthTokenRepo) DeleteAuthTokenByUserID(token string) error {
	var authToken model.AuthToken
	err := repo.DB.Where("refresh_token = ?", token).Delete(&authToken).Error
	if err != nil {
		return err
	}
	return nil
}
