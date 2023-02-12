package product_controller

import "wkey-stock/src/data/models"

func (controller *Controller) filterBrandAdd(model *models.BrandAdd) {
	model.Title = controller.FilterStringSimple(model.Title)
}
