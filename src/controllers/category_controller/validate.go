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

func (controller *Controller) validateCategoryUpload(model *models.CategoryUpload) error {
	if err := controller.RequiredField(model.Image, "image"); err != nil {
		return err
	}

	if err := controller.RequiredField(model.Image.Name, "image_name"); err != nil {
		return err
	}

	if err := controller.RequiredField(model.Image.Buffer, "image_buffer"); err != nil {
		return err
	}

	return nil
}
