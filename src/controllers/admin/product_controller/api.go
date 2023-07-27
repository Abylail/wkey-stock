package product_controller

import (
	"fmt"
	"math"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

func (controller *Controller) _get(from, pageSize int, searchQuery, categoryKey string, subcategoryKey string) (*models.AdminProductGet, error) {
	var products []entities.AdminProductGet
	var err error

	// Ищу сабкатегорию если она есть
	var subcategory *entities.SubCategory
	if subcategoryKey != "" {
		fmt.Println("subcategoryKey", subcategoryKey)
		subcategory, err = controller.subCategoryRepo.GetByCode(subcategoryKey)
		if err != nil {
			return nil, ErrorAdminProductGet(err)
		}
	}

	fmt.Println(subcategory)

	// получаем сами продукты
	if len(searchQuery) == 0 {
		products, err = controller.productRepo.GetAdmin(from, pageSize)
	} else {
		products, err = controller.productRepo.GetAdminByQuery(from, pageSize, searchQuery)
	}
	if err != nil {
		return nil, ErrorAdminProductGet(err)
	}

	// получаем массив с ID продуктов
	productIDs := make([]int, 0, len(products))
	for _, product := range products {
		productIDs = append(productIDs, product.ID)
	}

	// получаем картинки продуктов
	images, err := controller.productRepo.GetImages(productIDs)
	if err != nil {
		return nil, ErrorProductGetImages(err)
	}

	imagesList := make([]entities.ProductImageGet, 0, len(images))
	for _, image := range images {
		imagesList = append(imagesList, image)
	}

	// получаем кол-во продуктов
	var productCount int
	if len(searchQuery) == 0 {
		productCount, err = controller.productRepo.Count()
	} else {
		productCount, err = controller.productRepo.CountQuery(searchQuery)
	}
	if err != nil {
		return nil, ErrorAdminProductGetCount(err)
	}

	// считаем кол-во страниц
	pageCount := int(math.Ceil(float64(productCount) / float64(20)))

	// получаем список категорий
	categoryPairs, err := controller.productRepo.GetSubCategoryPairs(productIDs)
	if err != nil {
		return nil, ErrorProductGetPairs(err)
	}

	productList := make([]models.AdminProductItem, 0, len(products))
	for _, product := range products {
		categories := make([]models.ProductCategoryPair, 0, len(categoryPairs))
		for _, category := range categoryPairs {
			categories = append(categories, models.ProductCategoryPair{
				SubCategoryCode: category.SubCategoryCode,
				SubCategoryName: category.SubCategoryName,
				CategoryCode:    category.CategoryCode,
				CategoryName:    category.CategoryName,
			})
		}

		productList = append(productList, models.AdminProductItem{
			ID:                product.ID,
			Title:             product.Title,
			Price:             product.Price,
			VendorCode:        product.VendorCode,
			Barcode:           product.Barcode,
			UnitName:          product.UnitName,
			Categories:        categories,
			CreatedAt:         product.CreatedAt,
			UpdatedAt:         product.UpdatedAt,
			AdditionalPercent: product.AdditionalPercent,
			DescriptionRU:     product.DescriptionRU,
			DescriptionKZ:     product.DescriptionKZ,
			Count:             product.Count,
			BrandTitle:        product.BrandTitle,
		})
	}

	for i := 0; i < len(productList); i++ {
		productID := productList[i].ID
		productImages := make([]entities.ProductImageGet, 0)
		for _, image := range imagesList {
			if image.ProductID != productID {
				continue
			}

			productImages = append(productImages, image)
		}

		productList[i].Images = make([]string, 0, len(productImages))

		for _, image := range productImages {
			productList[i].Images = append(productList[i].Images, image.Path)
		}
	}

	return &models.AdminProductGet{
		PageCount: pageCount,
		List:      productList,
	}, nil
}

func (controller *Controller) _getSingle(productID int) (*models.AdminProductItem, error) {
	product, err := controller.productRepo.GetAdminByID(productID)
	if err != nil {
		return nil, ErrorAdminProductGet(err)
	}

	if product == nil {
		return nil, ErrorAdminProductNotFound(productID)
	}

	images, err := controller.productRepo.GetImages([]int{productID})
	if err != nil {
		return nil, ErrorProductGetImages(err)
	}

	// получаем список категорий
	categoryPairs, err := controller.productRepo.GetSubCategoryPairs([]int{productID})
	if err != nil {
		return nil, ErrorProductGetPairs(err)
	}

	categories := make([]models.ProductCategoryPair, 0, len(categoryPairs))
	for _, pair := range categoryPairs {
		categories = append(categories, models.ProductCategoryPair{
			SubCategoryCode: pair.SubCategoryCode,
			SubCategoryName: pair.SubCategoryName,
			CategoryCode:    pair.CategoryCode,
			CategoryName:    pair.CategoryName,
		})
	}

	productImages := make([]string, 0, len(images))
	for _, image := range images {
		productImages = append(productImages, image.Path)
	}

	return &models.AdminProductItem{
		ID:                product.ID,
		Title:             product.Title,
		Price:             product.Price,
		VendorCode:        product.VendorCode,
		Barcode:           product.Barcode,
		UnitName:          product.UnitName,
		Categories:        categories,
		CreatedAt:         product.CreatedAt,
		UpdatedAt:         product.UpdatedAt,
		AdditionalPercent: product.AdditionalPercent,
		DescriptionRU:     product.DescriptionRU,
		DescriptionKZ:     product.DescriptionKZ,
		Count:             product.Count,
		BrandTitle:        product.BrandTitle,
		Images:            productImages,
	}, nil
}

func (controller *Controller) _update(productID int, model *models.ProductUpdate) error {
	if err := controller.productRepo.Update(productID, model); err != nil {
		return ErrorProductUpdate(err)
	}

	return nil
}

func (controller *Controller) _upload(productID int, model *models.ProductUpload) error {
	if len(model.Images) == 0 {
		return nil
	}

	pathList, err := controller.image.UploadProductImages(productID, model)
	if err != nil {
		return ErrorProductUpdateFileImages(err)
	}

	if err = controller.productRepo.UpdateImages(productID, model, pathList); err != nil {
		return ErrorProductUpdateImages(err)
	}

	return nil
}
