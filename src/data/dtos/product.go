package dtos

import (
	"github.com/google/uuid"
	"github.com/lowl11/boost/pkg/types"
	"time"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
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
	havInventory bool
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
	DescriptionRU string
	DescriptionKZ string
	Count         int
	CreatedAt     time.Time
	UpdatedAt     time.Time
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
		havInventory: entity.HasInventory,
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
		DescriptionRU: types.ToString(entity.DescriptionRU),
		DescriptionKZ: types.ToString(entity.DescriptionKZ),
		Count:         entity.Count,
		CreatedAt:     entity.CreatedAt,
		UpdatedAt:     entity.UpdatedAt,
	}
}

func (product Product) ID() uuid.UUID {
	return product.id
}

func (product Product) Model() models.Product {
	return models.Product{
		ID:    product.id.String(),
		Title: product.title,
	}
}
