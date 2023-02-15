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

func (controller *Controller) _getClientSub() ([]models.SubCategoryClientGet, *models.Error) {
	return nil, nil
}

func (controller *Controller) _getAdmin(searchQuery string) ([]models.CategoryAdminItem, *models.Error) {
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
		return []models.CategoryAdminItem{}, nil
	}

	categories := type_list.NewWithList[entities.CategoryGet, models.CategoryAdminItem](list...).
		Select(func(item entities.CategoryGet) models.CategoryAdminItem {
			return models.CategoryAdminItem{
				ID:      item.ID,
				Code:    item.Code,
				TitleRU: item.TitleRU,
				TitleKZ: item.TitleKZ,
				Image:   item.Icon,
				Status:  item.Status,
			}
		}).
		Slice()

	return categories, nil
}

func (controller *Controller) _getAdminSingle(code string) (*models.CategoryAdminGetSingle, *models.Error) {
	logger := definition.Logger

	category, err := controller.categoryRepo.GetByCode(code)
	if err != nil {
		logger.Error(err, "Get category by code error", layers.Database)
		return nil, errors.CategoryGetByCode.With(err)
	}

	if category == nil {
		return nil, errors.CategoryNotFound
	}

	subCategories, err := controller.subCategoryRepo.GetByParent(category.ID)
	if err != nil {
		logger.Error(err, "Get sub categories list error", layers.Database)
		return nil, errors.CategoryGetList.With(err)
	}

	return &models.CategoryAdminGetSingle{
		ID:      category.ID,
		Code:    category.Code,
		TitleRU: category.TitleRU,
		TitleKZ: category.TitleKZ,
		Image:   category.Icon,
		Status:  category.Status,

		SubCategories: type_list.NewWithList[entities.SubCategoryGet, models.SubCategoryAdminGet](subCategories...).
			Select(func(item entities.SubCategoryGet) models.SubCategoryAdminGet {
				return models.SubCategoryAdminGet{
					ID:      item.ID,
					Code:    item.Code,
					TitleRU: item.TitleRU,
					TitleKZ: item.TitleKZ,
					Image:   item.Icon,
					Status:  item.Status,
				}
			}).
			Slice(),
	}, nil
}

func (controller *Controller) _getAdminSub(parentCode string, searchQuery string) ([]models.SubCategoryAdminGet, *models.Error) {
	logger := definition.Logger

	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		logger.Error(err, "Get category by code error", layers.Database)
		return nil, errors.CategoryGetByCode.With(err)
	}

	if category == nil {
		return nil, errors.CategoryNotFound
	}

	var list []entities.SubCategoryGet

	if len(searchQuery) == 0 {
		list, err = controller.subCategoryRepo.GetByParent(category.ID)
	} else {
		list, err = controller.subCategoryRepo.GetByQuery(category.ID, searchQuery)
	}
	if err != nil {
		logger.Error(err, "Get sub categories list error", layers.Database)
		return nil, errors.CategoryGetList.With(err)
	}

	if len(list) == 0 {
		return []models.SubCategoryAdminGet{}, nil
	}

	return type_list.NewWithList[entities.SubCategoryGet, models.SubCategoryAdminGet](list...).
		Select(func(item entities.SubCategoryGet) models.SubCategoryAdminGet {
			return models.SubCategoryAdminGet{
				ID:      item.ID,
				Code:    item.Code,
				TitleRU: item.TitleRU,
				TitleKZ: item.TitleKZ,
				Image:   item.Icon,
				Status:  item.Status,
			}
		}).
		Slice(), nil
}

