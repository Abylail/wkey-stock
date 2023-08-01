package category_gateway

import (
	"wkey-stock/src/repositories/category_repository"
)

type Gateway struct {
	categoryRepo *category_repository.Repository
}

func New(categoryRepo *category_repository.Repository) *Gateway {
	return &Gateway{
		categoryRepo: categoryRepo,
	}
}
