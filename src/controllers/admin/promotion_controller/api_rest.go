package promotion_controller

import (
	"github.com/lowl11/boost"
	"wkey-stock/src/adaptors/promotion_adaptor"
	"wkey-stock/src/data/models"
)

// GetListREST Список всех акций
func (controller Controller) GetListREST(ctx boost.Context) error {
	list, err := controller._getListAdmin()
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, promotion_adaptor.DtoToModel(list))
}

// GetSingleREST Получить промо акцию по id
func (controller Controller) GetSingleREST(ctx boost.Context) error {
	id := ctx.Param("id").Int()

	promotion, err := controller._getSingleAdmin(id)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, promotion.Model())
}

// GetSingleByCodeREST Получить промо акцию по code
func (controller Controller) GetSingleByCodeREST(ctx boost.Context) error {
	code := ctx.Param("code").String()

	promotion, err := controller._getSingleCodeAdmin(code)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, promotion.Model())
}

// CreateREST Создать промо акцию
func (controller Controller) CreateREST(ctx boost.Context) error {
	model := models.PromotionAdminCreate{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, err)
	}

	code, err := controller._createAdmin(&model)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, code)
}

// UpdateREST Обновить промо акцию
func (controller Controller) UpdateREST(ctx boost.Context) error {
	model := models.PromotionAdminUpdate{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, err)
	}

	if err := controller._updateAdmin(&model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

// UploadREST Загрузить фото в промо акцию
func (controller Controller) UploadREST(ctx boost.Context) error {
	model := models.PromotionAdminUpload{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, err)
	}

	if err := controller._uploadAdmin(&model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}

// DeleteREST Загрузить фото в промо акцию
func (controller Controller) DeleteREST(ctx boost.Context) error {
	code := ctx.Param("code").String()

	if err := controller._deleteAdmin(&code); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, "OK")
}
