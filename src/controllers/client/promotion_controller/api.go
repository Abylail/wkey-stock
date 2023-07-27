package promotion_controller

import (
	"wkey-stock/src/adaptors/promotion_adaptor"
	"wkey-stock/src/data/dtos"
)

// _getListClient список промо акций (в админке)
func (controller Controller) _getList() ([]dtos.Promotion, error) {
	list, err := controller.promotionRepo.GetAll()
	if err != nil {
		return nil, ErrorPromotionGetList(err)
	}

	return promotion_adaptor.EntityToDTO(list), nil
}

// _getSingleCodeAdmin промо акция по code
func (controller Controller) _getByCode(code string) (*dtos.Promotion, error) {
	promotion, err := controller.promotionRepo.GetByCode(code)
	if err != nil {
		return nil, ErrorPromotionGetByCode(err)
	}

	// Если не нашелся
	if promotion == nil {
		return nil, ErrorPromotionNotFoundByCode(code)
	}

	return dtos.NewPromotion(promotion), nil
}
