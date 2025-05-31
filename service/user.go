package service

import (
	"log"

	"github.com/abiyyu03/auth-service/constant/message"
	"github.com/abiyyu03/auth-service/entity/dto"
	"github.com/abiyyu03/auth-service/repository"
	"github.com/abiyyu03/auth-service/service/utils"
	"github.com/google/uuid"
)

type UserServiceInterface interface {
	Find() (resp []*dto.UserResponse, err error)
	FindById(id string) (resp *dto.UserResponse, err error)
	Register(user *dto.UserCreate) (resp *dto.UserResponse, err error)
	// Update(user *dto.UserCreate, id string) (err error)
	// Delete(id string) (err error)
}

type UserService struct {
	Repo *repository.UserRepo
}

func NewUserService(repo *repository.UserRepo) UserServiceInterface {
	return &UserService{
		Repo: repo,
	}
}

func (service *UserService) Find() (resp []*dto.UserResponse, err error) {
	user, err := service.Repo.Fetch()

	if err != nil {
		err = message.ErrInternalServer
		return nil, err
	}
	log.Println("UserHandler Find", user)
	for _, u := range user {
		resp = append(resp, &dto.UserResponse{
			ID:       u.ID,
			Email:    u.Email,
			Fullname: u.Fullname,
			RoleID:   u.Role.ID,
			RoleName: u.Role.RoleName,
		})
	}
	return
}

func (service *UserService) FindById(id string) (resp *dto.UserResponse, err error) {
	user, err := service.Repo.FetchById(id)

	if user == nil {
		err = message.ErrNotFound
		return
	}

	if err != nil {
		err = message.ErrInternalServer
		return
	}
	resp = &dto.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Fullname: user.Fullname,
		RoleID:   user.Role.ID,
		RoleName: user.Role.RoleName,
	}
	return
}

func (service *UserService) Register(user *dto.UserCreate) (resp *dto.UserResponse, err error) {
	hashedPassword, err := utils.HashPassword(user.Password)

	id, _ := uuid.NewV7()

	err = service.Repo.Create(&dto.UserCreate{
		ID:       id.String(),
		Email:    user.Email,
		Fullname: user.Fullname,
		Password: hashedPassword,
		RoleID:   user.RoleID,
	})

	if err != nil {
		err = message.ErrInternalServer
	}

	return
}

// func (service *UserService) Update(user *dto.UserCreate, id string) (err error) {

// }
// func (service *UserService) Delete(id string) (err error) {

// }
