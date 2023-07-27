package brand_controller

import (
	"github.com/lowl11/boost/pkg/base/controller"
	"wkey-stock/src/events"
	"wkey-stock/src/events/image_event"
	"wkey-stock/src/repositories"
	"wkey-stock/src/repositories/brand_repository"
)

type Controller struct {
	controller.Base

	brandRepo *brand_repository.Repository
	image     *image_event.Event
}

func New(apiRepositories *repositories.ApiRepositories, apiEvents *events.ApiEvents) *Controller {
	return &Controller{
		brandRepo: apiRepositories.Brand,
		image:     apiEvents.Image,
	}
}
