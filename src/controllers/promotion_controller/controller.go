package promotion_controller

import (
	"wkey-stock/src/controllers/controller"
	"wkey-stock/src/events"
	"wkey-stock/src/events/image_event"
	"wkey-stock/src/repositories"
	"wkey-stock/src/repositories/product_repository"
	"wkey-stock/src/repositories/promotion_repository"
)

type Controller struct {
	controller.Base
	promotionRepo *promotion_repository.Repository
	productRepo   *product_repository.Repository
	image         *image_event.Event
}

func Create(apiRepositories *repositories.ApiRepositories, apiEvents *events.ApiEvents) *Controller {
	return &Controller{
		promotionRepo: apiRepositories.Promotion,
		productRepo:   apiRepositories.Product,
		image:         apiEvents.Image,
	}
}
