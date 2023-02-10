package errors

import "wkey-stock/src/data/models"

var (
	AdminProductGet = &models.Error{
		TechMessage:     "Get products list error",
		BusinessMessage: defaultMessage,
	}
	AdminProductCountGet = &models.Error{
		TechMessage:     "Get products count error",
		BusinessMessage: defaultMessage,
	}
	ProductImagesGet = &models.Error{
		TechMessage:     "Get product images list error",
		BusinessMessage: defaultMessage,
	}
)
