package dtos

import (
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

type ProductCategory struct {
	productID       int
	subCategoryID   int
	subCategoryName string
	subCategoryCode string
	categoryName    string
	categoryCode    string
}

func NewProductCategory(entity *entities.ProductCategoryPair) *ProductCategory {
	return &ProductCategory{
		productID:       entity.ProductID,
		subCategoryID:   entity.SubCategoryID,
		subCategoryName: entity.SubCategoryName,
		subCategoryCode: entity.SubCategoryCode,
		categoryName:    entity.CategoryName,
		categoryCode:    entity.CategoryCode,
	}
}

func (productCategory ProductCategory) ProductID() int {
	return productCategory.productID
}

func (productCategory ProductCategory) SubCategoryID() int {
	return productCategory.subCategoryID
}

func (productCategory ProductCategory) SubCategoryName() string {
	return productCategory.subCategoryName
}

func (productCategory ProductCategory) SubCategoryCode() string {
	return productCategory.subCategoryCode
}

func (productCategory ProductCategory) CategoryName() string {
	return productCategory.categoryName
}

func (productCategory ProductCategory) CategoryCode() string {
	return productCategory.categoryCode
}

func (productCategory ProductCategory) Model() models.ProductCategoryPair {
	return models.ProductCategoryPair{
		SubCategoryCode: productCategory.SubCategoryCode(),
		SubCategoryName: productCategory.SubCategoryName(),
		CategoryCode:    productCategory.CategoryCode(),
		CategoryName:    productCategory.CategoryName(),
	}
}