func (controller *Controller) _getAdminSubSingle(parentCode, code string) (*models.SubCategoryAdminGet, *models.Error) {
	logger := definition.Logger

	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		logger.Error(err, "Get category by code error", layers.Database)
		return nil, errors.CategoryGetByCode.With(err)
	}

	if category == nil {
		return nil, errors.CategoryNotFound
	}

	subCategory, err := controller.subCategoryRepo.GetByCode(category.ID, code)
	if err != nil {
		logger.Error(err, "Get sub category by code error", layers.Database)
		return nil, errors.CategoryGetByCode.With(err)
	}

	if subCategory == nil {
		return nil, errors.SubCategoryNotFound
	}

	return &models.SubCategoryAdminGet{
		ID:      subCategory.ID,
		Code:    subCategory.Code,
		TitleRU: subCategory.TitleRU,
		TitleKZ: subCategory.TitleKZ,
		Image:   subCategory.Icon,
		Status:  subCategory.Status,
	}, nil
}

func (controller *Controller) _create(model *models.CategoryAdd) (string, *models.Error) {
	logger := definition.Logger

	// генерируем код категории
	categoryCode := strings.TrimSpace(strings.ToLower(iuliia.Wikipedia.Translate(model.TitleRU)))
	categoryCode = strings.ReplaceAll(categoryCode, " ", "_")

	// ищем категорию с таким же кодом
	category, err := controller.categoryRepo.GetByCode(categoryCode)
	if err != nil {
		logger.Error(err, "Get category by code error", layers.Database)
		return "", errors.CategoryGetByCode.With(err)
	}

	// выдаем ошибку если нашли категорию с таким же кодом
	if category != nil {
		return "", errors.CategoryAlreadyExist
	}

	// создаем запись в БД
	if err = controller.categoryRepo.Create(model, categoryCode); err != nil {
		logger.Error(err, "Create category error", layers.Database)
		return "", errors.CategoryAdd.With(err)
	}

	return categoryCode, nil
}

func (controller *Controller) _createSub(parentCode string, model *models.SubCategoryAdd) (string, *models.Error) {
	logger := definition.Logger

	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		logger.Error(err, "Get category by code error", layers.Database)
		return "", errors.CategoryGetByCode.With(err)
	}

	if category == nil {
		return "", errors.CategoryNotFound
	}

	// генерируем код категории
	categoryCode := strings.TrimSpace(strings.ToLower(iuliia.Wikipedia.Translate(model.TitleRU)))
	categoryCode = strings.ReplaceAll(categoryCode, " ", "_")

	// ищем категорию с таким же кодом
	subCategory, err := controller.subCategoryRepo.GetByCode(category.ID, categoryCode)
	if err != nil {
		logger.Error(err, "Get sub category by code error", layers.Database)
		return "", errors.CategoryGetByCode.With(err)
	}

	// выдаем ошибку если нашли категорию с таким же кодом
	if subCategory != nil {
		return "", errors.CategoryAlreadyExist
	}

	// создаем запись в БД
	if err = controller.subCategoryRepo.Create(category.ID, model, categoryCode); err != nil {
		logger.Error(err, "Create sub category error", layers.Database)
		return "", errors.CategoryAdd.With(err)
	}

	return categoryCode, nil
}

func (controller *Controller) _update(code string, model *models.CategoryUpdate) *models.Error {
	logger := definition.Logger

	if err := controller.categoryRepo.Update(code, model); err != nil {
		logger.Error(err, "Update category error", layers.Database)
		return errors.CategoryUpdate.With(err)
	}

	return nil
}

func (controller *Controller) _updateSub(parentCode, code string, model *models.SubCategoryUpdate) *models.Error {
	logger := definition.Logger

	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		logger.Error(err, "Get category by code error", layers.Database)
		return errors.CategoryGetByCode.With(err)
	}

	if category == nil {
		return errors.CategoryNotFound
	}

	if err = controller.subCategoryRepo.Update(category.ID, code, model); err != nil {
		logger.Error(err, "Update sub category error", layers.Database)
		return errors.CategoryUpdate.With(err)
	}

	return nil
}

