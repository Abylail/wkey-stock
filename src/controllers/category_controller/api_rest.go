package category_controller

import (
	"github.com/labstack/echo/v4"
	"strconv"
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

func (controller *Controller) GetClientSubREST(ctx echo.Context) error {
	list, err := controller._getClientSub()
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

func (controller *Controller) GetAdminSingleREST(ctx echo.Context) error {
	code := ctx.Param("code")

	category, err := controller._getAdminSingle(code)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, category)
}

func (controller *Controller) GetAdminSubREST(ctx echo.Context) error {
	parentCode := ctx.Param("parent_code")
	searchQuery := ctx.QueryParam("query")

	list, err := controller._getAdminSub(parentCode, searchQuery)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, list)
}

func (controller *Controller) GetAdminSingleSubREST(ctx echo.Context) error {
	parentCode := ctx.Param("parent_code")
	code := ctx.Param("code")

	list, err := controller._getAdminSubSingle(parentCode, code)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, list)
}

func (controller *Controller) AddREST(ctx echo.Context) error {
	model := models.CategoryAdd{}
	if err := ctx.Bind(&model); err != nil {
		return controller.Error(ctx, errors.CategoryAddBind.With(err))
	}

	if err := controller.validateCategoryAdd(&model); err != nil {
		return controller.Error(ctx, errors.CategoryAddValidate.With(err))
	}

	categoryCode, err := controller._create(&model)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, categoryCode)
}

func (controller *Controller) AddSubREST(ctx echo.Context) error {
	parentCode := ctx.Param("parent_code")

	model := models.SubCategoryAdd{}
	if err := ctx.Bind(&model); err != nil {
		return controller.Error(ctx, errors.CategoryAddBind.With(err))
	}

	if err := controller.validateCategoryAddSub(&model); err != nil {
		return controller.Error(ctx, errors.CategoryAddValidate.With(err))
	}

	subCategoryCode, err := controller._createSub(parentCode, &model)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, subCategoryCode)
}

func (controller *Controller) UpdateREST(ctx echo.Context) error {
	code := ctx.Param("code")
	if code == "" {
		return controller.Error(ctx, errors.CategoryUpdateParam)
	}

	model := models.CategoryUpdate{}
	if err := ctx.Bind(&model); err != nil {
		return controller.Error(ctx, errors.CategoryUpdateBind.With(err))
	}

	if err := controller.validateCategoryUpdate(&model); err != nil {
		return controller.Error(ctx, errors.CategoryUpdateValidate.With(err))
	}

	if err := controller._update(code, &model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) UpdateSubREST(ctx echo.Context) error {
	parentCode := ctx.Param("parent_code")
	code := ctx.Param("code")
	if code == "" {
		return controller.Error(ctx, errors.CategoryUpdateParam)
	}

	model := models.SubCategoryUpdate{}
	if err := ctx.Bind(&model); err != nil {
		return controller.Error(ctx, errors.CategoryUpdateBind.With(err))
	}

	if err := controller.validateCategoryUpdateSub(&model); err != nil {
		return controller.Error(ctx, errors.CategoryUpdateValidate.With(err))
	}

	if err := controller._updateSub(parentCode, code, &model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) UploadREST(ctx echo.Context) error {
	code := ctx.Param("code")
	if code == "" {
		return controller.Error(ctx, errors.CategoryUploadParam)
	}

	model := models.CategoryUpload{}
	if err := ctx.Bind(&model); err != nil {
		return controller.Error(ctx, errors.CategoryUploadBind.With(err))
	}

	if err := controller.validateCategoryUpload(&model); err != nil {
		return controller.Error(ctx, errors.CategoryUploadValidate.With(err))
	}

	imagePath, err := controller._upload(code, &model)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, imagePath)
}

func (controller *Controller) UploadSubREST(ctx echo.Context) error {
	parentCode := ctx.Param("parent_code")

	code := ctx.Param("code")
	if code == "" {
		return controller.Error(ctx, errors.CategoryUploadParam)
	}

	model := models.SubCategoryUpload{}
	if err := ctx.Bind(&model); err != nil {
		return controller.Error(ctx, errors.CategoryUploadBind.With(err))
	}

	if err := controller.validateCategoryUploadSub(&model); err != nil {
		return controller.Error(ctx, errors.CategoryUploadValidate.With(err))
	}

	imagePath, err := controller._uploadSub(parentCode, code, &model)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, imagePath)
}

func (controller *Controller) DeleteREST(ctx echo.Context) error {
	categoryCode := ctx.Param("code")
	if categoryCode == "" {
		return controller.Error(ctx, errors.CategoryDeleteParam)
	}

	if err := controller._delete(categoryCode); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) DeleteSubREST(ctx echo.Context) error {
	parentCode := ctx.Param("parent_code")
	categoryCode := ctx.Param("code")
	if categoryCode == "" {
		return controller.Error(ctx, errors.CategoryDeleteParam)
	}

	if err := controller._deleteSub(parentCode, categoryCode); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) ActivateREST(ctx echo.Context) error {
	code := ctx.Param("code")

	if err := controller._activate(code); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) DeactivateREST(ctx echo.Context) error {
	code := ctx.Param("code")

	if err := controller._deactivate(code); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) ActivateSubREST(ctx echo.Context) error {
	parentCode := ctx.Param("parent_code")
	code := ctx.Param("code")

	if err := controller._activateSub(parentCode, code); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) DeactivateSubREST(ctx echo.Context) error {
	parentCode := ctx.Param("parent_code")
	code := ctx.Param("code")

	if err := controller._deactivateSub(parentCode, code); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) BindProductListREST(ctx echo.Context) error {
	parentCode := ctx.Param("parent_code")
	code := ctx.Param("code")

	model := models.SubCategoryBindProductList{}
	if err := ctx.Bind(&model); err != nil {
		return controller.Error(ctx, errors.SubCategoryBindProductListBind.With(err))
	}

	if err := controller.validateSubCategoryProductList(&model); err != nil {
		return controller.Error(ctx, errors.SubCategoryBindProductListValidate.With(err))
	}

	if err := controller._bindProductList(parentCode, code, &model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) UnbindProductItemREST(ctx echo.Context) error {
	parentCode := ctx.Param("parent_code")
	code := ctx.Param("code")
	productID, _ := strconv.Atoi(ctx.Param("product_id"))
	if productID == 0 {
		return errors.UnbindProductFromSubCategory
	}

	if err := controller._unbindProductItem(parentCode, code, productID); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}
