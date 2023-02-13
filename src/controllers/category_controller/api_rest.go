package category_controller

import (
	"github.com/labstack/echo/v4"
	"wkey-stock/src/data/errors"
	"wkey-stock/src/data/models"
)

func (controller *Controller) GetClientREST(ctx echo.Context) error {
	list, err := controller._getClient()
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, list)
}

func (controller *Controller) GetAdminREST(ctx echo.Context) error {
	searchQuery := ctx.QueryParam("query")

	list, err := controller._getAdmin(searchQuery)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, list)
}

func (controller *Controller) AddREST(ctx echo.Context) error {
	model := models.CategoryAdd{}
	if err := ctx.Bind(&model); err != nil {
		return errors.CategoryAddBind.With(err)
	}

	if err := controller.validateCategoryAdd(&model); err != nil {
		return errors.CategoryAddValidate.With(err)
	}

	if err := controller._create(&model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) UpdateREST(ctx echo.Context) error {
	code := ctx.Param("code")
	if code == "" {
		return errors.CategoryUpdateParam
	}

	model := models.CategoryUpdate{}
	if err := ctx.Bind(&model); err != nil {
		return errors.CategoryUpdateBind.With(err)
	}

	if err := controller.validateCategoryUpdate(&model); err != nil {
		return errors.CategoryUpdateValidate.With(err)
	}

	if err := controller._update(code, &model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) UploadREST(ctx echo.Context) error {
	code := ctx.Param("code")
	if code == "" {
		return errors.CategoryUploadParam
	}

	model := models.CategoryUpload{}
	if err := ctx.Bind(&model); err != nil {
		return errors.CategoryUploadBind.With(err)
	}

	if err := controller.validateCategoryUpload(&model); err != nil {
		return errors.CategoryUploadValidate.With(err)
	}

	imagePath, err := controller._upload(code, &model)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, imagePath)
}

func (controller *Controller) DeleteREST(ctx echo.Context) error {
	categoryCode := ctx.Param("code")
	if categoryCode == "" {
		return errors.CategoryDeleteParam
	}

	if err := controller._delete(categoryCode); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}
