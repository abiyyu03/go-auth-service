package driver

import (
	"fmt"

	"github.com/abiyyu03/auth-service/entity/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresOption struct {
	Host                   string
	Username               string
	Password               string
	Name                   string
	Port                   string
	Timezone               string
	DBName                 string
	MaxPoolSize            int
	BatchSize              int
	PrepareStmt            bool
	SkipDefaultTransaction bool
}

func InitDB(opt PostgresOption) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		opt.Host,
		opt.Username,
		opt.Password,
		opt.Name,
		opt.Port,
		opt.Timezone,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt:            opt.PrepareStmt,
		SkipDefaultTransaction: opt.SkipDefaultTransaction,
		CreateBatchSize:        opt.BatchSize,
	})

	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return nil, err
	}

	return
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(
		&model.Role{},
		&model.User{},
		&model.AuthToken{},
	)
}
