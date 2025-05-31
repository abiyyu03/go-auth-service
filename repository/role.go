package repository

import (
	"github.com/abiyyu03/auth-service/entity/model"
	"gorm.io/gorm"
)

type RoleRepoInterface interface {
	Fetch() (roleResponse []*model.Role, err error)
	FetchById(id int) (roleResponse *model.Role, err error)
	Create(role *model.Role) (err error)
	Update(role *model.Role, id int) (err error)
	Delete(id int) (err error)
}

type RoleRepo struct {
	DB *gorm.DB
}

func NewRoleRepo(db *gorm.DB) RoleRepoInterface {
	return &RoleRepo{
		DB: db,
	}
}

func (repo *RoleRepo) Fetch() (roleResponse []*model.Role, err error) {
	err = repo.DB.Find(&roleResponse).Error
	if err != nil {
		return nil, err
	}
	return
}

func (repo *RoleRepo) FetchById(id int) (roleResponse *model.Role, err error) {
	err = repo.DB.First(&roleResponse, "id =?", id).Error
	if err != nil {
		return nil, err
	}
	return
}

func (repo *RoleRepo) Create(role *model.Role) (err error) {
	err = repo.DB.Create(&role).Error
	return
}

func (repo *RoleRepo) Update(role *model.Role, id int) (err error) {
	err = repo.DB.Where("id =?", id).Updates(&role).Error
	return
}

func (repo *RoleRepo) Delete(id int) (err error) {
	var role *model.Role
	err = repo.DB.Delete(&role, "id =?", id).Error
	return
}
