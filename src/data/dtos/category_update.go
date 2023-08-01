package dtos

import "wkey-stock/src/data/models"

type CategoryUpdateCount struct {
	count int
}

func NewCategoryUpdateCount(model *models.CategoryUpdateCount) *CategoryUpdateCount {
	return &CategoryUpdateCount{
		count: model.Count,
	}
}

func (update CategoryUpdateCount) Count() int {
	return update.count
}
