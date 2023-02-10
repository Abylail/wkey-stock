package errors

import "wkey-stock/src/data/models"

var (
	AdminProductGet = &models.Error{
		TechMessage:     "Get products list error",
		BusinessMessage: defaultMessage,
	}
)
