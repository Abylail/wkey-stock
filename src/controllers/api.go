package controllers

import (
	"wkey-stock/src/controllers/admin/admin_category_controller"
	"wkey-stock/src/controllers/admin/admin_product_controller"
	"wkey-stock/src/controllers/admin/admin_promotion_controller"
	"wkey-stock/src/controllers/client/client_category_controller"
	"wkey-stock/src/controllers/client/client_product_controller"
	"wkey-stock/src/controllers/client/client_promotion_controller"
	"wkey-stock/src/events"
	"wkey-stock/src/repositories"
)

type ApiControllers struct {
	AdminProduct  *admin_product_controller.Controller
	ClientProduct *client_product_controller.Controller

	AdminCategory  *admin_category_controller.Controller
	ClientCategory *client_category_controller.Controller

	AdminPromotion  *admin_promotion_controller.Controller
	ClientPromotion *client_promotion_controller.Controller
}

func Get(apiEvents *events.ApiEvents, apiRepositories *repositories.ApiRepositories) *ApiControllers {
	return &ApiControllers{
		AdminProduct:  admin_product_controller.New(apiRepositories, apiEvents),
		ClientProduct: client_product_controller.New(apiRepositories, apiEvents),

		AdminCategory:  admin_category_controller.New(apiRepositories, apiEvents),
		ClientCategory: client_category_controller.New(apiRepositories, apiEvents),

		AdminPromotion:  admin_promotion_controller.New(apiRepositories, apiEvents),
		ClientPromotion: client_promotion_controller.New(apiRepositories, apiEvents),
	}
}
