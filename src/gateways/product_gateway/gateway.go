package product_gateway

import "wkey-stock/src/repositories/product_repository"

type Gateway struct {
	productRepo *product_repository.Repository
}

func New(productRepo *product_repository.Repository) *Gateway {
	return &Gateway{
		productRepo: productRepo,
	}
}
