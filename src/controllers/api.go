package controllers

import (
	"wkey-stock/src/controllers/category_controller"
	"wkey-stock/src/controllers/product_controller"
	"wkey-stock/src/controllers/promotion_controller"
	"wkey-stock/src/events"
	"wkey-stock/src/repositories"
)

type ApiControllers struct {
	Product   *product_controller.Controller
	Category  *category_controller.Controller
	Promotion *promotion_controller.Controller
}

func Get(apiEvents *events.ApiEvents, apiRepositories *repositories.ApiRepositories) *ApiControllers {
	return &ApiControllers{
		Product:   product_controller.New(apiRepositories, apiEvents),
		Category:  category_controller.New(apiRepositories, apiEvents),
		Promotion: promotion_controller.New(apiRepositories, apiEvents),
	}
}
