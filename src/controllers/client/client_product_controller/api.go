package client_product_controller

import (
	"math"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

func (controller *Controller) _get(from, pageSize int, searchQuery string) (*models.ClientProductList, error) {
	var products []entities.ClientProductShort
	var err error

	if len(searchQuery) == 0 {
		products, err = controller.productRepo.GetClient(from, pageSize)
	} else {
		products, err = controller.productRepo.GetClientQuery(from, pageSize, searchQuery)
	}

	if err != nil {
		return nil, ErrorProductGet(err)
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
		return nil, ErrorProductGetCount(err)
	}

	// считаем кол-во страниц
	pageCount := int(math.Ceil(float64(productCount) / float64(20)))

	return &models.ClientProductList{
		List:      productList,
		PageCount: pageCount,
	}, nil
}
