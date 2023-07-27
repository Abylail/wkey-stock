package category_controller

import (
	"github.com/lowl11/boost"
	"wkey-stock/src/data/models"
)

func (controller *Controller) GetClientREST(ctx boost.Context) error {
	searchQuery := ctx.QueryParam("query").String()
	list, err := controller._getClient(searchQuery)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, list)
}

func (controller *Controller) GetClientSingleREST(ctx boost.Context) error {
	code := ctx.Param("code").String()

	category, err := controller._getClientSingle(code)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, category)
}

func (controller *Controller) GetClientSubREST(ctx boost.Context) error {
	parentCode := ctx.Param("par_code").String()
	searchQuery := ctx.QueryParam("query").String()

	list, err := controller._getClientSub(parentCode, searchQuery)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, list)
}

func (controller *Controller) GetClientSubSingleREST(ctx boost.Context) error {
	parentCode := ctx.Param("par_code").String()
	code := ctx.Param("code").String()

	list, err := controller._getClientSubSingle(parentCode, code)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, list)
}

func (controller *Controller) GetAdminREST(ctx boost.Context) error {
	searchQuery := ctx.QueryParam("query").String()

	list, err := controller._getAdmin(searchQuery)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, list)
}

func (controller *Controller) GetAdminSingleREST(ctx boost.Context) error {
	code := ctx.Param("code").String()

	category, err := controller._getAdminSingle(code)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, category)
}

func (controller *Controller) GetAdminSubREST(ctx boost.Context) error {
	parentCode := ctx.Param("parent_code").String()
	searchQuery := ctx.QueryParam("query").String()

	list, err := controller._getAdminSub(parentCode, searchQuery)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, list)
}

func (controller *Controller) GetAdminSingleSubREST(ctx boost.Context) error {
	parentCode := ctx.Param("parent_code").String()
	code := ctx.Param("code").String()

	list, err := controller._getAdminSubSingle(parentCode, code)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, list)
}

func (controller *Controller) AddREST(ctx boost.Context) error {
	model := models.CategoryAdd{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, ErrorCategoryBind(err))
	}

	categoryCode, err := controller._create(&model)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, categoryCode)
}

func (controller *Controller) AddSubREST(ctx boost.Context) error {
	parentCode := ctx.Param("parent_code").String()

	model := models.SubCategoryAdd{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, ErrorCategoryBind(err))
	}

	subCategoryCode, err := controller._createSub(parentCode, &model)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, subCategoryCode)
}

func (controller *Controller) UpdateREST(ctx boost.Context) error {
	code := ctx.Param("code").String()
	if code == "" {
		return controller.Error(ctx, ErrorCategoryParamRequired("code"))
	}

	model := models.CategoryUpdate{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, ErrorCategoryBind(err))
	}

	if err := controller._update(code, &model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) UpdateSubREST(ctx boost.Context) error {
	parentCode := ctx.Param("parent_code").String()
	code := ctx.Param("code").String()
	if code == "" {
		return controller.Error(ctx, ErrorCategoryParamRequired("parent_code"))
	}

	model := models.SubCategoryUpdate{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, ErrorCategoryBind(err))
	}

	if err := controller._updateSub(parentCode, code, &model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) UploadREST(ctx boost.Context) error {
	code := ctx.Param("code").String()
	if code == "" {
		return controller.Error(ctx, ErrorCategoryParamRequired("code"))
	}

	model := models.CategoryUpload{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, ErrorCategoryBind(err))
	}

	imagePath, err := controller._upload(code, &model)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, imagePath)
}

func (controller *Controller) UploadSubREST(ctx boost.Context) error {
	parentCode := ctx.Param("parent_code").String()

	code := ctx.Param("code").String()
	if code == "" {
		return controller.Error(ctx, ErrorCategoryParamRequired("code"))
	}

	model := models.SubCategoryUpload{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, ErrorCategoryBind(err))
	}

	imagePath, err := controller._uploadSub(parentCode, code, &model)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, imagePath)
}

func (controller *Controller) DeleteREST(ctx boost.Context) error {
	categoryCode := ctx.Param("code").String()
	if categoryCode == "" {
		return controller.Error(ctx, ErrorCategoryParamRequired("code"))
	}

	if err := controller._delete(categoryCode); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) DeleteSubREST(ctx boost.Context) error {
	parentCode := ctx.Param("parent_code").String()
	categoryCode := ctx.Param("code").String()
	if categoryCode == "" {
		return controller.Error(ctx, ErrorCategoryParamRequired("code"))
	}

	if err := controller._deleteSub(parentCode, categoryCode); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) ActivateREST(ctx boost.Context) error {
	code := ctx.Param("code").String()

	if err := controller._activate(code); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) DeactivateREST(ctx boost.Context) error {
	code := ctx.Param("code").String()

	if err := controller._deactivate(code); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) ActivateSubREST(ctx boost.Context) error {
	parentCode := ctx.Param("parent_code").String()
	code := ctx.Param("code").String()

	if err := controller._activateSub(parentCode, code); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) DeactivateSubREST(ctx boost.Context) error {
	parentCode := ctx.Param("parent_code").String()
	code := ctx.Param("code").String()

	if err := controller._deactivateSub(parentCode, code); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) BindProductListREST(ctx boost.Context) error {
	parentCode := ctx.Param("parent_code").String()
	code := ctx.Param("code").String()

	model := models.SubCategoryBindProductList{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, ErrorSubCategoryBind(err))
	}

	if err := controller._bindProductList(parentCode, code, &model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

func (controller *Controller) UnbindProductItemREST(ctx boost.Context) error {
	parentCode := ctx.Param("parent_code").String()
	code := ctx.Param("code").String()
	productID := ctx.Param("product_id").Int()
	if productID == 0 {
		return controller.Error(ctx, ErrorCategoryParamRequired("product_id"))
	}

	if err := controller._unbindProductItem(parentCode, code, productID); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}
