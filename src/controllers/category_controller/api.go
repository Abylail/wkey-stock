package category_controller

import (
	"github.com/lowl11/lazy-collection/type_list"
	"github.com/lowl11/lazylog/layers"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/errors"
	"wkey-stock/src/data/models"
	"wkey-stock/src/definition"
)

func (controller *Controller) _getClient() ([]models.CategoryClientGet, *models.Error) {
	return nil, nil
}

func (controller *Controller) _getAdmin(searchQuery string) ([]models.CategoryAdminGet, *models.Error) {
	var list []entities.CategoryGet
	var err error

	if searchQuery == "" {
		list, err = controller.categoryRepo.GetAll()
	} else {
		list, err = controller.categoryRepo.GetByQuery(searchQuery)
	}
	if err != nil {
		return nil, errors.CategoryGetList.With(err)
	}

	if len(list) == 0 {
		return []models.CategoryAdminGet{}, nil
	}

	categories := type_list.NewWithList[entities.CategoryGet, models.CategoryAdminGet](list...).
		Select(func(item entities.CategoryGet) models.CategoryAdminGet {
			return models.CategoryAdminGet{
				ID:      item.ID,
				Code:    item.Code,
				TitleRU: item.TitleRU,
				TitleKZ: item.TitleKZ,
				Icon:    item.Icon,
			}
		}).
		Slice()

	return categories, nil
}

func (controller *Controller) _create(model *models.CategoryAdd) *models.Error {
	logger := definition.Logger

	if err := controller.categoryRepo.Create(model); err != nil {
		logger.Error(err, "Update category error", layers.Database)
		return errors.CategoryUpdate.With(err)
	}

	return nil
}

func (controller *Controller) _update(code string, model *models.CategoryUpdate) *models.Error {
	logger := definition.Logger

	if err := controller.categoryRepo.Update(code, model); err != nil {
		logger.Error(err, "Update category error", layers.Database)
		return errors.CategoryUpdate.With(err)
	}

	return nil
}

func (controller *Controller) _delete(categoryCode string) *models.Error {
	logger := definition.Logger

	if err := controller.categoryRepo.Delete(categoryCode); err != nil {
		logger.Error(err, "Delete category error", layers.Database)
		return errors.CategoryDelete.With(err)
	}

	return nil
}
