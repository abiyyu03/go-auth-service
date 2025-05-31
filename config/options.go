package config

import "gorm.io/gorm"

type Option struct {
	DB *gorm.DB
}
