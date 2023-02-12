package product_controller

import "wkey-stock/src/data/models"

func (controller *Controller) validateBrandUpdate(model *models.BrandUpdate) error {
	if err := controller.RequiredField(model.Title, "title"); err != nil {
		return err
	}

	if err := controller.RequiredField(model.Image, "image"); err != nil {
		return err
	}

	return nil
}

func (controller *Controller) validateBrandAdd(model *models.BrandAdd) error {
	if err := controller.RequiredField(model.Title, "title"); err != nil {
		return err
	}

	if err := controller.RequiredField(model.Image, "image"); err != nil {
		return err
	}

	return nil
}
