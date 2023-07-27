package promotion_controller

import (
	"github.com/lowl11/boost"
	"wkey-stock/src/data/errors"
	"wkey-stock/src/data/models"
)

// GetListAdmin Список всех акций
func (controller *Controller) GetListAdmin(ctx boost.Context) error {
	list, err := controller._getListAdmin()
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, list)
}

// GetSingleAdmin Получить промоакцию по id
func (controller *Controller) GetSingleAdmin(ctx boost.Context) error {
	id := ctx.Param("id").Int()

	promotion, err := controller._getSingleAdmin(id)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, promotion)
}

// GetSingleCodeAdmin Получить промоакцию по code
func (controller *Controller) GetSingleCodeAdmin(ctx boost.Context) error {
	code := ctx.Param("code").String()

	promotion, err := controller._getSingleCodeAdmin(code)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, promotion)
}

// CreateAdmin Создать промоакцию
func (controller *Controller) CreateAdmin(ctx boost.Context) error {
	model := models.PromotionAdminCreate{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, errors.PromotionCreateBind.With(err))
	}

	code, err := controller._createAdmin(&model)

	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, code)
}

// UpdateAdmin Обновить промоакцию
func (controller *Controller) UpdateAdmin(ctx boost.Context) error {
	model := models.PromotionAdminUpdate{}

	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, errors.PromotionCreateBind.With(err))
	}

	if err := controller._updateAdmin(&model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

// UploadAdmin Загрузить фото в промоакцию
func (controller *Controller) UploadAdmin(ctx boost.Context) error {
	model := models.PromotionAdminUpload{}

	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, errors.PromotionCreateBind.With(err))
	}

	if err := controller._uploadAdmin(&model); err != nil {
		return controller.Error(ctx, errors.PromotionUpload.With(err))
	}

	return controller.Ok(ctx, "OK")
}

// DeleteAdmin Загрузить фото в промоакцию
func (controller *Controller) DeleteAdmin(ctx boost.Context) error {
	code := ctx.Param("code").String()

	if err := controller._deleteAdmin(&code); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

// GetListClient Список всех акций
func (controller *Controller) GetListClient(ctx boost.Context) error {
	list, err := controller._getListClient()
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, list)
}

// GetSingleClient Получить промоакцию по code
func (controller *Controller) GetSingleClient(ctx boost.Context) error {
	code := ctx.Param("code").String()

	promotion, err := controller._getSingleClient(code)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, promotion)
}
