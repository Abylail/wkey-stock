package product_controller

import (
	"github.com/lowl11/boost/pkg/base/controller"
	"wkey-stock/src/gateways"
	"wkey-stock/src/gateways/product_gateway"
)

type Controller struct {
	controller.Base

	products *product_gateway.Gateway
}

func New(gateways *gateways.Gateways) *Controller {
	return &Controller{
		products: gateways.Products,
	}
}
