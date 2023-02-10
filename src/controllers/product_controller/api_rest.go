package product_controller

import (
	"github.com/labstack/echo/v4"
	"strconv"
	"wkey-stock/src/data/errors"
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

func (controller *Controller) GetAdminSingleREST(ctx echo.Context) error {
	productID, _ := strconv.Atoi(ctx.Param("id"))
	if productID == 0 {
		return controller.NotFound(ctx, errors.AdminProductNotFound)
	}

	product, err := controller._getAdminSingle(productID)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, product)
}

func (controller *Controller) GetClientREST(ctx echo.Context) error {
	return controller.Ok(ctx, "OK")
}
