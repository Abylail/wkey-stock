package product_controller

import (
	"strings"
	"wkey-stock/src/data/models"
)

func (controller *Controller) filterBrandAdd(model *models.BrandAdd) {
	model.Title = strings.TrimSpace(model.Title)
}
