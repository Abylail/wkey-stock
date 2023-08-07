package dtos

import (
	"github.com/google/uuid"
	"time"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

type Promotion struct {
	id         uuid.UUID
	link       string
	linkTextRU string
	linkTextKZ string
	image      string
	position   int

	createdAt time.Time
	updatedAt time.Time
}

func NewPromotion(entity *entities.Promotion) *Promotion {
	id, _ := uuid.Parse(entity.ID)

	return &Promotion{
		id:         id,
		link:       entity.Link,
		linkTextRU: entity.LinkTextRU,
		linkTextKZ: entity.LinkTextKZ,
		image:      entity.Image,
		position:   entity.Position,

		createdAt: entity.CreatedAt,
		updatedAt: entity.UpdatedAt,
	}
}

func (promotion *Promotion) ID() uuid.UUID {
	return promotion.id
}

func (promotion *Promotion) Entity() entities.Promotion {
	return entities.Promotion{
		ID:         promotion.id.String(),
		Link:       promotion.link,
		LinkTextRU: promotion.linkTextRU,
		LinkTextKZ: promotion.linkTextKZ,
		Image:      promotion.image,
		Position:   promotion.position,

		CreatedAt: promotion.createdAt,
		UpdatedAt: promotion.updatedAt,
	}
}

func (promotion *Promotion) Model() *models.Promotion {
	return &models.Promotion{
		ID:         promotion.id.String(),
		Link:       promotion.link,
		LinkTextRU: promotion.linkTextRU,
		LinkTextKZ: promotion.linkTextKZ,
		Image:      promotion.image,
		Position:   promotion.position,

		CreatedAt: promotion.createdAt.Format("2006-01-02 15:05:04"),
		UpdatedAt: promotion.updatedAt.Format("2006-01-02 15:05:04"),
	}
}
