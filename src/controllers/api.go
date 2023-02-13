package controllers

import (
	"wkey-stock/src/controllers/category_controller"
	"wkey-stock/src/controllers/product_controller"
	"wkey-stock/src/controllers/static_controller"
	"wkey-stock/src/events"
	"wkey-stock/src/repositories"
)

type ApiControllers struct {
	Static *static_controller.Controller

	Product  *product_controller.Controller
	Category *category_controller.Controller
}

func Get(apiEvents *events.ApiEvents, apiRepositories *repositories.ApiRepositories) *ApiControllers {
	return &ApiControllers{
		Static: static_controller.Create(),

		Product:  product_controller.Create(apiRepositories, apiEvents),
		Category: category_controller.Create(apiRepositories, apiEvents),
	}
}
