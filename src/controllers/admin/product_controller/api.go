package product_controller

import (
	"math"
	"wkey-stock/src/adaptors/product_adaptor"
	"wkey-stock/src/adaptors/product_category_adaptor"
	"wkey-stock/src/data/dtos"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

func (controller Controller) _get(page int, searchQuery string) (*dtos.ProductList, error) {
	pageSize := 20
	from := (page - 1) * pageSize

	var products []entities.AdminProductGet
	var err error

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
	pageCount := int(math.Ceil(float64(productCount / 20)))

	// получаем список категорий
	categoryPairs, err := controller.productRepo.GetSubCategoryPairs(productIDs)
	if err != nil {
		return nil, ErrorProductGetPairs(err)
	}

	categories := product_category_adaptor.EntityToDTO(categoryPairs)
	productList := product_adaptor.EntityToDTO(products)

	for i := 0; i < len(productList); i++ {
		productList[i].EditImage(images)
	}

	return dtos.NewProductList(pageCount, productList, categories), nil
}

func (controller Controller) _getSingle(productID int) (*dtos.Product, error) {
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

	productDTO := dtos.NewProduct(product)
	productDTO.EditImage(images)
	productDTO.SaveCategories(product_category_adaptor.EntityToDTO(categoryPairs))

	return productDTO, nil
}

func (controller Controller) _update(productID int, model *models.ProductUpdate) error {
	if err := controller.productRepo.Update(productID, model); err != nil {
		return ErrorProductUpdate(err)
	}

	return nil
}

func (controller Controller) _upload(productID int, model *models.ProductUpload) error {
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
