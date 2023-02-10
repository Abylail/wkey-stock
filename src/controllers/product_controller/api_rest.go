package product_controller

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

func (controller *Controller) GetAdminREST(ctx echo.Context) error {
	searchQuery := ctx.QueryParam("query")
	categoryKey := ctx.QueryParam("category")
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	if page == 0 {
		page = 1
	}

	from := (page - 1) * 20
	to := from + 20

	products, err := controller._getAdmin(from, to, searchQuery, categoryKey)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, products)
}

func (controller *Controller) GetClientREST(ctx echo.Context) error {
	return controller.Ok(ctx, "OK")
}
