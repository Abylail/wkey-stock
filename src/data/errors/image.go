package errors

import "wkey-stock/src/data/models"

var (
	ImageUploadCategoryIcon = &models.Error{
		TechMessage:     "Upload category icon error",
		BusinessMessage: defaultMessage,
	}
	ImageUploadBrandIcon = &models.Error{
		TechMessage:     "Upload brand icon error",
		BusinessMessage: defaultMessage,
	}
)
