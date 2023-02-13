package errors

import "wkey-stock/src/data/models"

var (
	CategoryNotFound = &models.Error{
		TechMessage:     "Category not found",
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

	CategoryAddBind = &models.Error{
		TechMessage:     "Add category model bind error",
		BusinessMessage: defaultMessage,
	}
	CategoryUpdateBind = &models.Error{
		TechMessage:     "Update category model bind error",
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

	CategoryGetList = &models.Error{
		TechMessage:     "Get categories list error",
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
