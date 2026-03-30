package schema

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required,oneof=admin user"`
	IsActive bool   `json:"is_active"`
}

type UpdateUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role" binding:"omitempty,oneof=admin user"`
	IsActive *bool  `json:"is_active"` 
}