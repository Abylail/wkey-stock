package gateways

import (
	"wkey-stock/src/gateways/brands_gateway"
	"wkey-stock/src/gateways/category_gateway"
	"wkey-stock/src/gateways/product_gateway"
	"wkey-stock/src/gateways/promotions_gateway"
	"wkey-stock/src/gateways/subcategory_gateway"
	"wkey-stock/src/repositories"
)

type Gateways struct {
	Products      *product_gateway.Gateway
	Categories    *category_gateway.Gateway
	Brands        *brands_gateway.Gateway
	Promotions    *promotions_gateway.Gateway
	SubCategories *subcategory_gateway.Gateway
}

func Get(repositories *repositories.ApiRepositories) *Gateways {
	return &Gateways{
		Products:      product_gateway.New(repositories.Product),
		Categories:    category_gateway.New(repositories.Category),
		Brands:        brands_gateway.New(repositories.Brand),
		Promotions:    promotions_gateway.New(repositories.Promotion),
		SubCategories: subcategory_gateway.New(repositories.SubCategory),
	}
}
