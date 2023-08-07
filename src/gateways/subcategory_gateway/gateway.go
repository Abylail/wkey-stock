package subcategory_gateway

import "wkey-stock/src/repositories/subcategory_repository"

type Gateway struct {
	subCatRepo *subcategory_repository.Repository
}

func New(subCatRepo *subcategory_repository.Repository) *Gateway {
	return &Gateway{
		subCatRepo: subCatRepo,
	}
}
