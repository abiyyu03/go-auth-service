package dto

type UserResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	RoleID   int    `json:"role_id"`
	RoleName string `json:"role_name"`
}

type UserCreate struct {
	ID       string `json:"id"`
	Email    string `json:"email" validate:"required"`
	Fullname string `json:"fullname" validate:"required"`
	Password string `json:"password" validate:"required"`
	RoleID   uint   `json:"role_id" validate:"required"`
}
