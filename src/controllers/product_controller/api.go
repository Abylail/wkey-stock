package product_controller

import (
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
	_ = searchQuery

	products, err := controller.productRepo.GetAdmin(from, to)
	if err != nil {
		logger.Error(err, "Get products list error", layers.Database)
		return nil, errors.AdminProductGet.With(err)
	}

	return type_list.NewWithList[entities.AdminProductGet, models.AdminProductGet](products...).
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
		}).Slice(), nil
}
