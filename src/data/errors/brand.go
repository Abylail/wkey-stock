package errors

import "wkey-stock/src/data/models"

var (
	BrandNotFound = &models.Error{
		TechMessage:     "Brand not found",
		BusinessMessage: defaultMessage,
	}
	BrandAlreadyExist = &models.Error{
		TechMessage:     "Brand already exist",
		BusinessMessage: defaultMessage,
	}

	BrandGetParam = &models.Error{
		TechMessage:     "Brand ID required",
		BusinessMessage: defaultMessage,
	}
	BrandUpdateParam = &models.Error{
		TechMessage:     "Brand ID required",
		BusinessMessage: defaultMessage,
	}
	BrandUploadParam = &models.Error{
		TechMessage:     "Brand ID required",
		BusinessMessage: defaultMessage,
	}
	BrandDeleteParam = &models.Error{
		TechMessage:     "Brand ID required",
		BusinessMessage: defaultMessage,
	}

	BrandAddBind = &models.Error{
		TechMessage:     "Add new brand model bind error",
		BusinessMessage: defaultMessage,
	}
	BrandUpdateBind = &models.Error{
		TechMessage:     "Update brand model bind error",
		BusinessMessage: defaultMessage,
	}
	BrandUploadBind = &models.Error{
		TechMessage:     "Upload brand model bind error",
		BusinessMessage: defaultMessage,
	}

	BrandAddValidate = &models.Error{
		TechMessage:     "Add new brand model validate error",
		BusinessMessage: defaultMessage,
	}
	BrandUpdateValidate = &models.Error{
		TechMessage:     "Update brand model validate error",
		BusinessMessage: defaultMessage,
	}
	BrandUploadValidate = &models.Error{
		TechMessage:     "Upload brand model validate error",
		BusinessMessage: defaultMessage,
	}

	BrandAdd = &models.Error{
		TechMessage:     "Add new brand error",
		BusinessMessage: defaultMessage,
	}
	BrandUpdate = &models.Error{
		TechMessage:     "Update brand error",
		BusinessMessage: defaultMessage,
	}
	BrandUpdateIcon = &models.Error{
		TechMessage:     "Update brand icon error",
		BusinessMessage: defaultMessage,
	}
	BrandDelete = &models.Error{
		TechMessage:     "Delete brand error",
		BusinessMessage: defaultMessage,
	}
	BrandGetList = &models.Error{
		TechMessage:     "Get list of brands error",
		BusinessMessage: defaultMessage,
	}
	BrandGetByTitle = &models.Error{
		TechMessage:     "Get brand by title error",
		BusinessMessage: defaultMessage,
	}
	BrandGetByID = &models.Error{
		TechMessage:     "Get brand by ID error",
		BusinessMessage: defaultMessage,
	}
)
