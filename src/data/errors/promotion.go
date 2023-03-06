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
		TechMessage:     "Create new promotion model bind error",
		BusinessMessage: defaultMessage,
	}
	PromotionCreateValidate = &models.Error{
		TechMessage:     "Validate promotion model error",
		BusinessMessage: defaultMessage,
	}
	PromotionCreate = &models.Error{
		TechMessage:     "Create new promotion error",
		BusinessMessage: defaultMessage,
	}
	PromotionUpdate = &models.Error{
		TechMessage:     "Update promotion error",
		BusinessMessage: defaultMessage,
	}
	PromotionUpload = &models.Error{
		TechMessage:     "Upload promotion error",
		BusinessMessage: defaultMessage,
	}
	PromotionImageUpload = &models.Error{
		TechMessage:     "Upload image error",
		BusinessMessage: defaultMessage,
	}
	PromotionImageUpdate = &models.Error{
		TechMessage:     "Update image error",
		BusinessMessage: defaultMessage,
	}
)
