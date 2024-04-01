package request

import (
	"go-fiber-gorm/util"
)

type UserCreateRequest struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"email"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
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
	return util.Validate.Struct(r)
}
