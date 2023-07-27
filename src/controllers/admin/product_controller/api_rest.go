package product_controller

import (
	"github.com/lowl11/boost"
	"wkey-stock/src/data/models"
)

func (controller Controller) GetREST(ctx boost.Context) error {
	searchQuery := ctx.QueryParam("query").String()
	categoryKey := ctx.QueryParam("category").String()
	subcategoryKey := ctx.QueryParam("subcategory").String()
	page := ctx.QueryParam("page").Int()
	if page == 0 {
		page = 1
	}

	pageSize := 20
	from := (page - 1) * pageSize

	products, err := controller._get(from, pageSize, searchQuery, categoryKey, subcategoryKey)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, products)
}

func (controller Controller) GetSingleREST(ctx boost.Context) error {
	productID := ctx.Param("id").Int()
	if productID == 0 {
		//return controller.NotFound(ctx, errors.AdminProductNotFound)
		return controller.NotFound(ctx)
	}

	product, err := controller._getSingle(productID)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, product)
}

func (controller Controller) UpdateProductREST(ctx boost.Context) error {
	productID := ctx.Param("id").Int()
	if productID == 0 {
		return controller.Error(ctx, ErrorProductParamRequired("id"))
	}

	model := models.ProductUpdate{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, ErrorProductBind(err))
	}

	if err := controller._update(productID, &model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller Controller) UploadProductREST(ctx boost.Context) error {
	productID := ctx.Param("id").Int()
	if productID == 0 {
		return controller.Error(ctx, ErrorProductParamRequired("id"))
	}

	model := models.ProductUpload{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, ErrorProductBind(err))
	}

	if err := controller._upload(productID, &model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}
