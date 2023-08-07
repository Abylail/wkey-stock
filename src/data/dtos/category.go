package dtos

import (
	"github.com/google/uuid"
	"time"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

type Category struct {
	id         uuid.UUID
	proskladID int
	titleRU    string
	titleKZ    string
	position   int
	count      int
	createdAt  time.Time
	updatedAt  time.Time
}

func NewCategory(entity *entities.Category) *Category {
	id, _ := uuid.Parse(entity.ID)
	return &Category{
		id:         id,
		proskladID: entity.ProskladID,
		titleRU:    entity.TitleRU,
		position:   entity.Position,
		count:      entity.Count,
		createdAt:  entity.CreatedAt,
		updatedAt:  entity.UpdatedAt,
	}
}

func (category *Category) ID() uuid.UUID {
	return category.id
}

func (category *Category) EditProsklad(model *models.CategoryProsklad) {
	defer category.updateDate()

	category.titleRU = model.Title
	category.count = model.Count
	category.position = model.Position
}

func (category *Category) EditCount(count int) {
	category.count = count
}

func (category *Category) updateDate() {
	category.updatedAt = time.Now()
}

func (category *Category) Model() *models.Category {
	return &models.Category{
		ID:         category.ID().String(),
		ProskladID: category.proskladID,
		TitleRU:    category.titleRU,
		TitleKZ:    category.titleKZ,
		Position:   category.position,
		Count:      category.count,
	}
}

func (category *Category) Entity() entities.Category {
	return entities.Category{
		ID:         category.ID().String(),
		ProskladID: category.proskladID,
		TitleRU:    category.titleRU,
		TitleKZ:    category.titleKZ,
		Position:   category.position,
		Count:      category.count,
		CreatedAt:  category.createdAt,
		UpdatedAt:  category.updatedAt,
	}
}
