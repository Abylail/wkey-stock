package category_controller

import (
	"github.com/lowl11/lazy-collection/type_list"
	"github.com/lowl11/lazylog/layers"
	"github.com/mehanizm/iuliia-go"
	"strings"
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
				Image:   item.Icon,
			}
		}).
		Slice()

	return categories, nil
}

func (controller *Controller) _create(model *models.CategoryAdd) *models.Error {
	logger := definition.Logger

	// генерируем код категории
	categoryCode := strings.TrimSpace(strings.ToLower(iuliia.Wikipedia.Translate(model.TitleRU)))

	// ищем категорию с таким же кодом
	category, err := controller.categoryRepo.GetByCode(categoryCode)
	if err != nil {
		logger.Error(err, "Get category by code error", layers.Database)
		return errors.CategoryGetByCode.With(err)
	}

	// выдаем ошибку если нашли категорию с таким же кодом
	if category != nil {
		return errors.CategoryAlreadyExist
	}

	// загружаем иконку
	fullPath, err := controller.image.UploadCategoryIcon(categoryCode, model.Image.Name, model.Image.Buffer)
	if err != nil {
		logger.Error(err, "Upload category icon error", layers.File)
		return errors.ImageUploadCategoryIcon.With(err)
	}

	// создаем запись в БД
	if err = controller.categoryRepo.Create(model, categoryCode, fullPath); err != nil {
		logger.Error(err, "Update category error", layers.Database)
		return errors.CategoryAdd.With(err)
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

func (controller *Controller) _upload(code string, model *models.CategoryUpload) (string, *models.Error) {
	logger := definition.Logger

	imagePath, err := controller.image.UploadCategoryIcon(code, model.Image.Name, model.Image.Buffer)
	if err != nil {
		logger.Error(err, "Upload category icon error", layers.File)
		return "", errors.ImageUploadCategoryIcon.With(err)
	}

	if err = controller.categoryRepo.UpdateImage(code, imagePath); err != nil {
		logger.Error(err, "Update category image error", layers.Database)
		return "", errors.CategoryUpdate.With(err)
	}

	return imagePath, nil
}

func (controller *Controller) _delete(categoryCode string) *models.Error {
	logger := definition.Logger

	if err := controller.categoryRepo.Delete(categoryCode); err != nil {
		logger.Error(err, "Delete category error", layers.Database)
		return errors.CategoryDelete.With(err)
	}

	return nil
}
