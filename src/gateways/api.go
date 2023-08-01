package gateways

import (
	"wkey-stock/src/gateways/category_gateway"
	"wkey-stock/src/gateways/product_gateway"
	"wkey-stock/src/repositories"
)

type Gateways struct {
	Products   *product_gateway.Gateway
	Categories *category_gateway.Gateway
}

func Get(repositories *repositories.ApiRepositories) *Gateways {
	return &Gateways{
		Products:   product_gateway.New(repositories.Product),
		Categories: category_gateway.New(repositories.Category),
	}
}
