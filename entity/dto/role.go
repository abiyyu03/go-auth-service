package dto

type RoleResponse struct {
	ID       int    `json:"id"`
	RoleName string `json:"role_name"`
	RoleCode string `json:"role_code"`
}
