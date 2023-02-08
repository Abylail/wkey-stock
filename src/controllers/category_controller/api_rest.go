package category_controller

import (
	"github.com/labstack/echo/v4"
)

func (controller *Controller) GetClientREST(ctx echo.Context) error {
	list, err := controller._getClient()
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, list)
}

func (controller *Controller) GetAdminREST(ctx echo.Context) error {
	searchQuery := ctx.QueryParam("query")

	list, err := controller._getAdmin(searchQuery)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, list)
}
