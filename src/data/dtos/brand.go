package dtos

import (
	"github.com/google/uuid"
	"time"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

type Brand struct {
	id    uuid.UUID
	title string
	code  string
	image string

	createdAt time.Time
	updatedAt time.Time
}

func NewBrand(entity *entities.Brand) *Brand {
	id, _ := uuid.Parse(entity.ID)

	return &Brand{
		id:    id,
		title: entity.Title,
		code:  entity.Code,
		image: entity.Image,

		createdAt: entity.CreatedAt,
		updatedAt: entity.UpdatedAt,
	}
}

func (brand *Brand) ID() uuid.UUID {
	return brand.id
}

func (brand *Brand) Model() *models.Brand {
	return &models.Brand{
		ID:    brand.id.String(),
		Title: brand.title,
		Code:  brand.code,
		Image: brand.image,

		CreatedAt: brand.createdAt.Format("2006-01-02 15:05:04"),
		UpdatedAt: brand.updatedAt.Format("2006-01-02 15:05:04"),
	}
}

func (brand *Brand) Entity() entities.Brand {
	return entities.Brand{
		ID:    brand.id.String(),
		Title: brand.title,
		Code:  brand.code,
		Image: brand.image,

		CreatedAt: brand.createdAt,
		UpdatedAt: brand.updatedAt,
	}
}
