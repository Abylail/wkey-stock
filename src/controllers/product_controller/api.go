package product_controller

import (
	"github.com/lowl11/lazy-collection/array"
	"github.com/lowl11/lazy-collection/type_list"
	"github.com/lowl11/lazylog/layers"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/errors"
	"wkey-stock/src/data/models"
	"wkey-stock/src/definition"
)

func (controller *Controller) _getAdmin(from, to int, searchQuery, categoryKey string) ([]models.AdminProductGet, *models.Error) {
	logger := definition.Logger

	// todo: что то придумать с этим
	_ = categoryKey

	var products []entities.AdminProductGet
	var err error

	// получаем сами продукты
	if len(searchQuery) == 0 {
		products, err = controller.productRepo.GetAdmin(from, to)
	} else {
		products, err = controller.productRepo.GetAdminByQuery(from, to, searchQuery)
	}
	if err != nil {
		logger.Error(err, "Get products list error", layers.Database)
		return nil, errors.AdminProductGet.With(err)
	}

	// получаем массив с ID продуктов
	productIDs := make([]int, 0, len(products))
	array.NewWithList[entities.AdminProductGet](products...).Each(func(item entities.AdminProductGet) {
		productIDs = append(productIDs, item.ID)
	})

	// получаем картинки продуктов
	images, err := controller.productRepo.GetImages(productIDs)
	if err != nil {
		logger.Error(err, "Get product images error", layers.Database)
		return nil, errors.ProductImagesGet.With(err)
	}

	imagesList := array.NewWithList[entities.ProductImageGet](images...)

	productList := type_list.NewWithList[entities.AdminProductGet, models.AdminProductGet](products...).
		Select(func(item entities.AdminProductGet) models.AdminProductGet {
			return models.AdminProductGet{
				ID:                item.ID,
				Title:             item.Title,
				Price:             item.Price,
				VendorCode:        item.VendorCode,
				Barcode:           item.Barcode,
				UnitName:          item.UnitName,
				CategoryID:        item.CategoryID,
				CategoryName:      item.CategoryName,
				CreatedAt:         item.CreatedAt,
				UpdatedAt:         item.UpdatedAt,
				AdditionalPercent: item.AdditionalPercent,
				DescriptionRU:     item.DescriptionRU,
				DescriptionKZ:     item.DescriptionKZ,
				Count:             item.Count,
			}
		}).Slice()

	for i := 0; i < len(productList); i++ {
		productID := productList[i].ID
		productImages := imagesList.Where(func(image entities.ProductImageGet) bool {
			return image.ProductID == productID
		})

		productList[i].Images = make([]string, 0, productImages.Size())

		productImages.Each(func(image entities.ProductImageGet) {
			productList[i].Images = append(productList[i].Images, image.Path)
		})
	}

	return productList, nil
}
