package client_promotion_controller

import (
	"github.com/lowl11/boost"
	"wkey-stock/src/adaptors/promotion_adaptor"
)

// GetListREST Список всех акций
func (controller *Controller) GetListREST(ctx boost.Context) error {
	list, err := controller._getList()
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, promotion_adaptor.DtoToModel(list))
}

// GetByCodeREST Получить промо акцию по code
func (controller *Controller) GetByCodeREST(ctx boost.Context) error {
	code := ctx.Param("code").String()

	promotion, err := controller._getByCode(code)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, promotion.Model())
}
