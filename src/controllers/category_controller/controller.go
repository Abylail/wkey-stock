package category_controller

import (
	"wkey-stock/src/controllers/controller"
	"wkey-stock/src/repositories"
	"wkey-stock/src/repositories/category_repository"
)

type Controller struct {
	controller.Base
	categoryRepo *category_repository.Repository
}

func Create(apiRepositories *repositories.ApiRepositories) *Controller {
	return &Controller{
		categoryRepo: apiRepositories.Category,
	}
}