func (controller *Controller) _upload(code string, model *models.CategoryUpload) (string, *models.Error) {
	logger := definition.Logger

	category, err := controller.categoryRepo.GetByCode(code)
	if err != nil {
		logger.Error(err, "Get category by code error", layers.Database)
		return "", errors.CategoryGetByCode.With(err)
	}

	if category == nil {
		return "", errors.CategoryNotFound
	}

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

func (controller *Controller) _uploadSub(parentCode, code string, model *models.SubCategoryUpload) (string, *models.Error) {
	logger := definition.Logger

	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		logger.Error(err, "Get category by code error", layers.Database)
		return "", errors.CategoryGetByCode.With(err)
	}

	if category == nil {
		return "", errors.CategoryNotFound
	}

	imagePath, err := controller.image.UploadSubCategoryIcon(parentCode, code, model.Image.Name, model.Image.Buffer)
	if err != nil {
		logger.Error(err, "Upload sub category icon error", layers.File)
		return "", errors.ImageUploadCategoryIcon.With(err)
	}

	if err = controller.subCategoryRepo.UpdateImage(code, imagePath); err != nil {
		logger.Error(err, "Update sub category image error", layers.Database)
		return "", errors.CategoryUpdate.With(err)
	}

	return imagePath, nil
}

func (controller *Controller) _delete(categoryCode string) *models.Error {
	logger := definition.Logger

	category, err := controller.categoryRepo.GetByCode(categoryCode)
	if err != nil {
		logger.Error(err, "Get category by code error", layers.Database)
		return errors.CategoryGetByCode.With(err)
	}

	if category == nil {
		return errors.CategoryNotFound
	}

	count, err := controller.subCategoryRepo.Count(category.ID)
	if err != nil {
		logger.Error(err, "Get count sub categories by parent id error", layers.Database)
		return errors.CategoryGetCount.With(err)
	}

	if count > 0 {
		return errors.CategoryHasSubCategories
	}

	if err = controller.categoryRepo.Delete(categoryCode); err != nil {
		logger.Error(err, "Delete category error", layers.Database)
		return errors.CategoryDelete.With(err)
	}

	return nil
}

func (controller *Controller) _deleteSub(parentCode, categoryCode string) *models.Error {
	logger := definition.Logger

	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		logger.Error(err, "Get category by code error", layers.Database)
		return errors.CategoryGetByCode.With(err)
	}

	if category == nil {
		return errors.CategoryNotFound
	}

	if err = controller.subCategoryRepo.Delete(category.ID, categoryCode); err != nil {
		logger.Error(err, "Delete sub category error", layers.Database)
		return errors.CategoryDelete.With(err)
	}

	return nil
}

func (controller *Controller) _activate(code string) *models.Error {
	logger := definition.Logger

	if err := controller.categoryRepo.Activate(code); err != nil {
		logger.Error(err, "Update category status 'active' error", layers.Database)
		return errors.CategoryUpdateStatus.With(err)
	}

	return nil
}

func (controller *Controller) _deactivate(code string) *models.Error {
	logger := definition.Logger

	if err := controller.categoryRepo.Deactivate(code); err != nil {
		logger.Error(err, "Update category status 'inactive' error", layers.Database)
		return errors.CategoryUpdateStatus.With(err)
	}

	return nil
}

func (controller *Controller) _activateSub(parentCode, code string) *models.Error {
	logger := definition.Logger

	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		logger.Error(err, "Get category by code error", layers.Database)
		return errors.CategoryGetByCode.With(err)
	}

	if category == nil {
		return errors.CategoryNotFound
	}

	if err = controller.subCategoryRepo.Activate(category.ID, code); err != nil {
		logger.Error(err, "Update sub category status 'active' error", layers.Database)
		return errors.SubCategoryUpdateStatus.With(err)
	}

	return nil
}

func (controller *Controller) _deactivateSub(parentCode, code string) *models.Error {
	logger := definition.Logger

	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		logger.Error(err, "Get category by code error", layers.Database)
		return errors.CategoryGetByCode.With(err)
	}

	if category == nil {
		return errors.CategoryNotFound
	}

	if err = controller.subCategoryRepo.Deactivate(category.ID, code); err != nil {
		logger.Error(err, "Update sub category status 'inactive' error", layers.Database)
		return errors.SubCategoryUpdateStatus.With(err)
	}

	return nil
}
