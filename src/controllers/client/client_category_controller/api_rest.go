package client_category_controller

import "github.com/lowl11/boost"

func (controller Controller) GetREST(ctx boost.Context) error {
	searchQuery := ctx.QueryParam("query").String()

	list, err := controller._get(searchQuery)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, list)
}

func (controller Controller) GetSingleREST(ctx boost.Context) error {
	code := ctx.Param("code").String()

	category, err := controller._getByCode(code)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, category)
}

func (controller Controller) GetSubREST(ctx boost.Context) error {
	parentCode := ctx.Param("par_code").String()
	searchQuery := ctx.QueryParam("query").String()

	list, err := controller._getSubList(parentCode, searchQuery)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, list)
}

func (controller Controller) GetSubSingleREST(ctx boost.Context) error {
	parentCode := ctx.Param("par_code").String()
	code := ctx.Param("code").String()

	list, err := controller._getSubSingle(parentCode, code)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, list)
}
