package model

import (
	"time"

	"gorm.io/gorm"
)

type AuthToken struct {
	*gorm.Model
	ID           int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Token        string `json:"token" gorm:"type:text;not null;unique"`
	RefreshToken string `json:"refresh_token" gorm:"type:text;not null;unique"`
	IPAddress    string `json:"ip_address" gorm:"type:varchar(45);not null"`
	Device       string `json:"device" gorm:"type:varchar(100);not null"`
	Revoked      bool   `json:"revoked" gorm:"type:boolean;default:false;not null"`
	ExpiresAt    string `json:"expires_at" gorm:"type:timestamp;not null"`
	UserID       string `json:"user_id" gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `json:"deleted_at,omitempty" gorm:"index"`
}
