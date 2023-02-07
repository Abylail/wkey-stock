package category_controller

import (
	"github.com/lowl11/lazy-collection/type_list"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

func (controller *Controller) _getClient() ([]models.CategoryClientGet, *models.Error) {
	return nil, nil
}

func (controller *Controller) _getAdmin() ([]models.CategoryAdminGet, *models.Error) {
	list, err := controller.categoryRepo.GetAll()
	if err != nil {
		return nil, nil
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
