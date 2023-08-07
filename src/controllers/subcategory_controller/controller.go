package subcategory_controller

import (
	"github.com/lowl11/boost/pkg/base/controller"
	"wkey-stock/src/gateways"
	"wkey-stock/src/gateways/subcategory_gateway"
)

type Controller struct {
	controller.Base
	subCategories *subcategory_gateway.Gateway
}

func New(gateways *gateways.Gateways) *Controller {
	return &Controller{
		subCategories: gateways.SubCategories,
	}
}
