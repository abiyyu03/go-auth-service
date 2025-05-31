package model

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	*gorm.Model
	ID        int    `json:"id"`
	RoleName  string `json:"role_name"`
	RoleCode  string `json:"role_code"`
	Users     []User `json:"-" gorm:"foreignKey:RoleID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
