package promotion_controller

import (
	"github.com/lowl11/boost/pkg/base/controller"
	"wkey-stock/src/gateways"
	"wkey-stock/src/gateways/promotions_gateway"
)

type Controller struct {
	controller.Base
	promotions *promotions_gateway.Gateway
}

func New(gateways *gateways.Gateways) *Controller {
	return &Controller{
		promotions: gateways.Promotions,
	}
}
