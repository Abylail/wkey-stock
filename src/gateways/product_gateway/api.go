package product_gateway

import (
	"wkey-stock/src/adaptors/product_adaptor"
	"wkey-stock/src/data/dtos"
)

func (gateway Gateway) GetAll() ([]dtos.Product, error) {
	products, err := gateway.productRepo.GetAll()
	if err != nil {
		return nil, ErrorGetAllProducts(err)
	}

	return product_adaptor.EntityToDTO(products), nil
}
