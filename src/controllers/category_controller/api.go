package category_controller

import (
	"github.com/mehanizm/iuliia-go"
	"strings"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/errors"
	"wkey-stock/src/data/models"
)

// _getClient список
func (controller *Controller) _getClient(searchQuery string) ([]models.CategoryClientGet, *models.Error) {
	var list []entities.Category
	var err error

	if len(searchQuery) == 0 {
		list, err = controller.categoryRepo.GetAll()
	} else {
		list, err = controller.categoryRepo.GetByQuery(searchQuery)
	}

	if err != nil {
		return nil, errors.CategoryGetList.With(err)
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

func (controller *Controller) _getClientSingle(code string) (*models.CategoryClientGetSingle, *models.Error) {
	category, err := controller.categoryRepo.GetByCode(code)
	if err != nil {
		return nil, errors.CategoryGetByCode.With(err)
	}

	if category == nil {
		return nil, errors.CategoryNotFound
	}

	return &models.CategoryClientGetSingle{
		Code:    category.Code,
		TitleRU: category.TitleRU,
		TitleKZ: category.TitleKZ,
		Image:   category.Icon,
		Status:  category.Status,
	}, nil
}

func (controller *Controller) _getClientSub(parentCode string, searchQuery string) ([]models.SubCategoryClientGet, *models.Error) {
	// Беру категорию по которой идет поиск
	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		return nil, errors.CategoryGetByCode.With(err)
	}
	if category == nil {
		return nil, errors.CategoryNotFound
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

func (controller *Controller) _getClientSubSingle(parentCode string, code string) (*models.SubCategoryClientGet, *models.Error) {
	// Беру категорию по которой идет поиск
	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		return nil, errors.CategoryGetByCode.With(err)
	}

	if category == nil {
		return nil, errors.CategoryNotFound
	}

	subCategory, err := controller.subCategoryRepo.GetByCode(code)
	if err != nil {
		return nil, errors.CategoryGetByCode.With(err)
	}

	if subCategory == nil {
		return nil, errors.SubCategoryNotFound
	}

	return &models.SubCategoryClientGet{
		Code:    subCategory.Code,
		TitleRU: subCategory.TitleRU,
		TitleKZ: subCategory.TitleKZ,
		Image:   subCategory.Icon,
	}, nil
}

func (controller *Controller) _getAdmin(searchQuery string) ([]models.CategoryAdminItem, *models.Error) {
	var list []entities.Category
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

	categories := make([]models.CategoryAdminItem, 0, len(list))
	for _, category := range list {
		categories = append(categories, models.CategoryAdminItem{
			ID:      category.ID,
			Code:    category.Code,
			TitleRU: category.TitleRU,
			TitleKZ: category.TitleKZ,
			Image:   category.Icon,
			Status:  category.Status,
		})
	}

	return categories, nil
}

func (controller *Controller) _getAdminSingle(code string) (*models.CategoryAdminGetSingle, *models.Error) {
	category, err := controller.categoryRepo.GetByCode(code)
	if err != nil {
		return nil, errors.CategoryGetByCode.With(err)
	}

	if category == nil {
		return nil, errors.CategoryNotFound
	}

	subCategories, err := controller.subCategoryRepo.GetByParent(category.ID)
	if err != nil {
		return nil, errors.CategoryGetList.With(err)
	}

	subCategoriesList := make([]models.SubCategoryAdminGet, 0, len(subCategories))
	for _, subCategory := range subCategories {
		subCategoriesList = append(subCategoriesList, models.SubCategoryAdminGet{
			ID:      subCategory.ID,
			Code:    subCategory.Code,
			TitleRU: subCategory.TitleRU,
			TitleKZ: subCategory.TitleKZ,
			Image:   subCategory.Icon,
			Status:  subCategory.Status,
		})
	}

	return &models.CategoryAdminGetSingle{
		ID:      category.ID,
		Code:    category.Code,
		TitleRU: category.TitleRU,
		TitleKZ: category.TitleKZ,
		Image:   category.Icon,
		Status:  category.Status,

		SubCategories: subCategoriesList,
	}, nil
}

func (controller *Controller) _getAdminSub(parentCode string, searchQuery string) ([]models.SubCategoryAdminGet, *models.Error) {
	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		return nil, errors.CategoryGetByCode.With(err)
	}

	if category == nil {
		return nil, errors.CategoryNotFound
	}

	var list []entities.SubCategory

	if len(searchQuery) == 0 {
		list, err = controller.subCategoryRepo.GetByParent(category.ID)
	} else {
		list, err = controller.subCategoryRepo.GetByQuery(category.ID, searchQuery)
	}
	if err != nil {
		return nil, errors.CategoryGetList.With(err)
	}

	if len(list) == 0 {
		return []models.SubCategoryAdminGet{}, nil
	}

	subCategories := make([]models.SubCategoryAdminGet, 0, len(list))
	for _, subCategory := range list {
		subCategories = append(subCategories, models.SubCategoryAdminGet{
			ID:      subCategory.ID,
			Code:    subCategory.Code,
			TitleRU: subCategory.TitleRU,
			TitleKZ: subCategory.TitleKZ,
			Image:   subCategory.Icon,
			Status:  subCategory.Status,
		})
	}

	return subCategories, nil
}

func (controller *Controller) _getAdminSubSingle(parentCode, code string) (*models.SubCategoryAdminGet, *models.Error) {
	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		return nil, errors.CategoryGetByCode.With(err)
	}

	if category == nil {
		return nil, errors.CategoryNotFound
	}

	subCategory, err := controller.subCategoryRepo.GetByCode(code)
	if err != nil {
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
	// генерируем код категории
	categoryCode := strings.TrimSpace(strings.ToLower(iuliia.Wikipedia.Translate(model.TitleRU)))
	categoryCode = strings.ReplaceAll(categoryCode, " ", "_")

	// ищем категорию с таким же кодом
	category, err := controller.categoryRepo.GetByCode(categoryCode)
	if err != nil {
		return "", errors.CategoryGetByCode.With(err)
	}

	// выдаем ошибку если нашли категорию с таким же кодом
	if category != nil {
		return "", errors.CategoryAlreadyExist
	}

	// создаем запись в БД
	if err = controller.categoryRepo.Create(model, categoryCode); err != nil {
		return "", errors.CategoryAdd.With(err)
	}

	return categoryCode, nil
}

func (controller *Controller) _createSub(parentCode string, model *models.SubCategoryAdd) (string, *models.Error) {
	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		return "", errors.CategoryGetByCode.With(err)
	}

	if category == nil {
		return "", errors.CategoryNotFound
	}

	// генерируем код категории
	categoryCode := strings.TrimSpace(strings.ToLower(iuliia.Wikipedia.Translate(model.TitleRU)))
	categoryCode = strings.ReplaceAll(categoryCode, " ", "_")

	// ищем категорию с таким же кодом
	subCategory, err := controller.subCategoryRepo.GetByCode(categoryCode)
	if err != nil {
		return "", errors.CategoryGetByCode.With(err)
	}

	// выдаем ошибку если нашли категорию с таким же кодом
	if subCategory != nil {
		return "", errors.CategoryAlreadyExist
	}

	// создаем запись в БД
	if err = controller.subCategoryRepo.Create(category.ID, model, categoryCode); err != nil {
		return "", errors.CategoryAdd.With(err)
	}

	return categoryCode, nil
}

