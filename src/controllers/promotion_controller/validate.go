package promotion_controller

import (
	"errors"
	"wkey-stock/src/data/models"
)

func (controller *Controller) validatePromotionCreate(model *models.PromotionAdminCreate) error {
	if err := controller.RequiredField(model.TitleRU, "title_ru"); err != nil {
		return err
	}
	if err := controller.RequiredField(model.TitleRU, "title_kz"); err != nil {
		return err
	}

	return nil
}

func (controller *Controller) validatePromotionUpdate(model *models.PromotionAdminUpdate) error {
	if err := controller.RequiredField(model.Code, "code"); err != nil {
		return err
	}
	if err := controller.RequiredField(model.TitleRU, "title_ru"); err != nil {
		return err
	}
	if err := controller.RequiredField(model.TitleRU, "title_kz"); err != nil {
		return err
	}

	return nil
}

func (controller *Controller) validatePromotionUpload(model *models.PromotionAdminUpload) error {
	if err := controller.RequiredField(model.Code, "code"); err != nil {
		return err
	}

	if model.Lang != "ru" && model.Lang != "kz" {
		return errors.New("lang can be only ru or kz")
	}

	return nil
}
