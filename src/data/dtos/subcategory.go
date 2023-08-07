package dtos

import (
	"github.com/google/uuid"
	"time"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

type SubCategory struct {
	id      uuid.UUID
	titleRU string
	titleKZ string
	code    string

	createdAt time.Time
	updatedAt time.Time
}

func NewSubCategory(entity *entities.SubCategory) *SubCategory {
	id, _ := uuid.Parse(entity.ID)

	return &SubCategory{
		id:      id,
		titleRU: entity.TitleRU,
		titleKZ: entity.TitleKZ,
		code:    entity.Code,

		createdAt: entity.CreatedAt,
		updatedAt: entity.UpdatedAt,
	}
}

func (subCategory *SubCategory) ID() uuid.UUID {
	return subCategory.id
}

func (subCategory *SubCategory) Model() *models.SubCategory {
	return &models.SubCategory{
		ID:      subCategory.id.String(),
		TitleRU: subCategory.titleRU,
		TitleKZ: subCategory.titleKZ,
		Code:    subCategory.code,

		CreatedAt: subCategory.createdAt.Format("2006-01-02 15:05:04"),
		UpdatedAt: subCategory.updatedAt.Format("2006-01-02 15:05:04"),
	}
}

func (subCategory *SubCategory) Entity() entities.SubCategory {
	return entities.SubCategory{
		ID:      subCategory.id.String(),
		TitleRU: subCategory.titleRU,
		TitleKZ: subCategory.titleKZ,
		Code:    subCategory.code,

		CreatedAt: subCategory.createdAt,
		UpdatedAt: subCategory.updatedAt,
	}
}
