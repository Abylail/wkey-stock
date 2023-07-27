package product_category_adaptor

import (
	"wkey-stock/src/data/dtos"
	"wkey-stock/src/data/entities"
)

func EntityToDTO(entityList []entities.ProductCategoryPair) []dtos.ProductCategory {
	dtoList := make([]dtos.ProductCategory, 0, len(entityList))

	for _, productCategory := range entityList {
		dtoList = append(dtoList, *dtos.NewProductCategory(&productCategory))
	}

	return dtoList
}
