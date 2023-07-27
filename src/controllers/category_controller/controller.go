package category_controller

import (
	"github.com/lowl11/boost/pkg/base/controller"
	"wkey-stock/src/events"
	"wkey-stock/src/events/image_event"
	"wkey-stock/src/repositories"
	"wkey-stock/src/repositories/category_repository"
	"wkey-stock/src/repositories/product_repository"
	"wkey-stock/src/repositories/sub_category_repository"
)

type Controller struct {
	controller.Base
	categoryRepo    *category_repository.Repository
	subCategoryRepo *sub_category_repository.Repository
	productRepo     *product_repository.Repository
	image           *image_event.Event
}

func Create(apiRepositories *repositories.ApiRepositories, apiEvents *events.ApiEvents) *Controller {
	return &Controller{
		categoryRepo:    apiRepositories.Category,
		subCategoryRepo: apiRepositories.SubCategory,
		productRepo:     apiRepositories.Product,
		image:           apiEvents.Image,
	}
}
