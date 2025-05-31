package repository

import (
	"github.com/abiyyu03/auth-service/entity/dto"
	"github.com/abiyyu03/auth-service/entity/model"
	"gorm.io/gorm"
)

type UserRepoInterface interface {
	Fetch() (userResponse []*model.User, err error)
	FetchById(id string) (userResponse *model.User, err error)
	FetchLogin(email string) (userResponse *model.User, err error)
	Create(user *dto.UserCreate) (err error)
	Update(user *dto.UserCreate, id string) (err error)
	Delete(id string) (err error)
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepoInterface {
	return &UserRepo{
		DB: db,
	}
}

func (repo *UserRepo) Fetch() (userResponse []*model.User, err error) {
	err = repo.DB.Preload("Role").Find(&userResponse).Error
	if err != nil {
		return nil, err
	}
	return
}
func (repo *UserRepo) FetchById(id string) (userResponse *model.User, err error) {
	err = repo.DB.First(&userResponse, "id =?", id).Error
	if err != nil {
		return nil, err
	}
	return
}
func (repo *UserRepo) FetchLogin(email string) (userResponse *model.User, err error) {
	err = repo.DB.Preload("Role").Where("email =?", email).First(&userResponse).Error
	return
}
func (repo *UserRepo) Create(user *dto.UserCreate) (err error) {
	err = repo.DB.Create(&model.User{
		ID:       user.ID,
		Email:    user.Email,
		Fullname: user.Fullname,
		Password: user.Password,
		RoleID:   user.RoleID,
	}).Error
	return
}
func (repo *UserRepo) Update(user *dto.UserCreate, id string) (err error) {
	err = repo.DB.Where("id =?", id).Updates(&user).Error
	return
}
func (repo *UserRepo) Delete(id string) (err error) {
	var user *model.User
	err = repo.DB.Delete(&user, "id =?", id).Error
	return
}
