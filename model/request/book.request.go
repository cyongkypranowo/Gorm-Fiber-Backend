package request

import "go-fiber-gorm/utils"

type BookCreateRequest struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
}

func ValidateBookCreateRequest(r *BookCreateRequest) error {
	return utils.Validate.Struct(r)
}
