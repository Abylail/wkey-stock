package category_controller

import (
	"github.com/mehanizm/iuliia-go"
	"strings"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

// _getClient список
func (controller *Controller) _getClient(searchQuery string) ([]models.CategoryClientGet, error) {
	var list []entities.Category
	var err error

	if len(searchQuery) == 0 {
		list, err = controller.categoryRepo.GetAll()
	} else {
		list, err = controller.categoryRepo.GetByQuery(searchQuery)
	}

	if err != nil {
		return nil, ErrorCategoryGetList()
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

func (controller *Controller) _getClientSingle(code string) (*models.CategoryClientGetSingle, error) {
	category, err := controller.categoryRepo.GetByCode(code)
	if err != nil {
		return nil, ErrorCategoryGetByCode()
	}

	if category == nil {
		return nil, ErrorCategoryNotFound()
	}

	return &models.CategoryClientGetSingle{
		Code:    category.Code,
		TitleRU: category.TitleRU,
		TitleKZ: category.TitleKZ,
		Image:   category.Icon,
		Status:  category.Status,
	}, nil
}

func (controller *Controller) _getClientSub(parentCode string, searchQuery string) ([]models.SubCategoryClientGet, error) {
	// Беру категорию по которой идет поиск
	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		return nil, ErrorCategoryGetByCode()
	}

	if category == nil {
		return nil, ErrorCategoryNotFound()
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

func (controller *Controller) _getClientSubSingle(parentCode string, code string) (*models.SubCategoryClientGet, error) {
	// Беру категорию по которой идет поиск
	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		return nil, ErrorCategoryGetByCode()
	}

	if category == nil {
		return nil, ErrorCategoryNotFound()
	}

	subCategory, err := controller.subCategoryRepo.GetByCode(code)
	if err != nil {
		return nil, ErrorSubCategoryGetByCode()
	}

	if subCategory == nil {
		return nil, ErrorSubCategoryNotFound()
	}

	return &models.SubCategoryClientGet{
		Code:    subCategory.Code,
		TitleRU: subCategory.TitleRU,
		TitleKZ: subCategory.TitleKZ,
		Image:   subCategory.Icon,
	}, nil
}

func (controller *Controller) _getAdmin(searchQuery string) ([]models.CategoryAdminItem, error) {
	var list []entities.Category
	var err error

	if searchQuery == "" {
		list, err = controller.categoryRepo.GetAll()
	} else {
		list, err = controller.categoryRepo.GetByQuery(searchQuery)
	}
	if err != nil {
		return nil, ErrorCategoryGetList()
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

func (controller *Controller) _getAdminSingle(code string) (*models.CategoryAdminGetSingle, error) {
	category, err := controller.categoryRepo.GetByCode(code)
	if err != nil {
		return nil, ErrorCategoryGetByCode()
	}

	if category == nil {
		return nil, ErrorCategoryNotFound()
	}

	subCategories, err := controller.subCategoryRepo.GetByParent(category.ID)
	if err != nil {
		return nil, ErrorSubCategoryGetList()
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

func (controller *Controller) _getAdminSub(parentCode string, searchQuery string) ([]models.SubCategoryAdminGet, error) {
	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		return nil, ErrorCategoryGetByCode()
	}

	if category == nil {
		return nil, ErrorCategoryNotFound()
	}

	var list []entities.SubCategory

	if len(searchQuery) == 0 {
		list, err = controller.subCategoryRepo.GetByParent(category.ID)
	} else {
		list, err = controller.subCategoryRepo.GetByQuery(category.ID, searchQuery)
	}
	if err != nil {
		return nil, ErrorSubCategoryGetList()
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

func (controller *Controller) _getAdminSubSingle(parentCode, code string) (*models.SubCategoryAdminGet, error) {
	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		return nil, ErrorCategoryGetByCode()
	}

	if category == nil {
		return nil, ErrorCategoryNotFound()
	}

	subCategory, err := controller.subCategoryRepo.GetByCode(code)
	if err != nil {
		return nil, ErrorSubCategoryGetByCode()
	}

	if subCategory == nil {
		return nil, ErrorSubCategoryNotFound()
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

func (controller *Controller) _create(model *models.CategoryAdd) (string, error) {
	// генерируем код категории
	categoryCode := strings.TrimSpace(strings.ToLower(iuliia.Wikipedia.Translate(model.TitleRU)))
	categoryCode = strings.ReplaceAll(categoryCode, " ", "_")

	// ищем категорию с таким же кодом
	category, err := controller.categoryRepo.GetByCode(categoryCode)
	if err != nil {
		return "", ErrorCategoryGetByCode()
	}

	// выдаем ошибку если нашли категорию с таким же кодом
	if category != nil {
		return "", ErrorCategoryAlreadyExist()
	}

	// создаем запись в БД
	if err = controller.categoryRepo.Create(model, categoryCode); err != nil {
		return "", ErrorCategoryAdd()
	}

	return categoryCode, nil
}

func (controller *Controller) _createSub(parentCode string, model *models.SubCategoryAdd) (string, error) {
	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		return "", ErrorCategoryGetByCode()
	}

	if category == nil {
		return "", ErrorCategoryNotFound()
	}

	// генерируем код категории
	categoryCode := strings.TrimSpace(strings.ToLower(iuliia.Wikipedia.Translate(model.TitleRU)))
	categoryCode = strings.ReplaceAll(categoryCode, " ", "_")

	// ищем категорию с таким же кодом
	subCategory, err := controller.subCategoryRepo.GetByCode(categoryCode)
	if err != nil {
		return "", ErrorSubCategoryGetByCode()
	}

	// выдаем ошибку если нашли категорию с таким же кодом
	if subCategory != nil {
		return "", ErrorSubCategoryAlreadyExist()
	}

	// создаем запись в БД
	if err = controller.subCategoryRepo.Create(category.ID, model, categoryCode); err != nil {
		return "", ErrorSubCategoryAdd()
	}

	return categoryCode, nil
}

func (controller *Controller) _update(code string, model *models.CategoryUpdate) error {
	if err := controller.categoryRepo.UpdateByCode(code, model); err != nil {
		return ErrorCategoryUpdate()
	}

	return nil
}

func (controller *Controller) _updateSub(parentCode, code string, model *models.SubCategoryUpdate) error {
	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		return ErrorCategoryGetByCode()
	}

	if category == nil {
		return ErrorCategoryNotFound()
	}

	if err = controller.subCategoryRepo.UpdateByParent(category.ID, code, model); err != nil {
		return ErrorSubCategoryUpdate()
	}

	return nil
}

func (controller *Controller) _upload(code string, model *models.CategoryUpload) (string, error) {
	category, err := controller.categoryRepo.GetByCode(code)
	if err != nil {
		return "", ErrorCategoryGetByCode()
	}

	if category == nil {
		return "", ErrorCategoryNotFound()
	}

	imagePath, err := controller.image.UploadCategoryIcon(code, model.Image.Name, model.Image.Buffer)
	if err != nil {
		return "", ErrorCategoryUpdateFileImage()
	}

	if err = controller.categoryRepo.UpdateImage(code, imagePath); err != nil {
		return "", ErrorCategoryUpdateImage()
	}

	return imagePath, nil
}

func (controller *Controller) _uploadSub(parentCode, code string, model *models.SubCategoryUpload) (string, error) {
	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		return "", ErrorCategoryGetByCode()
	}

	if category == nil {
		return "", ErrorCategoryNotFound()
	}

	imagePath, err := controller.image.UploadSubCategoryIcon(parentCode, code, model.Image.Name, model.Image.Buffer)
	if err != nil {
		return "", ErrorSubCategoryUpdateFileImage()
	}

	if err = controller.subCategoryRepo.UpdateImage(category.ID, code, imagePath); err != nil {
		return "", ErrorSubCategoryUpdateImage()
	}

	return imagePath, nil
}

func (controller *Controller) _delete(categoryCode string) error {
	category, err := controller.categoryRepo.GetByCode(categoryCode)
	if err != nil {
		return ErrorCategoryGetByCode()
	}

	if category == nil {
		return ErrorCategoryNotFound()
	}

	count, err := controller.subCategoryRepo.CountByParent(category.ID)
	if err != nil {
		return ErrorSubCategoryGetCount()
	}

	if count > 0 {
		return ErrorCategoryHasChildren()
	}

	if err = controller.categoryRepo.DeleteByCode(categoryCode); err != nil {
		return ErrorCategoryDelete()
	}

	return nil
}

func (controller *Controller) _deleteSub(parentCode, categoryCode string) error {
	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		return ErrorCategoryGetByCode()
	}

	if category == nil {
		return ErrorCategoryNotFound()
	}

	if err = controller.subCategoryRepo.DeleteByParent(category.ID, categoryCode); err != nil {
		return ErrorSubCategoryDelete()
	}

	return nil
}

func (controller *Controller) _activate(code string) error {
	if err := controller.categoryRepo.Activate(code); err != nil {
		return ErrorCategoryUpdateStatus()
	}

	return nil
}

func (controller *Controller) _deactivate(code string) error {
	if err := controller.categoryRepo.Deactivate(code); err != nil {
		return ErrorCategoryUpdateStatus()
	}

	return nil
}

func (controller *Controller) _activateSub(parentCode, code string) error {
	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		return ErrorCategoryGetByCode()
	}

	if category == nil {
		return ErrorCategoryNotFound()
	}

	subCategory, err := controller.subCategoryRepo.GetByCode(code)
	if err != nil {
		return ErrorCategoryGetByCode()
	}

	if subCategory == nil {
		return ErrorCategoryNotFound()
	}

	count, err := controller.productRepo.CountBySubCategory(subCategory.ID)
	if err != nil {
		return ErrorSubCategoryGetCount()
	}

	if count == 0 {
		return ErrorCategoryNoChildren()
	}

	if err = controller.subCategoryRepo.Activate(category.ID, code); err != nil {
		return ErrorSubCategoryUpdateStatus()
	}

	return nil
}

func (controller *Controller) _deactivateSub(parentCode, code string) error {
	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		return ErrorCategoryGetByCode()
	}

	if category == nil {
		return ErrorCategoryNotFound()
	}

	if err = controller.subCategoryRepo.Deactivate(category.ID, code); err != nil {
		return ErrorCategoryUpdateStatus()
	}

	return nil
}

func (controller *Controller) _bindProductList(parentCode, code string, model *models.SubCategoryBindProductList) error {
	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		return ErrorCategoryGetByCode()
	}

	if category == nil {
		return ErrorCategoryNotFound()
	}

	subCategory, err := controller.subCategoryRepo.GetByCode(code)
	if err != nil {
		return ErrorSubCategoryGetByCode()
	}

	if subCategory == nil {
		return ErrorSubCategoryNotFound()
	}

	if err = controller.productRepo.BindSubCategory(subCategory.ID, model.List); err != nil {
		return ErrorSubCategoryBindSub()
	}

	return nil
}

func (controller *Controller) _unbindProductItem(parentCode, code string, productID int) error {
	category, err := controller.categoryRepo.GetByCode(parentCode)
	if err != nil {
		return ErrorCategoryGetByCode()
	}

	if category == nil {
		return ErrorCategoryNotFound()
	}

	subCategory, err := controller.subCategoryRepo.GetByCode(code)
	if err != nil {
		return ErrorCategoryGetByCode()
	}

	if subCategory == nil {
		return ErrorCategoryNotFound()
	}

	if err = controller.productRepo.UnbindSubCategory(productID, subCategory.ID); err != nil {
		return ErrorSubCategoryUnbindSub()
	}

	return nil
}
