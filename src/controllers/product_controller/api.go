package product_controller

import (
	"github.com/lowl11/lazy-collection/array"
	"github.com/lowl11/lazy-collection/type_list"
	"github.com/lowl11/lazylog/layers"
	"math"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/errors"
	"wkey-stock/src/data/models"
	"wkey-stock/src/definition"
)

func (controller *Controller) _getAdmin(from, to int, searchQuery, categoryKey string) (*models.AdminProductGet, *models.Error) {
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

	// получаем кол-во продуктов
	var productCount int
	if len(searchQuery) == 0 {
		productCount, err = controller.productRepo.Count()
	} else {
		productCount, err = controller.productRepo.CountQuery(searchQuery)
	}
	if err != nil {
		logger.Error(err, "Get products count error", layers.Database)
		return nil, errors.AdminProductCountGet.With(err)
	}

	// считаем кол-во страниц
	pageCount := int(math.Ceil(float64(productCount / 20)))

	productList := type_list.NewWithList[entities.AdminProductGet, models.AdminProductItem](products...).
		Select(func(item entities.AdminProductGet) models.AdminProductItem {
			return models.AdminProductItem{
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

	return &models.AdminProductGet{
		PageCount: pageCount,
		List:      productList,
	}, nil
}

func (controller *Controller) _getAdminSingle(productID int) (*models.AdminProductItem, *models.Error) {
	logger := definition.Logger

	product, err := controller.productRepo.GetAdminByID(productID)
	if err != nil {
		logger.Error(err, "Get admin product by id error", layers.Database)
		return nil, errors.AdminProductGet.With(err)
	}

	if product == nil {
		return nil, errors.AdminProductNotFound
	}

	images, err := controller.productRepo.GetImages([]int{productID})
	if err != nil {
		logger.Error(err, "Get product images error", layers.Database)
		return nil, errors.ProductImagesGet.With(err)
	}

	return &models.AdminProductItem{
		ID:                product.ID,
		Title:             product.Title,
		Price:             product.Price,
		VendorCode:        product.VendorCode,
		Barcode:           product.Barcode,
		UnitName:          product.UnitName,
		CategoryID:        product.CategoryID,
		CategoryName:      product.CategoryName,
		CreatedAt:         product.CreatedAt,
		UpdatedAt:         product.UpdatedAt,
		AdditionalPercent: product.AdditionalPercent,
		DescriptionRU:     product.DescriptionRU,
		DescriptionKZ:     product.DescriptionKZ,
		Count:             product.Count,
		Images: type_list.NewWithList[entities.ProductImageGet, string](images...).
			Select(func(item entities.ProductImageGet) string {
				return item.Path
			}).Slice(),
	}, nil
}

func (controller *Controller) _getBrand(searchQuery string) ([]models.BrandGet, *models.Error) {
	logger := definition.Logger

	var brands []entities.BrandGet
	var err error

	if len(searchQuery) == 0 {
		brands, err = controller.brandRepo.GetAll()
	} else {
		brands, err = controller.brandRepo.GetByQuery(searchQuery)
	}
	if err != nil {
		logger.Error(err, "Get list of brands error", layers.Database)
		return nil, errors.BrandGetList.With(err)
	}

	return type_list.NewWithList[entities.BrandGet, models.BrandGet](brands...).
		Select(func(item entities.BrandGet) models.BrandGet {
			return models.BrandGet{
				ID:    item.ID,
				Title: item.Title,
				Image: item.Image,
			}
		}).Slice(), nil
}

func (controller *Controller) _addBrand(model *models.BrandAdd) *models.Error {
	logger := definition.Logger

	brand, err := controller.brandRepo.GetByTitle(model.Title)
	if err != nil {
		logger.Error(err, "Get brand by title error", layers.Database)
		return errors.BrandGetByTitle.With(err)
	}

	if brand != nil {
		return errors.BrandAlreadyExist
	}

	if err = controller.brandRepo.Create(model); err != nil {
		logger.Error(err, "Create brand error", layers.Database)
		return errors.BrandAdd.With(err)
	}

	return nil
}

func (controller *Controller) _updateBrand(id int, model *models.BrandUpdate) *models.Error {
	logger := definition.Logger

	if err := controller.brandRepo.Update(id, model); err != nil {
		logger.Error(err, "Update brand error", layers.Database)
		return errors.BrandUpdate.With(err)
	}

	return nil
}

func (controller *Controller) _deleteBrand(id int) *models.Error {
	logger := definition.Logger

	if err := controller.brandRepo.Delete(id); err != nil {
		logger.Error(err, "Delete brand error", layers.Database)
		return errors.BrandDelete.With(err)
	}
	return nil
}
