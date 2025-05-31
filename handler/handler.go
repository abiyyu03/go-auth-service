package handler

import (
	"github.com/abiyyu03/auth-service/service"
)

type Option struct {
	*service.Service
}

type Handlers struct {
	User UserHandlerInterface
	Auth AuthHandlerInterface
	Role RoleHandlerInterface
}
