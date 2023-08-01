package product_controller

import (
	"github.com/lowl11/boost"
	"wkey-stock/src/adaptors/product_adaptor"
	"wkey-stock/src/data/models"
)

func (controller Controller) Get(ctx boost.Context) error {
	products, err := controller.products.GetAll()
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, product_adaptor.DtoToModel(products))
}

func (controller Controller) GetByID(ctx boost.Context) error {
	productID := ctx.Param("product-id").String()

	product, err := controller.products.GetByID(productID)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, product.Model())
}

func (controller Controller) Add(ctx boost.Context) error {
	model := models.ProductAdd{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, err)
	}

	newID, err := controller.products.Add(&model)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.CreatedID(ctx, newID)
}

func (controller Controller) UpdateDescription(ctx boost.Context) error {
	productID := ctx.Param("product-id").String()
	if productID == "" {
		return controller.Error(ctx, ErrorProductIDRequired())
	}

	model := models.ProductUpdateDescription{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, err)
	}

	if err := controller.products.UpdateDescription(productID, &model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx)
}

func (controller Controller) UpdateCount(ctx boost.Context) error {
	productID := ctx.Param("product-id").String()
	if productID == "" {
		return controller.Error(ctx, ErrorProductIDRequired())
	}

	model := models.ProductUpdateCount{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, err)
	}

	if err := controller.products.UpdateCount(productID, &model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx)
}

func (controller Controller) Delete(ctx boost.Context) error {
	productID := ctx.Param("product-id").String()

	if productID == "" {
		return controller.Error(ctx, ErrorProductIDRequired())
	}

	if err := controller.products.RemoveByID(productID); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx)
}
