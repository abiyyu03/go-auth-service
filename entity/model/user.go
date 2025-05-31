package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	ID        string `json:"id"`
	Email     string `json:"email"`
	Fullname  string `json:"fullname"`
	Password  string `json:"password"`
	RoleID    uint   `json:"role_id" gorm:"column:role_id;not null"`
	Role      Role   `json:"role" gorm:"foreignKey:RoleID;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
