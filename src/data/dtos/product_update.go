package dtos

import (
	"wkey-stock/src/data/models"
)

type ProductUpdateDescription struct {
	descriptionRU string
	descriptionKZ string
}

type ProductUpdateCount struct {
	count int
}

func NewProductUpdateDescription(model *models.ProductUpdateDescription) *ProductUpdateDescription {
	return &ProductUpdateDescription{
		descriptionRU: model.DescriptionRU,
		descriptionKZ: model.DescriptionKZ,
	}
}

func NewProductUpdateCount(model *models.ProductUpdateCount) *ProductUpdateCount {
	return &ProductUpdateCount{
		count: model.Count,
	}
}

func (update ProductUpdateDescription) DescriptionRU() string {
	return update.descriptionRU
}

func (update ProductUpdateDescription) DescriptionKZ() string {
	return update.descriptionKZ
}

func (update ProductUpdateCount) Count() int {
	return update.count
}