func (controller *Controller) _update(code string, model *models.CategoryUpdate) *models.Error {
	if err := controller.categoryRepo.UpdateByCode(code, model); err != nil {
		return errors.CategoryUpdate.With(err)
	}

	return nil
}

func (controller *Controller) _updateSub(parentCode, code string, model *models.SubCategoryUpdate) *models.Error {
	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		return errors.CategoryGetByCode.With(err)
	}

	if category == nil {
		return errors.CategoryNotFound
	}

	if err = controller.subCategoryRepo.UpdateByParent(category.ID, code, model); err != nil {
		return errors.CategoryUpdate.With(err)
	}

	return nil
}

func (controller *Controller) _upload(code string, model *models.CategoryUpload) (string, *models.Error) {
	category, err := controller.categoryRepo.GetByCode(code)
	if err != nil {
		return "", errors.CategoryGetByCode.With(err)
	}

	if category == nil {
		return "", errors.CategoryNotFound
	}

	imagePath, err := controller.image.UploadCategoryIcon(code, model.Image.Name, model.Image.Buffer)
	if err != nil {
		return "", errors.ImageUploadCategoryIcon.With(err)
	}

	if err = controller.categoryRepo.UpdateImage(code, imagePath); err != nil {
		return "", errors.CategoryUpdate.With(err)
	}

	return imagePath, nil
}

