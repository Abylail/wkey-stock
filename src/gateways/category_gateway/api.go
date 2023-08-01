package category_gateway

import (
	"wkey-stock/src/adaptors/category_adaptor"
	"wkey-stock/src/data/dtos"
	"wkey-stock/src/data/models"
)

func (gateway Gateway) GetAll() ([]dtos.Category, error) {
	categories, err := gateway.categoryRepo.All()
	if err != nil {
		return nil, ErrorGetAllCategories(err)
	}

	return category_adaptor.EntityToDTO(categories), nil
}

func (gateway Gateway) GetByID(id string) (*dtos.Category, error) {
	categoryEntity, err := gateway.categoryRepo.ByID(id)
	if err != nil {
		return nil, ErrorGetByID(id, err)
	}

	if categoryEntity == nil {
		return nil, ErrorCategoryNotFound(id)
	}

	return dtos.NewCategory(categoryEntity), nil
}

func (gateway Gateway) GetByProsklad(proskladID int) (*dtos.Category, error) {
	categoryEntity, err := gateway.categoryRepo.ByProsklad(proskladID)
	if err != nil {
		return nil, ErrorGetByProsklad(proskladID, err)
	}

	if categoryEntity == nil {
		return nil, ErrorCategoryNotFound(proskladID)
	}

	return dtos.NewCategory(categoryEntity), nil
}

func (gateway Gateway) Add(model *models.CategoryProsklad) (string, error) {
	// пытаемся найти категорию по Prosklad ID
	categoryEntity, err := gateway.categoryRepo.ByProsklad(model.ID)
	if err != nil {
		return "", err
	}

	if categoryEntity != nil {
		return "", ErrorCategoryAlreadyExist(model.ID)
	}

	// создаем dto категории
	category := dtos.NewCategoryAdd(model)

	// создаем запись в БД
	if err = gateway.categoryRepo.Create(category); err != nil {
		return "", ErrorAddCategory(err)
	}

	// возвращаем ID сгенерированного продукта
	return category.ID().String(), nil
}

func (gateway Gateway) UpdateProsklad(proskladID int, model *models.CategoryProsklad) error {
	category, err := gateway.GetByProsklad(proskladID)
	if err != nil {
		return err
	}

	category.EditProsklad(model)

	if err = gateway.categoryRepo.UpdateCategory(category); err != nil {
		return ErrorUpdateCategory(proskladID, err)
	}

	return nil
}

func (gateway Gateway) UpdateCount(id string, model *models.CategoryUpdateCount) error {
	category, err := gateway.GetByID(id)
	if err != nil {
		return err
	}

	updateDTO := dtos.NewCategoryUpdateCount(model)

	category.EditCount(updateDTO.Count())

	if err = gateway.categoryRepo.UpdateCategory(category); err != nil {
		return ErrorUpdateCategory(id, err)
	}

	return nil
}

func (gateway Gateway) RemoveByID(id string) error {
	category, err := gateway.GetByID(id)
	if err != nil {
		return err
	}

	if err = gateway.categoryRepo.MakeDeleted(category); err != nil {
		return ErrorRemoveCategory(id, err)
	}

	return nil
}

func (gateway Gateway) RemoveByProsklad(proskladID int) error {
	category, err := gateway.GetByProsklad(proskladID)
	if err != nil {
		return ErrorGetByProsklad(proskladID, err)
	}

	if err = gateway.categoryRepo.Remove(category); err != nil {
		return ErrorRemoveCategory(proskladID, err)
	}

	return nil
}
