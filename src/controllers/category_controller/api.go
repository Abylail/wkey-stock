package category_controller

import (
	"github.com/lowl11/boost"
	"wkey-stock/src/adaptors/category_adaptor"
	"wkey-stock/src/data/models"
)

func (controller Controller) Get(ctx boost.Context) error {
	categories, err := controller.categories.GetAll()
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, category_adaptor.DtoToModel(categories))
}

func (controller Controller) GetByID(ctx boost.Context) error {
	categoryID := ctx.Param("category-id").String()

	category, err := controller.categories.GetByID(categoryID)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx, category.Model())
}

func (controller Controller) Add(ctx boost.Context) error {
	model := models.CategoryProsklad{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, err)
	}

	newID, err := controller.categories.Add(&model)
	if err != nil {
		return controller.Error(ctx, err)
	}

	return controller.CreatedID(ctx, newID)
}

func (controller Controller) UpdateProsklad(ctx boost.Context) error {
	proskladID := ctx.Param("prosklad-id").Int()
	if proskladID == 0 {
		return controller.Error(ctx, ErrorProskladIDRequired())
	}

	model := models.CategoryProsklad{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, err)
	}

	if err := controller.categories.UpdateProsklad(proskladID, &model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx)
}

func (controller Controller) UpdateCount(ctx boost.Context) error {
	categoryID := ctx.Param("category-id").String()
	if categoryID == "" {
		return controller.Error(ctx, ErrorCategoryIDRequired())
	}

	model := models.CategoryUpdateCount{}
	if err := ctx.Parse(&model); err != nil {
		return controller.Error(ctx, err)
	}

	if err := controller.categories.UpdateCount(categoryID, &model); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx)
}

func (controller Controller) Delete(ctx boost.Context) error {
	categoryID := ctx.Param("category-id").String()

	if categoryID == "" {
		return controller.Error(ctx, ErrorCategoryIDRequired())
	}

	if err := controller.categories.RemoveByID(categoryID); err != nil {
		return controller.Error(ctx, err)
	}

	return controller.Ok(ctx)
}
