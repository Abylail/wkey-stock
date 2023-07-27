package dtos

import (
	"time"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
	"wkey-stock/src/enums/languages"
)

type Product struct {
	id                int
	title             string
	price             int
	vendorCode        string
	barcode           string
	unitName          string
	createdAt         time.Time
	updatedAt         time.Time
	additionalPercent any
	descriptionRU     *string
	descriptionKZ     *string
	count             int
	brandTitle        string

	images     []string
	categories []ProductCategory
}

func NewProduct(entity *entities.AdminProductGet) *Product {
	return &Product{
		id:                entity.ID,
		title:             entity.Title,
		price:             entity.Price,
		vendorCode:        entity.VendorCode,
		barcode:           entity.Barcode,
		unitName:          entity.UnitName,
		createdAt:         entity.CreatedAt,
		updatedAt:         entity.UpdatedAt,
		additionalPercent: entity.AdditionalPercent,
		descriptionRU:     entity.DescriptionRU,
		descriptionKZ:     entity.DescriptionKZ,
		count:             entity.Count,
		brandTitle:        entity.BrandTitle,
	}
}

func (product Product) ID() int {
	return product.id
}

func (product Product) Title() string {
	return product.title
}

func (product Product) Price() int {
	return product.price
}

func (product Product) VendorCode() string {
	return product.vendorCode
}

func (product Product) Barcode() string {
	return product.barcode
}

func (product Product) UnitName() string {
	return product.unitName
}

func (product Product) CreateAt() time.Time {
	return product.createdAt
}

func (product Product) UpdatedAt() time.Time {
	return product.updatedAt
}

func (product Product) AdditionalPercent() any {
	return product.additionalPercent
}

func (product Product) Description(language string) *string {
	if language == languages.KZ {
		return product.descriptionKZ
	}

	return product.descriptionRU
}

func (product Product) Count() int {
	return product.count
}

func (product Product) BrandTitle() string {
	return product.brandTitle
}

func (product Product) EditImage(images []entities.ProductImageGet) {
	productImages := make([]entities.ProductImageGet, 0)
	for _, image := range images {
		if image.ProductID != product.ID() {
			continue
		}

		productImages = append(productImages, image)
	}

	product.images = make([]string, 0, len(productImages))
	for _, image := range productImages {
		product.images = append(product.images, image.Path)
	}
}

func (product Product) SaveCategories(categories []ProductCategory) {
	product.categories = categories
}

func (product Product) Categories() []ProductCategory {
	return product.categories
}

func (product Product) Model(categories []ProductCategory) models.AdminProductItem {
	categoriesModel := make([]models.ProductCategoryPair, 0, len(categories))
	for _, category := range categories {
		categoriesModel = append(categoriesModel, category.Model())
	}

	return models.AdminProductItem{
		ID:                product.ID(),
		Title:             product.Title(),
		Price:             product.Price(),
		VendorCode:        product.VendorCode(),
		Barcode:           product.Barcode(),
		UnitName:          product.UnitName(),
		CreatedAt:         product.CreateAt(),
		UpdatedAt:         product.UpdatedAt(),
		AdditionalPercent: product.AdditionalPercent(),
		DescriptionRU:     product.Description(languages.RU),
		DescriptionKZ:     product.Description(languages.KZ),
		Count:             product.Count(),
		BrandTitle:        product.BrandTitle(),
		Categories:        categoriesModel,
	}
}
