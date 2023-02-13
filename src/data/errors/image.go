package errors

import "wkey-stock/src/data/models"

var (
	ImageUploadCategoryIcon = &models.Error{
		TechMessage:     "Upload category icon error",
		BusinessMessage: defaultMessage,
	}
)
