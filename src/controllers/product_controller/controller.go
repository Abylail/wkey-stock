package product_controller

import (
	"github.com/lowl11/boost/pkg/base/controller"
	"wkey-stock/src/events"
	"wkey-stock/src/events/image_event"
	"wkey-stock/src/repositories"
	"wkey-stock/src/repositories/brand_repository"
	"wkey-stock/src/repositories/product_repository"
	"wkey-stock/src/repositories/sub_category_repository"
)

type Controller struct {
	controller.Base
	productRepo     *product_repository.Repository
	subCategoryRepo *sub_category_repository.Repository
	brandRepo       *brand_repository.Repository
	image           *image_event.Event
}

func New(apiRepositories *repositories.ApiRepositories, apiEvents *events.ApiEvents) *Controller {
	return &Controller{
		productRepo:     apiRepositories.Product,
		subCategoryRepo: apiRepositories.SubCategory,
		brandRepo:       apiRepositories.Brand,
		image:           apiEvents.Image,
	}
}
