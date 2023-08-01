package dtos

import (
	"github.com/google/uuid"
	"time"
	"wkey-stock/src/data/models"
)

func NewCategoryAdd(model *models.CategoryProsklad) *Category {
	return &Category{
		id:         uuid.New(),
		proskladID: model.ID,
		title:      model.Title,
		position:   model.Position,
		count:      model.Count,
		createdAt:  time.Now(),
		updatedAt:  time.Now(),
	}
}
