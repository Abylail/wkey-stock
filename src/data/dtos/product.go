package dtos

import (
	"github.com/google/uuid"
	"strings"
	"time"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
	"wkey-stock/src/enums/languages"
)

type Product struct {
	// common
	id               uuid.UUID
	proskladID       int
	titleRU          string
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
	descriptionRU   *string
	descriptionKZ   *string
	count           int
	price           float32
	vendorCode      string
	primaryImage    *string
	secondaryImages []string

	createdAt time.Time
	updatedAt time.Time
}

func NewProduct(entity *entities.Product) *Product {
	id, _ := uuid.Parse(entity.ID)

	return &Product{
		// common
		id:               id,
		proskladID:       entity.ProskladID,
		titleRU:          entity.TitleRU,
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
		descriptionRU:   entity.DescriptionRU,
		descriptionKZ:   entity.DescriptionKZ,
		count:           entity.Count,
		price:           entity.Price,
		vendorCode:      entity.VendorCode,
		primaryImage:    entity.PrimaryImage,
		secondaryImages: getSecondaryImages(entity.SecondaryImages),

		createdAt: entity.CreatedAt,
		updatedAt: entity.UpdatedAt,
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

func (product *Product) EditProsklad(models *models.ProductProsklad) {
	defer product.updateDate()

	// common
	product.titleRU = models.Title
	product.barcode = models.Barcode
	product.companyID = models.CompanyID
	product.itemCategoryName = models.ItemCategoryName

	// price
	product.sellingPrice = models.SellingPrice
	product.oldSellingPrice = models.OldSellingPrice
	product.previousPurchasePrice = models.PreviousPurchasePrice
	product.additionalPercent = models.AdditionalPercent

	// flags
	product.hasInventory = models.HasInventory
	product.isVirtual = models.IsVirtual
	product.marked = models.Marked
	product.isQuick = models.IsQuick

	// unit
	product.unitID = models.UnitID
	product.unitName = models.UnitName
	product.unitType = models.UnitType

	// vendor
	product.brandID = models.VendorID
}

func (product *Product) EditDescription(description, language string) {
	defer product.updateDate()

	if language == languages.KZ {
		product.descriptionKZ = &description
		return
	}

	product.descriptionRU = &description
}

func (product *Product) EditCount(count int) {
	defer product.updateDate()

	product.count = count
}

func (product *Product) EditImages(primary string, secondary []string) {
	defer product.updateDate()

	product.primaryImage = &primary
	product.secondaryImages = secondary
}

func (product *Product) Model() models.Product {
	return models.Product{
		ID:              product.id.String(),
		ProskladID:      product.proskladID,
		TitleRU:         product.titleRU,
		DescriptionRU:   product.Description(languages.RU),
		DescriptionKZ:   product.Description(languages.KZ),
		Count:           product.count,
		Price:           product.price,
		PrimaryImage:    product.primaryImage,
		SecondaryImages: product.secondaryImages,
		Type:            product.vendorCode,
	}
}

func (product *Product) Entity() entities.Product {
	return entities.Product{
		ID:               product.id.String(),
		ProskladID:       product.proskladID,
		TitleRU:          product.titleRU,
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

		DescriptionRU:   product.descriptionRU,
		DescriptionKZ:   product.descriptionKZ,
		Count:           product.count,
		PrimaryImage:    product.primaryImage,
		SecondaryImages: strings.Join(product.secondaryImages, ","),

		CreatedAt: product.createdAt,
		UpdatedAt: product.updatedAt,
	}
}

func (product *Product) updateDate() {
	product.updatedAt = time.Now()
}

func getSecondaryImages(images string) []string {
	imagesArray := strings.Split(images, ",")
	if len(imagesArray) == 1 && imagesArray[0] == "" {
		return []string{}
	}

	return imagesArray
}