func (controller *Controller) _uploadSub(parentCode, code string, model *models.SubCategoryUpload) (string, *models.Error) {
	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		return "", errors.CategoryGetByCode.With(err)
	}

	if category == nil {
		return "", errors.CategoryNotFound
	}

	imagePath, err := controller.image.UploadSubCategoryIcon(parentCode, code, model.Image.Name, model.Image.Buffer)
	if err != nil {
		return "", errors.ImageUploadCategoryIcon.With(err)
	}

	if err = controller.subCategoryRepo.UpdateImage(category.ID, code, imagePath); err != nil {
		return "", errors.CategoryUpdate.With(err)
	}

	return imagePath, nil
}

func (controller *Controller) _delete(categoryCode string) *models.Error {
	category, err := controller.categoryRepo.GetByCode(categoryCode)
	if err != nil {
		return errors.CategoryGetByCode.With(err)
	}

	if category == nil {
		return errors.CategoryNotFound
	}

	count, err := controller.subCategoryRepo.CountByParent(category.ID)
	if err != nil {
		return errors.CategoryGetCount.With(err)
	}

	if count > 0 {
		return errors.CategoryHasSubCategories
	}

	if err = controller.categoryRepo.DeleteByCode(categoryCode); err != nil {
		return errors.CategoryDelete.With(err)
	}

	return nil
}

func (controller *Controller) _deleteSub(parentCode, categoryCode string) *models.Error {
	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		return errors.CategoryGetByCode.With(err)
	}

	if category == nil {
		return errors.CategoryNotFound
	}

	if err = controller.subCategoryRepo.DeleteByParent(category.ID, categoryCode); err != nil {
		return errors.CategoryDelete.With(err)
	}

	return nil
}

func (controller *Controller) _activate(code string) *models.Error {
	if err := controller.categoryRepo.Activate(code); err != nil {
		return errors.CategoryUpdateStatus.With(err)
	}

	return nil
}

func (controller *Controller) _deactivate(code string) *models.Error {
	if err := controller.categoryRepo.Deactivate(code); err != nil {
		return errors.CategoryUpdateStatus.With(err)
	}

	return nil
}

func (controller *Controller) _activateSub(parentCode, code string) *models.Error {
	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		return errors.CategoryGetByCode.With(err)
	}

	if category == nil {
		return errors.CategoryNotFound
	}

	subCategory, err := controller.subCategoryRepo.GetByCode(code)
	if err != nil {
		return errors.CategoryGetByCode.With(err)
	}

	if subCategory == nil {
		return errors.CategoryNotFound
	}

	count, err := controller.productRepo.CountBySubCategory(subCategory.ID)
	if err != nil {
		return errors.AdminProductCountGet.With(err)
	}

	if count == 0 {
		return errors.SubCategoryActivateEmpty
	}

	if err = controller.subCategoryRepo.Activate(category.ID, code); err != nil {
		return errors.SubCategoryUpdateStatus.With(err)
	}

	return nil
}

func (controller *Controller) _deactivateSub(parentCode, code string) *models.Error {
	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		return errors.CategoryGetByCode.With(err)
	}

	if category == nil {
		return errors.CategoryNotFound
	}

	if err = controller.subCategoryRepo.Deactivate(category.ID, code); err != nil {
		return errors.SubCategoryUpdateStatus.With(err)
	}

	return nil
}

func (controller *Controller) _bindProductList(parentCode, code string, model *models.SubCategoryBindProductList) *models.Error {
	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		return errors.CategoryGetByCode.With(err)
	}

	if category == nil {
		return errors.CategoryNotFound
	}

	subCategory, err := controller.subCategoryRepo.GetByCode(code)
	if err != nil {
		return errors.CategoryGetByCode.With(err)
	}

	if subCategory == nil {
		return errors.CategoryNotFound
	}

	if err = controller.productRepo.BindSubCategory(subCategory.ID, model.List); err != nil {
		return errors.SubCategoryBindProductList.With(err)
	}

	return nil
}

func (controller *Controller) _unbindProductItem(parentCode, code string, productID int) *models.Error {
	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		return errors.CategoryGetByCode.With(err)
	}

	if category == nil {
		return errors.CategoryNotFound
	}

	subCategory, err := controller.subCategoryRepo.GetByCode(code)
	if err != nil {
		return errors.CategoryGetByCode.With(err)
	}

	if subCategory == nil {
		return errors.CategoryNotFound
	}

	if err = controller.productRepo.UnbindSubCategory(productID, subCategory.ID); err != nil {
		return errors.SubCategoryUnbindProductItem.With(err)
	}

	return nil
}
