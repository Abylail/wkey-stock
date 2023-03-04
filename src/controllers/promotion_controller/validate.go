package promotion_controller

import "wkey-stock/src/data/models"

func (controller *Controller) validatePromotionCreate(model *models.PromotionAdminCreate) error {
	if err := controller.RequiredField(model.TitleRU, "title_ru"); err != nil {
		return err
	}
	if err := controller.RequiredField(model.TitleRU, "title_kz"); err != nil {
		return err
	}

	return nil
}
