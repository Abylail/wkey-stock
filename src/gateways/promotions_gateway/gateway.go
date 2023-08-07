package promotions_gateway

import "wkey-stock/src/repositories/promotion_repository"

type Gateway struct {
	promRepo *promotion_repository.Repository
}

func New(promRepo *promotion_repository.Repository) *Gateway {
	return &Gateway{
		promRepo: promRepo,
	}
}
