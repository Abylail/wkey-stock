package product_controller

import (
	"github.com/labstack/echo/v4"
	"strconv"
	"wkey-stock/src/data/errors"
	"wkey-stock/src/data/models"
)

func (controller *Controller) GetAdminREST(ctx echo.Context) error {
	searchQuery := ctx.QueryParam("query")
	categoryKey := ctx.QueryParam("category")
	subcategoryKey := ctx.QueryParam("subcategory")
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
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

func (controller *Controller) GetAdminSingleREST(ctx echo.Context) error {
	productID, _ := strconv.Atoi(ctx.Param("id"))
	if productID == 0 {
		return controller.NotFound(ctx, errors.AdminProductNotFound)
	}

	product, err := controller._getAdminSingle(productID)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, product)
}

func (controller *Controller) GetClientREST(ctx echo.Context) error {
	searchQuery := ctx.QueryParam("query")
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
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

func (controller *Controller) UpdateProductREST(ctx echo.Context) error {
	productID, _ := strconv.Atoi(ctx.Param("id"))
	if productID == 0 {
		return controller.Error(ctx, errors.ProductUpdateParam)
	}

	model := models.ProductUpdate{}
	if err := ctx.Bind(&model); err != nil {
		return controller.Error(ctx, errors.ProductUpdateBind.With(err))
	}

	if err := controller.validateProductUpdate(&model); err != nil {
		return controller.Error(ctx, errors.ProductUpdateValidate.With(err))
	}

	if err := controller._update(productID, &model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) UploadProductREST(ctx echo.Context) error {
	productID, _ := strconv.Atoi(ctx.Param("id"))
	if productID == 0 {
		return controller.Error(ctx, errors.ProductUpdateParam)
	}

	model := models.ProductUpload{}
	if err := ctx.Bind(&model); err != nil {
		return controller.Error(ctx, errors.ProductUploadBind.With(err))
	}

	if err := controller.validateProductUpload(&model); err != nil {
		return controller.Error(ctx, errors.ProductUploadValidate.With(err))
	}

	if err := controller._upload(productID, &model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) GetBrandREST(ctx echo.Context) error {
	searchQuery := ctx.QueryParam("query")

	brands, err := controller._getBrand(searchQuery)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, brands)
}

func (controller *Controller) GetBrandSingleREST(ctx echo.Context) error {
	brandID, _ := strconv.Atoi(ctx.Param("id"))
	if brandID == 0 {
		return errors.BrandGetParam
	}

	brands, err := controller._getBrandSingle(brandID)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, brands)
}

func (controller *Controller) AddBrandREST(ctx echo.Context) error {
	model := models.BrandAdd{}
	if err := ctx.Bind(&model); err != nil {
		return controller.Error(ctx, errors.BrandAddBind.With(err))
	}

	if err := controller.validateBrandAdd(&model); err != nil {
		return controller.Error(ctx, errors.BrandAddValidate.With(err))
	}

	controller.filterBrandAdd(&model)

	if err := controller._addBrand(&model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) UpdateBrandREST(ctx echo.Context) error {
	brandID, _ := strconv.Atoi(ctx.Param("id"))
	if brandID == 0 {
		return controller.Error(ctx, errors.BrandUpdateParam)
	}

	model := models.BrandUpdate{}
	if err := ctx.Bind(&model); err != nil {
		return controller.Error(ctx, errors.BrandUpdateBind.With(err))
	}

	if err := controller.validateBrandUpdate(&model); err != nil {
		return controller.Error(ctx, errors.BrandUpdateValidate.With(err))
	}

	if err := controller._updateBrand(brandID, &model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) UploadBrandREST(ctx echo.Context) error {
	brandID, _ := strconv.Atoi(ctx.Param("id"))
	if brandID == 0 {
		return controller.Error(ctx, errors.BrandUploadParam)
	}

	model := models.BrandUpload{}
	if err := ctx.Bind(&model); err != nil {
		return controller.Error(ctx, errors.BrandUploadBind.With(err))
	}

	if err := controller.validateBrandUpload(&model); err != nil {
		return controller.Error(ctx, errors.BrandUploadValidate.With(err))
	}

	imagePath, err := controller._uploadBrand(brandID, &model)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, imagePath)
}

func (controller *Controller) DeleteBrandREST(ctx echo.Context) error {
	brandID, _ := strconv.Atoi(ctx.Param("id"))
	if brandID == 0 {
		return errors.BrandDeleteParam
	}

	if err := controller._deleteBrand(brandID); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}
