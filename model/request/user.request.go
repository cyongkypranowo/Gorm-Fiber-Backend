package request

import "go-fiber-gorm/utils"

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required,min=6"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type UserUpdateRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type UserEmailRequest struct {
	Email string `json:"email" validate:"email"`
}

func ValidateUserCreateRequest(r *UserCreateRequest) error {
	return utils.Validate.Struct(r)
}
