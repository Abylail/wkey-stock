package product_controller

import (
	"wkey-stock/src/controllers/controller"
	"wkey-stock/src/repositories"
	"wkey-stock/src/repositories/product_repository"
)

type Controller struct {
	controller.Base
	productRepo *product_repository.Repository
}

func Create(apiRepositories *repositories.ApiRepositories) *Controller {
	return &Controller{
		productRepo: apiRepositories.Product,
	}
}
