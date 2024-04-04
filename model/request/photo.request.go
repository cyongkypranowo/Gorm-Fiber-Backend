package request

import "go-fiber-gorm/utils"

type PhotoCreateRequest struct {
	CategoryId int64 `json:"category_id" form:"category_id" validate:"required"`
}

func ValidatePhotoCreateRequest(r *PhotoCreateRequest) error {
	return utils.Validate.Struct(r)
}
