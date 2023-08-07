package dtos

import (
	"github.com/google/uuid"
	"time"
	"wkey-stock/src/data/models"
)

func NewProductAdd(model *models.ProductProsklad) *Product {
	return &Product{
		id:         uuid.New(),
		proskladID: model.ID,
		titleRU:    model.Title,
		barcode:    model.Barcode,
		companyID:  model.CompanyID,

		sellingPrice:          model.SellingPrice,
		oldSellingPrice:       model.OldSellingPrice,
		previousPurchasePrice: model.PreviousPurchasePrice,
		additionalPercent:     model.AdditionalPercent,

		hasInventory: model.HasInventory,
		isVirtual:    model.IsVirtual,
		marked:       model.Marked,
		isQuick:      model.IsQuick,

		unitID:   model.UnitID,
		unitName: model.UnitName,
		unitType: model.UnitType,

		brandID: model.VendorID,

		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}
