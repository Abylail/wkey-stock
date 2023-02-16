package errors

import "wkey-stock/src/data/models"

var (
	SubCategoryActivateEmpty = &models.Error{
		TechMessage:     "Sub category has no binded products",
		BusinessMessage: defaultMessage,
	}

	SubCategoryBindProductListBind = &models.Error{
		TechMessage:     "Bind product list to sub category model bind error",
		BusinessMessage: defaultMessage,
	}

	SubCategoryBindProductListValidate = &models.Error{
		TechMessage:     "Bind product list to sub category model validate error",
		BusinessMessage: defaultMessage,
	}

	SubCategoryBindProductItem = &models.Error{
		TechMessage:     "Bind product item to sub category error",
		BusinessMessage: defaultMessage,
	}
	SubCategoryBindProductList = &models.Error{
		TechMessage:     "Bind product list to sub category error",
		BusinessMessage: defaultMessage,
	}
)
