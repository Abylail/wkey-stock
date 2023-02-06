package category_controller

import "github.com/labstack/echo/v4"

func (controller *Controller) Get(ctx echo.Context) error {
	return controller.Ok(ctx, "OK")
}
