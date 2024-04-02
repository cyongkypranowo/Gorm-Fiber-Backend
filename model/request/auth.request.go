package request

import "go-fiber-gorm/util"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func ValidateLoginRequest(r *LoginRequest) error {
	return util.Validate.Struct(r)
}
