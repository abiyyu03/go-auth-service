package repository

import "gorm.io/gorm"

type Option struct {
	DB *gorm.DB
}

type Repository struct {
	User UserRepoInterface
	Role RoleRepoInterface
}
