package product_controller

import "github.com/lowl11/boost"

func (controller *Controller) GetREST(ctx boost.Context) error {
	searchQuery := ctx.QueryParam("query").String()
	page := ctx.QueryParam("page").Int()
	if page == 0 {
		page = 1
	}

	pageSize := 20
	from := (page - 1) * pageSize

	products, err := controller._get(from, pageSize, searchQuery)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, products)
}
