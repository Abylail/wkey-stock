package category_controller

import (
	"github.com/lowl11/boost/pkg/base/controller"
	"wkey-stock/src/gateways"
	"wkey-stock/src/gateways/category_gateway"
)

type Controller struct {
	controller.Base
	categories *category_gateway.Gateway
}

func New(gateways *gateways.Gateways) *Controller {
	return &Controller{
		categories: gateways.Categories,
	}
}
