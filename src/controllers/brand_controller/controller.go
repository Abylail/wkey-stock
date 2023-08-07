package brand_controller

import (
	"github.com/lowl11/boost/pkg/base/controller"
	"wkey-stock/src/gateways"
	"wkey-stock/src/gateways/brands_gateway"
)

type Controller struct {
	controller.Base
	brands *brands_gateway.Gateway
}

func New(gateways *gateways.Gateways) *Controller {
	return &Controller{
		brands: gateways.Brands,
	}
}
