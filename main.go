package main

import (
	"github.com/abiyyu03/auth-service/driver"
	"github.com/abiyyu03/auth-service/handler"
	"github.com/abiyyu03/auth-service/handler/middleware"
	"github.com/abiyyu03/auth-service/repository"
	"github.com/abiyyu03/auth-service/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func setupDatabase() (db *gorm.DB, err error) {
	db, err = driver.InitDB(driver.PostgresOption{
		Host:                   "localhost",
		Username:               "user_pemula",
		Password:               "abiyyucakra99",
		Name:                   "db_auth",
		Port:                   "5432",
		Timezone:               "Asia/Jakarta",
		DBName:                 "db_auth",
		MaxPoolSize:            100,
		BatchSize:              100,
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	return
}

func main() {
	app := fiber.New()
	api := app.Group("/api")

	postgres, err := setupDatabase()
	driver.MigrateDB(postgres)

	if err != nil {
		panic("DB Failed to connect: " + err.Error())
	}

	users := &repository.UserRepo{
		DB: postgres,
	}
	roles := &repository.RoleRepo{
		DB: postgres,
	}

	userServices := &service.UserService{
		Repo: users,
	}
	authServices := &service.AuthService{
		Repo: users,
	}
	roleServices := &service.RoleService{
		Repo: roles,
	}

	handlers := &handler.Handlers{
		User: handler.NewUserHandler(userServices),
		Auth: handler.NewAuthHandler(authServices),
		Role: handler.NewRoleHandler(roleServices),
	}

	api.Get("/users", middleware.HandleJWTMiddleware(roleServices, []int{1}), handlers.User.Find)
	api.Get("/users/:id", middleware.HandleJWTMiddleware(roleServices, []int{1}), handlers.User.FindById)
	api.Put("/users/:id", middleware.HandleJWTMiddleware(roleServices, []int{1}), handlers.User.Find)
	api.Post("/users", middleware.HandleJWTMiddleware(roleServices, []int{1}), handlers.User.Register)
	api.Post("/users", middleware.HandleJWTMiddleware(roleServices, []int{1}), handlers.User.Register)
	api.Get("/roles", middleware.HandleJWTMiddleware(roleServices, []int{1}), handlers.Role.Find)
	api.Get("/roles/:id", middleware.HandleJWTMiddleware(roleServices, []int{1}), handlers.Role.FindById)
	api.Post("/roles", middleware.HandleJWTMiddleware(roleServices, []int{1}), handlers.Role.Create)
	api.Put("/roles/:id", middleware.HandleJWTMiddleware(roleServices, []int{1}), handlers.Role.Update)
	api.Post("/login", handlers.Auth.Login)

	app.Listen(":3000")
}
