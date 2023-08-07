package promotion_controller

import "github.com/lowl11/boost"

func (controller Controller) Get(ctx boost.Context) error {
	return controller.Ok(ctx)
}
