package service

import (
	"github.com/abiyyu03/auth-service/constant/message"
	"github.com/abiyyu03/auth-service/entity/dto"
	"github.com/abiyyu03/auth-service/entity/model"
	"github.com/abiyyu03/auth-service/repository"
)

type RoleServiceInterface interface {
	Find() (resp []*dto.RoleResponse, err error)
	FindById(d int) (resp *dto.RoleResponse, err error)
	Create(role *model.Role) (err error)
	Update(role *model.Role, d int) (err error)
	// Delete(d int) (err error)
}

type RoleService struct {
	Repo *repository.RoleRepo
}

func NewRoleService(repo *repository.RoleRepo) RoleServiceInterface {
	return &RoleService{
		Repo: repo,
	}
}

func (service *RoleService) Find() (resp []*dto.RoleResponse, err error) {
	roles, err := service.Repo.Fetch()

	if err != nil {
		err = message.ErrInternalServer
		return nil, err
	}

	for _, r := range roles {
		resp = append(resp, &dto.RoleResponse{
			ID:       r.ID,
			RoleName: r.RoleName,
			RoleCode: r.RoleCode,
		})
	}
	return
}

func (service *RoleService) FindById(id int) (resp *dto.RoleResponse, err error) {
	role, err := service.Repo.FetchById(id)

	if role == nil {
		err = message.ErrNotFound
		return
	}

	if err != nil {
		err = message.ErrInternalServer
		return
	}

	resp = &dto.RoleResponse{
		ID:       role.ID,
		RoleName: role.RoleName,
		RoleCode: role.RoleCode,
	}
	return
}

func (service *RoleService) Create(role *model.Role) (err error) {
	err = service.Repo.Create(role)

	if err != nil {
		err = message.ErrInternalServer
		return err
	}

	return
}

func (service *RoleService) Update(role *model.Role, id int) (err error) {
	err = service.Repo.Update(role, id)

	if err != nil {
		err = message.ErrInternalServer
	}

	return
}
