package errors

import "wkey-stock/src/data/models"

var (
	CategoryNotFound = &models.Error{
		TechMessage:     "Category not found",
		BusinessMessage: defaultMessage,
	}
	SubCategoryNotFound = &models.Error{
		TechMessage:     "Sub category not found",
		BusinessMessage: defaultMessage,
	}
	CategoryAlreadyExist = &models.Error{
		TechMessage:     "Category already exist",
		BusinessMessage: defaultMessage,
	}

	CategoryDeleteParam = &models.Error{
		TechMessage:     "Category code required",
		BusinessMessage: defaultMessage,
	}
	CategoryUpdateParam = &models.Error{
		TechMessage:     "Category code required",
		BusinessMessage: defaultMessage,
	}
	CategoryUploadParam = &models.Error{
		TechMessage:     "Category code required",
		BusinessMessage: defaultMessage,
	}

	CategoryAddBind = &models.Error{
		TechMessage:     "Add category model bind error",
		BusinessMessage: defaultMessage,
	}
	CategoryUpdateBind = &models.Error{
		TechMessage:     "Update category model bind error",
		BusinessMessage: defaultMessage,
	}
	CategoryUploadBind = &models.Error{
		TechMessage:     "Upload category icon model bind error",
		BusinessMessage: defaultMessage,
	}

	CategoryAddValidate = &models.Error{
		TechMessage:     "Add category model validate error",
		BusinessMessage: defaultMessage,
	}
	CategoryUpdateValidate = &models.Error{
		TechMessage:     "Update category model validate error",
		BusinessMessage: defaultMessage,
	}
	CategoryUploadValidate = &models.Error{
		TechMessage:     "Upload category icon model validate error",
		BusinessMessage: defaultMessage,
	}

	CategoryGetList = &models.Error{
		TechMessage:     "Get categories list error",
		BusinessMessage: defaultMessage,
	}
	CategoryGetByCode = &models.Error{
		TechMessage:     "Get category by code error",
		BusinessMessage: defaultMessage,
	}
	CategoryAdd = &models.Error{
		TechMessage:     "Add category error",
		BusinessMessage: defaultMessage,
	}
	CategoryUpdate = &models.Error{
		TechMessage:     "Update category error",
		BusinessMessage: defaultMessage,
	}
	CategoryDelete = &models.Error{
		TechMessage:     "Delete category error",
		BusinessMessage: defaultMessage,
	}
)
