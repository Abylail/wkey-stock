package product_controller

import (
	"github.com/lowl11/boost"
	"wkey-stock/src/adaptors/product_adaptor"
	"wkey-stock/src/data/models"
)

func (controller Controller) Get(ctx boost.Context) error {
	products, err := controller.products.GetAll(ctx.Context())
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, product_adaptor.DtoToModel(products))
}

func (controller Controller) GetByID(ctx boost.Context) error {
	productID := ctx.Param("product-id").String()

	product, err := controller.products.GetByID(ctx.Context(), productID)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, product.Model())
}

func (controller Controller) Add(ctx boost.Context) error {
	model := models.ProductProsklad{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, err)
	}

	newID, err := controller.products.Add(ctx.Context(), &model)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.CreatedID(ctx, newID)
}

func (controller Controller) UpdateProsklad(ctx boost.Context) error {
	proskladID := ctx.Param("prosklad-id").Int()
	if proskladID == 0 {
		return controller.Error(ctx, ErrorProskladIDRequired())
	}

	model := models.ProductProsklad{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, err)
	}

	if err := controller.products.UpdateProsklad(ctx.Context(), proskladID, &model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx)
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

	if err := controller.products.UpdateDescription(ctx.Context(), productID, &model); err != nil {
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

	if err := controller.products.UpdateCount(ctx.Context(), productID, &model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx)
}

func (controller Controller) UpdateImages(ctx boost.Context) error {
	productID := ctx.Param("product-id").String()
	if productID == "" {
		return controller.Error(ctx, ErrorProductIDRequired())
	}

	model := models.ProductUpdateImages{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, err)
	}

	if err := controller.products.UpdateImages(ctx.Context(), productID, &model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx)
}

func (controller Controller) Delete(ctx boost.Context) error {
	productID := ctx.Param("product-id").String()

	if productID == "" {
		return controller.Error(ctx, ErrorProductIDRequired())
	}

	if err := controller.products.RemoveByID(ctx.Context(), productID); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx)
}
