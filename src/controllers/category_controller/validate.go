package category_controller

import "wkey-stock/src/data/models"

func (controller *Controller) validateCategoryAdd(model *models.CategoryAdd) error {
	if err := controller.RequiredField(model.TitleRU, "title_ru"); err != nil {
		return err
	}

	if err := controller.RequiredField(model.TitleKZ, "title_kz"); err != nil {
		return err
	}

	return nil
}

func (controller *Controller) validateCategoryUpdate(model *models.CategoryUpdate) error {
	if err := controller.RequiredField(model.TitleRU, "title_ru"); err != nil {
		return err
	}

	if err := controller.RequiredField(model.TitleKZ, "title_kz"); err != nil {
		return err
	}

	return nil
}
