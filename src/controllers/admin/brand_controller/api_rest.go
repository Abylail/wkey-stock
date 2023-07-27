package brand_controller

import (
	"github.com/lowl11/boost"
	"wkey-stock/src/data/models"
)

func (controller *Controller) GetREST(ctx boost.Context) error {
	searchQuery := ctx.QueryParam("query").String()

	brands, err := controller._get(searchQuery)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, brands)
}

func (controller *Controller) GetSingleREST(ctx boost.Context) error {
	brandID := ctx.Param("id").Int()
	if brandID == 0 {
		return controller.Error(ctx, ErrorBrandParamRequired("id"))
	}

	brands, err := controller._getSingle(brandID)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, brands)
}

func (controller *Controller) AddREST(ctx boost.Context) error {
	model := models.BrandAdd{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, ErrorBrandBind(err))
	}

	if err := controller._add(&model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) UpdateREST(ctx boost.Context) error {
	brandID := ctx.Param("id").Int()
	if brandID == 0 {
		return controller.Error(ctx, ErrorBrandParamRequired("id"))
	}

	model := models.BrandUpdate{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, ErrorBrandBind(err))
	}

	if err := controller._update(brandID, &model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) UploadREST(ctx boost.Context) error {
	brandID := ctx.Param("id").Int()
	if brandID == 0 {
		return controller.Error(ctx, ErrorBrandParamRequired("id"))
	}

	model := models.BrandUpload{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, ErrorBrandBind(err))
	}

	imagePath, err := controller._upload(brandID, &model)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, imagePath)
}

func (controller *Controller) DeleteREST(ctx boost.Context) error {
	brandID := ctx.Param("id").Int()
	if brandID == 0 {
		return controller.Error(ctx, ErrorBrandParamRequired("id"))
	}

	if err := controller._delete(brandID); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}
