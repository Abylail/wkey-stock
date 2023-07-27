package product_controller

import (
	"github.com/lowl11/boost"
	"wkey-stock/src/data/models"
)

func (controller *Controller) GetAdminREST(ctx boost.Context) error {
	searchQuery := ctx.QueryParam("query").String()
	categoryKey := ctx.QueryParam("category").String()
	subcategoryKey := ctx.QueryParam("subcategory").String()
	page := ctx.QueryParam("page").Int()
	if page == 0 {
		page = 1
	}

	pageSize := 20
	from := (page - 1) * pageSize

	products, err := controller._getAdmin(from, pageSize, searchQuery, categoryKey, subcategoryKey)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, products)
}

func (controller *Controller) GetAdminSingleREST(ctx boost.Context) error {
	productID := ctx.Param("id").Int()
	if productID == 0 {
		//return controller.NotFound(ctx, errors.AdminProductNotFound)
		return controller.NotFound(ctx)
	}

	product, err := controller._getAdminSingle(productID)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, product)
}

func (controller *Controller) GetClientREST(ctx boost.Context) error {
	searchQuery := ctx.QueryParam("query").String()
	page := ctx.QueryParam("page").Int()
	if page == 0 {
		page = 1
	}

	pageSize := 20
	from := (page - 1) * pageSize

	products, err := controller._getClient(from, pageSize, searchQuery)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, products)
}

func (controller *Controller) UpdateProductREST(ctx boost.Context) error {
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

func (controller *Controller) UploadProductREST(ctx boost.Context) error {
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

func (controller *Controller) GetBrandREST(ctx boost.Context) error {
	searchQuery := ctx.QueryParam("query").String()

	brands, err := controller._getBrand(searchQuery)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, brands)
}

func (controller *Controller) GetBrandSingleREST(ctx boost.Context) error {
	brandID := ctx.Param("id").Int()
	if brandID == 0 {
		return controller.Error(ctx, ErrorBrandParamRequired("id"))
	}

	brands, err := controller._getBrandSingle(brandID)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, brands)
}

func (controller *Controller) AddBrandREST(ctx boost.Context) error {
	model := models.BrandAdd{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, ErrorBrandBind(err))
	}

	controller.filterBrandAdd(&model)

	if err := controller._addBrand(&model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) UpdateBrandREST(ctx boost.Context) error {
	brandID := ctx.Param("id").Int()
	if brandID == 0 {
		return controller.Error(ctx, ErrorBrandParamRequired("id"))
	}

	model := models.BrandUpdate{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, ErrorBrandBind(err))
	}

	if err := controller._updateBrand(brandID, &model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) UploadBrandREST(ctx boost.Context) error {
	brandID := ctx.Param("id").Int()
	if brandID == 0 {
		return controller.Error(ctx, ErrorBrandParamRequired("id"))
	}

	model := models.BrandUpload{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, ErrorBrandBind(err))
	}

	imagePath, err := controller._uploadBrand(brandID, &model)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, imagePath)
}

func (controller *Controller) DeleteBrandREST(ctx boost.Context) error {
	brandID := ctx.Param("id").Int()
	if brandID == 0 {
		return controller.Error(ctx, ErrorBrandParamRequired("id"))
	}

	if err := controller._deleteBrand(brandID); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}
