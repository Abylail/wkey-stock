package errors

import "wkey-stock/src/data/models"

var (
	PromotionGetList = &models.Error{
		TechMessage:     "Promotion get list error",
		BusinessMessage: defaultMessage,
	}
	PromotionGetById = &models.Error{
		TechMessage:     "Promotion get by id error",
		BusinessMessage: defaultMessage,
	}
	PromotionGetByCode = &models.Error{
		TechMessage:     "Promotion get by code error",
		BusinessMessage: defaultMessage,
	}
	PromotionNotFound = &models.Error{
		TechMessage:     "Promotion not found",
		BusinessMessage: defaultMessage,
	}
	PromotionCreateBind = &models.Error{
		TechMessage:     "Add new promotion model bind error",
		BusinessMessage: defaultMessage,
	}
	PromotionCreate = &models.Error{
		TechMessage:     "Create new promotion error",
		BusinessMessage: defaultMessage,
	}
)
