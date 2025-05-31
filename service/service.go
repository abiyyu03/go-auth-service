package service

type Service struct {
	Auth AuthServiceInterface
	User UserServiceInterface
	Role RoleServiceInterface
}
