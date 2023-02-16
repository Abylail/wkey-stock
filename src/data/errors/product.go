package errors

import "wkey-stock/src/data/models"

var (
	AdminProductNotFound = &models.Error{
		TechMessage:     "Product not found",
		BusinessMessage: defaultMessage,
	}

	ProductUpdateParam = &models.Error{
		TechMessage:     "Product ID required",
		BusinessMessage: defaultMessage,
	}

	ProductUpdateBind = &models.Error{
		TechMessage:     "Update product model bind error",
		BusinessMessage: defaultMessage,
	}
	ProductUploadBind = &models.Error{
		TechMessage:     "Update product model bind error",
		BusinessMessage: defaultMessage,
	}

	ProductUpdateValidate = &models.Error{
		TechMessage:     "Update product model validate error",
		BusinessMessage: defaultMessage,
	}
	ProductUploadValidate = &models.Error{
		TechMessage:     "Update product model validate error",
		BusinessMessage: defaultMessage,
	}

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
	ProductUpdate = &models.Error{
		TechMessage:     "Update product error",
		BusinessMessage: defaultMessage,
	}
	ProductUpload = &models.Error{
		TechMessage:     "Upload product images error",
		BusinessMessage: defaultMessage,
	}
)
