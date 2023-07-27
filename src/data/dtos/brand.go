package dtos

import (
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

type Brand struct {
	id         int
	title      string
	image      *string
	proskladID int
}

func NewBrand(entity *entities.Brand) *Brand {
	return &Brand{
		id:         entity.ID,
		title:      entity.Title,
		image:      entity.Image,
		proskladID: entity.ProskladID,
	}
}

func (brand Brand) ID() int {
	return brand.id
}

func (brand Brand) Title() string {
	return brand.title
}

func (brand Brand) Image() *string {
	return brand.image
}

func (brand Brand) ProskladID() int {
	return brand.proskladID
}

func (brand Brand) Edit(title string, image *string) {
	if title != "" {
		brand.title = title
	}

	if image != nil {
		brand.image = image
	}
}

func (brand Brand) Model() models.Brand {
	return models.Brand{
		ID:    brand.ID(),
		Title: brand.Title(),
		Image: brand.Image(),
	}
}
