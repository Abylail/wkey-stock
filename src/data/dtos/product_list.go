package dtos

import "wkey-stock/src/data/models"

type ProductList struct {
	pageCount  int
	products   []Product
	categories []ProductCategory
}

func NewProductList(pageCount int, products []Product, categories []ProductCategory) *ProductList {
	return &ProductList{
		pageCount:  pageCount,
		products:   products,
		categories: categories,
	}
}

func (list ProductList) Model() models.AdminProductGet {
	products := make([]models.AdminProductItem, 0, len(list.products))
	for _, product := range list.products {
		products = append(products, product.Model(list.categories))
	}

	return models.AdminProductGet{
		PageCount: list.pageCount,
		List:      products,
	}
}
