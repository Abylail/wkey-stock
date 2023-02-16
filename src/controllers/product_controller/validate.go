package product_controller

import "wkey-stock/src/data/models"

func (controller *Controller) validateProductUpdate(model *models.ProductUpdate) error {
	if err := controller.RequiredField(model.DescriptionRU, "description_ru"); err != nil {
		return err
	}

	if err := controller.RequiredField(model.DescriptionKZ, "description_kz"); err != nil {
		return err
	}

	return nil
}

func (controller *Controller) validateProductUpload(_ *models.ProductUpload) error {
	//

	return nil
}

func (controller *Controller) validateBrandUpdate(model *models.BrandUpdate) error {
	if err := controller.RequiredField(model.Title, "title"); err != nil {
		return err
	}

	return nil
}

func (controller *Controller) validateBrandUpload(model *models.BrandUpload) error {
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

func (controller *Controller) validateBrandAdd(model *models.BrandAdd) error {
	if err := controller.RequiredField(model.Title, "title"); err != nil {
		return err
	}

	if err := controller.RequiredField(model.Image, "image"); err != nil {
		return err
	}

	return nil
}
