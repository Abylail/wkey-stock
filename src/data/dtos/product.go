package dtos

import (
	"github.com/google/uuid"
	"time"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
	"wkey-stock/src/enums/languages"
)

type Product struct {
	// common
	id               uuid.UUID
	proskladID       int
	title            string
	barcode          string
	companyID        int
	itemCategoryName string

	// price
	sellingPrice          float32
	oldSellingPrice       float32
	previousPurchasePrice float32
	additionalPercent     float32

	// flags
	hasInventory bool
	isVirtual    bool
	marked       bool
	isQuick      bool

	// unit
	unitID   int
	unitName string
	unitType int

	// brand (vendor)
	brandID int

	// custom
	descriptionRU *string
	descriptionKZ *string
	count         int
	createdAt     time.Time
	updatedAt     time.Time
}

func NewProduct(entity *entities.Product) *Product {
	id, _ := uuid.Parse(entity.ID)

	return &Product{
		// common
		id:               id,
		proskladID:       entity.ProskladID,
		title:            entity.Title,
		barcode:          entity.Barcode,
		companyID:        entity.CompanyID,
		itemCategoryName: entity.ItemCategoryName,

		// price
		sellingPrice:          entity.SellingPrice,
		oldSellingPrice:       entity.OldSellingPrice,
		previousPurchasePrice: entity.PreviousPurchasePrice,
		additionalPercent:     entity.AdditionalPercent,

		// flags
		hasInventory: entity.HasInventory,
		isVirtual:    entity.IsVirtual,
		marked:       entity.Marked,
		isQuick:      entity.IsQuick,

		// unit
		unitID:   entity.UnitID,
		unitName: entity.UnitName,
		unitType: entity.UnitType,

		// brand (vendor)
		brandID: entity.BrandID,

		// custom
		descriptionRU: entity.DescriptionRU,
		descriptionKZ: entity.DescriptionKZ,
		count:         entity.Count,
		createdAt:     entity.CreatedAt,
		updatedAt:     entity.UpdatedAt,
	}
}

func (product *Product) ID() uuid.UUID {
	return product.id
}

func (product *Product) Description(language string) *string {
	if language == languages.KZ {
		return product.descriptionKZ
	}

	return product.descriptionRU
}

func (product *Product) EditDescription(description, language string) {
	defer product.update()

	if language == languages.KZ {
		product.descriptionKZ = &description
		return
	}

	product.descriptionRU = &description
}

func (product *Product) EditCount(count int) {
	defer product.update()

	product.count = count
}

func (product *Product) Model() models.Product {
	return models.Product{
		ID:            product.id.String(),
		ProskladID:    product.proskladID,
		Title:         product.title,
		DescriptionRU: product.Description(languages.RU),
		DescriptionKZ: product.Description(languages.KZ),
		Count:         product.count,
	}
}

func (product *Product) Entity() entities.Product {
	return entities.Product{
		ID:               product.id.String(),
		ProskladID:       product.proskladID,
		Title:            product.title,
		Barcode:          product.barcode,
		ItemCategoryName: product.itemCategoryName,

		SellingPrice:          product.sellingPrice,
		OldSellingPrice:       product.oldSellingPrice,
		PreviousPurchasePrice: product.previousPurchasePrice,
		AdditionalPercent:     product.additionalPercent,

		HasInventory: product.hasInventory,
		IsVirtual:    product.isVirtual,
		Marked:       product.marked,
		IsQuick:      product.isQuick,

		UnitID:   product.unitID,
		UnitName: product.unitName,
		UnitType: product.unitType,

		BrandID: product.brandID,

		DescriptionRU: product.descriptionRU,
		DescriptionKZ: product.descriptionKZ,
		Count:         product.count,
		CreatedAt:     product.createdAt,
		UpdatedAt:     product.updatedAt,
	}
}

func (product *Product) update() {
	product.updatedAt = time.Now()
}
