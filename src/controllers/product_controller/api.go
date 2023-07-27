package product_controller

import (
	"fmt"
	"math"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

func (controller *Controller) _getAdmin(from, pageSize int, searchQuery, categoryKey string, subcategoryKey string) (*models.AdminProductGet, error) {
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

func (controller *Controller) _getAdminSingle(productID int) (*models.AdminProductItem, error) {
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

func (controller *Controller) _getBrand(searchQuery string) ([]models.BrandGet, error) {
	var brands []entities.Brand
	var err error

	if len(searchQuery) == 0 {
		brands, err = controller.brandRepo.GetAll()
	} else {
		brands, err = controller.brandRepo.GetByQuery(searchQuery)
	}
	if err != nil {
		return nil, ErrorBrandGetList(err)
	}

	list := make([]models.BrandGet, 0, len(brands))
	for _, brand := range brands {
		list = append(list, models.BrandGet{
			ID:    brand.ProskladID,
			Title: brand.Title,
			Image: brand.Image,
		})
	}

	return list, nil
}

func (controller *Controller) _getBrandSingle(id int) (*models.BrandGet, error) {
	brand, err := controller.brandRepo.GetByID(id)
	if err != nil {
		return nil, ErrorBrandGetByID(err)
	}

	return &models.BrandGet{
		ID:    brand.ProskladID,
		Title: brand.Title,
		Image: brand.Image,
	}, nil
}

func (controller *Controller) _addBrand(model *models.BrandAdd) error {
	brand, err := controller.brandRepo.GetByTitle(model.Title)
	if err != nil {
		return ErrorBrandGetByTitle(model.Title)
	}

	if brand != nil {
		return ErrorBrandAlreadyExist(model.Title)
	}

	if err = controller.brandRepo.Create(model); err != nil {
		return ErrorBrandAdd(err)
	}

	return nil
}

func (controller *Controller) _updateBrand(id int, model *models.BrandUpdate) error {
	if err := controller.brandRepo.Update(id, model); err != nil {
		return ErrorBrandUpdate(err)
	}

	return nil
}

func (controller *Controller) _uploadBrand(brandID int, model *models.BrandUpload) (string, error) {
	brand, err := controller.brandRepo.GetByID(brandID)
	if err != nil {
		return "", ErrorBrandGetByID(err)
	}

	if brand == nil {
		return "", ErrorBrandNotFound(brandID)
	}

	imagePath, err := controller.image.UploadBrandIcon(brandID, model.Image.Name, model.Image.Buffer)
	if err != nil {
		return "", ErrorBrandUpdateFileIcon(err)
	}

	if err = controller.brandRepo.UpdateIcon(brandID, imagePath); err != nil {
		return "", ErrorBrandUpdateIcon(err)
	}

	return imagePath, nil
}

func (controller *Controller) _deleteBrand(id int) error {
	if err := controller.brandRepo.DeleteByID(id); err != nil {
		return ErrorBrandDelete(err)
	}

	return nil
}

func (controller *Controller) _getClient(from, pageSize int, searchQuery string) (*models.ClientProductList, error) {
	var products []entities.ClientProductShort
	var err error

	if len(searchQuery) == 0 {
		products, err = controller.productRepo.GetClient(from, pageSize)
	} else {
		products, err = controller.productRepo.GetClientQuery(from, pageSize, searchQuery)
	}

	if err != nil {
		return nil, ErrorClientProductGet(err)
	}

	productList := make([]models.ClientProductItemShort, 0, len(products))
	for _, product := range products {
		productList = append(productList, models.ClientProductItemShort{
			ID:         product.ID,
			Title:      product.Title,
			Price:      product.Price,
			VendorCode: product.VendorCode,
			Count:      product.Count,
		})
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

	for i := 0; i < len(productList); i++ {
		productID := productList[i].ID

		productImages := make([]entities.ProductImageGet, 0)
		for _, image := range images {
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

	// получаем кол-во продуктов
	var productCount int
	if len(searchQuery) == 0 {
		productCount, err = controller.productRepo.GetClientCount()
	} else {
		productCount, err = controller.productRepo.GetClientCountQuery(searchQuery)
	}
	if err != nil {
		return nil, ErrorClientProductGetCount(err)
	}

	// считаем кол-во страниц
	pageCount := int(math.Ceil(float64(productCount) / float64(20)))

	return &models.ClientProductList{
		List:      productList,
		PageCount: pageCount,
	}, nil
}
