package brands_gateway

import "wkey-stock/src/repositories/brand_repository"

type Gateway struct {
	brandRepo *brand_repository.Repository
}

func New(brandRepo *brand_repository.Repository) *Gateway {
	return &Gateway{
		brandRepo: brandRepo,
	}
}
