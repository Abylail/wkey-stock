package category_controller

import (
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

// _get список
func (controller Controller) _get(searchQuery string) ([]models.CategoryClientGet, error) {
	var list []entities.Category
	var err error

	if len(searchQuery) == 0 {
		list, err = controller.categoryRepo.GetAll()
	} else {
		list, err = controller.categoryRepo.GetByQuery(searchQuery)
	}

	if err != nil {
		return nil, ErrorCategoryGetList(err)
	}

	categories := make([]models.CategoryClientGet, 0, len(list))
	for _, category := range list {
		categories = append(categories, models.CategoryClientGet{
			Code:    category.Code,
			TitleRU: category.TitleRU,
			TitleKZ: category.TitleKZ,
			Image:   category.Icon,
		})
	}

	return categories, nil
}

func (controller Controller) _getByCode(code string) (*models.CategoryClientGetSingle, error) {
	category, err := controller.categoryRepo.GetByCode(code)
	if err != nil {
		return nil, ErrorCategoryGetByCode(code)
	}

	if category == nil {
		return nil, ErrorCategoryNotFound(code)
	}

	return &models.CategoryClientGetSingle{
		Code:    category.Code,
		TitleRU: category.TitleRU,
		TitleKZ: category.TitleKZ,
		Image:   category.Icon,
		Status:  category.Status,
	}, nil
}

func (controller Controller) _getSubList(parentCode string, searchQuery string) ([]models.SubCategoryClientGet, error) {
	// Беру категорию по которой идет поиск
	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		return nil, ErrorCategoryGetByCode(parentCode)
	}

	if category == nil {
		return nil, ErrorCategoryNotFound(parentCode)
	}

	// Беру список по id категории
	var list []entities.SubCategory
	if len(searchQuery) == 0 {
		list, err = controller.subCategoryRepo.GetByParent(category.ID)
	} else {
		list, err = controller.subCategoryRepo.GetByQuery(category.ID, searchQuery)
	}

	subCategories := make([]models.SubCategoryClientGet, 0, len(list))
	for _, subCategory := range list {
		subCategories = append(subCategories, models.SubCategoryClientGet{
			Code:    subCategory.Code,
			TitleRU: subCategory.TitleRU,
			TitleKZ: subCategory.TitleKZ,
			Image:   subCategory.Icon,
		})
	}

	return subCategories, nil
}

func (controller Controller) _getSubSingle(parentCode string, code string) (*models.SubCategoryClientGet, error) {
	// Беру категорию по которой идет поиск
	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		return nil, ErrorCategoryGetByCode(code)
	}

	if category == nil {
		return nil, ErrorCategoryNotFound(code)
	}

	subCategory, err := controller.subCategoryRepo.GetByCode(code)
	if err != nil {
		return nil, ErrorSubCategoryGetByCode(err)
	}

	if subCategory == nil {
		return nil, ErrorSubCategoryNotFound(code)
	}

	return &models.SubCategoryClientGet{
		Code:    subCategory.Code,
		TitleRU: subCategory.TitleRU,
		TitleKZ: subCategory.TitleKZ,
		Image:   subCategory.Icon,
	}, nil
}
