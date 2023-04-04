package promotion_controller

import (
	"github.com/labstack/echo/v4"
	"strconv"
	"wkey-stock/src/data/errors"
	"wkey-stock/src/data/models"
)

// GetListAdmin Список всех акций
func (controller *Controller) GetListAdmin(ctx echo.Context) error {
	list, err := controller._getListAdmin()
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, list)
}

// GetSingleAdmin Получить промоакцию по id
func (controller *Controller) GetSingleAdmin(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	promotion, err := controller._getSingleAdmin(id)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, promotion)
}

// GetSingleCodeAdmin Получить промоакцию по code
func (controller *Controller) GetSingleCodeAdmin(ctx echo.Context) error {
	code := ctx.Param("code")

	promotion, err := controller._getSingleCodeAdmin(code)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, promotion)
}

// CreateAdmin Создать промоакцию
func (controller *Controller) CreateAdmin(ctx echo.Context) error {
	model := models.PromotionAdminCreate{}
	if err := ctx.Bind(&model); err != nil {
		return controller.Error(ctx, errors.PromotionCreateBind.With(err))
	}

	if err := controller.validatePromotionCreate(&model); err != nil {
		return controller.Error(ctx, errors.BrandAddValidate.With(err))
	}

	code, err := controller._createAdmin(&model)

	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, code)
}

// UpdateAdmin Обновить промоакцию
func (controller *Controller) UpdateAdmin(ctx echo.Context) error {
	model := models.PromotionAdminUpdate{}

	if err := ctx.Bind(&model); err != nil {
		return controller.Error(ctx, errors.PromotionCreateBind.With(err))
	}

	if err := controller.validatePromotionUpdate(&model); err != nil {
		return controller.Error(ctx, errors.BrandAddValidate.With(err))
	}

	if err := controller._updateAdmin(&model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

// UploadAdmin Загрузить фото в промоакцию
func (controller *Controller) UploadAdmin(ctx echo.Context) error {
	model := models.PromotionAdminUpload{}

	if err := ctx.Bind(&model); err != nil {
		return controller.Error(ctx, errors.PromotionCreateBind.With(err))
	}

	if err := controller.validatePromotionUpload(&model); err != nil {
		return controller.Error(ctx, errors.PromotionCreateValidate.With(err))
	}

	if err := controller._uploadAdmin(&model); err != nil {
		return controller.Error(ctx, errors.PromotionUpload.With(err))
	}

	return controller.Ok(ctx, "OK")
}

// DeleteAdmin Загрузить фото в промоакцию
func (controller *Controller) DeleteAdmin(ctx echo.Context) error {
	code := ctx.Param("code")

	if err := controller._deleteAdmin(&code); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

// GetListClient Список всех акций
func (controller *Controller) GetListClient(ctx echo.Context) error {
	list, err := controller._getListClient()
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, list)
}
