package static_controller

import "wkey-stock/src/controllers/controller"

type Controller struct {
	controller.Base
}

func Create() *Controller {
	return &Controller{}
}
