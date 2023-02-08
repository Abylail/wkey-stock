package errors

import "wkey-stock/src/data/models"

var (
	CategoryGetList = &models.Error{
		TechMessage:     "Get categories list error",
		BusinessMessage: defaultMessage,
	}
)
