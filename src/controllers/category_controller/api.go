package category_controller

import (
	"github.com/lowl11/lazy-collection/type_list"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/errors"
	"wkey-stock/src/data/models"
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
		list, err = controller.categoryRepo.ByQuery(searchQuery)
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
				ID:         item.ID,
				Key:        item.Key,
				Title:      item.Title,
				ParentID:   item.ParentID,
				Position:   item.Position,
				ItemsCount: item.ItemsCount,
			}
		}).
		Slice()

	return categories, nil
}
