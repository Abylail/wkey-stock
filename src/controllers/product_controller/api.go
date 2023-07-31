package product_controller

import (
	"github.com/lowl11/boost"
	"wkey-stock/src/adaptors/product_adaptor"
)

func (controller Controller) Get(ctx boost.Context) error {
	products, err := controller.products.GetAll()
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, product_adaptor.DtoToModel(products))
}
