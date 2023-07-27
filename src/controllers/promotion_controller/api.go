package promotion_controller

import (
	"wkey-stock/src/data/models"
)

// _getListAdmin список промоакций (в админке)
func (controller *Controller) _getListAdmin() ([]models.PromotionAdminGet, error) {
	list, err := controller.promotionRepo.GetAll()

	if err != nil {
		return nil, ErrorPromotionGetList()
	}

	promotions := make([]models.PromotionAdminGet, 0, len(list))
	for _, promotion := range list {
		promotions = append(promotions, models.PromotionAdminGet{
			ID:            promotion.ID,
			CODE:          promotion.Code,
			TitleRU:       promotion.TitleRU,
			TitleKZ:       promotion.TitleKZ,
			ImageRU:       promotion.ImageRU,
			ImageKZ:       promotion.ImageKZ,
			DescriptionRU: promotion.DescriptionRU,
			DescriptionKZ: promotion.DescriptionKZ,
		})
	}

	return promotions, nil
}

// _getSingleAdmin промоакция по id
func (controller *Controller) _getSingleAdmin(id int) (*models.PromotionAdminGet, error) {
	rawPromotion, err := controller.promotionRepo.GetByID(id)
	if err != nil {
		return nil, ErrorPromotionGetByID()
	}

	// Если не нашелся
	if rawPromotion == nil {
		return nil, ErrorPromotionNotFound()
	}

	return &models.PromotionAdminGet{
		ID:            rawPromotion.ID,
		CODE:          rawPromotion.Code,
		TitleRU:       rawPromotion.TitleRU,
		TitleKZ:       rawPromotion.TitleKZ,
		ImageRU:       rawPromotion.ImageRU,
		ImageKZ:       rawPromotion.ImageKZ,
		DescriptionRU: rawPromotion.DescriptionRU,
		DescriptionKZ: rawPromotion.DescriptionKZ,
	}, nil
}

// _getSingleCodeAdmin промоакция по code
func (controller *Controller) _getSingleCodeAdmin(code string) (*models.PromotionAdminGet, error) {
	rawPromotion, err := controller.promotionRepo.GetByCode(code)
	if err != nil {
		return nil, ErrorPromotionGetByCode()
	}

	// Если не нашелся
	if rawPromotion == nil {
		return nil, ErrorPromotionNotFound()
	}

	return &models.PromotionAdminGet{
		ID:            rawPromotion.ID,
		CODE:          rawPromotion.Code,
		TitleRU:       rawPromotion.TitleRU,
		TitleKZ:       rawPromotion.TitleKZ,
		ImageRU:       rawPromotion.ImageRU,
		ImageKZ:       rawPromotion.ImageKZ,
		DescriptionRU: rawPromotion.DescriptionRU,
		DescriptionKZ: rawPromotion.DescriptionKZ,
	}, nil
}

// _createAdmin создание промоации
func (controller *Controller) _createAdmin(model *models.PromotionAdminCreate) (*string, error) {
	code, err := controller.promotionRepo.Create(model)
	if err != nil {
		return nil, ErrorPromotionAdd()
	}

	return code, nil
}

// _updateAdmin обновление промоакции
func (controller *Controller) _updateAdmin(model *models.PromotionAdminUpdate) error {
	if err := controller.promotionRepo.UpdateByCode(model); err != nil {
		return ErrorPromotionUpdate()
	}
	return nil
}

// _uploadAdmin загрузка фотографий
func (controller *Controller) _uploadAdmin(model *models.PromotionAdminUpload) error {

	//if err := controller.promotionRepo.UpdateImage(model.Code, "imagepattext", model.Lang); err != nil {
	//	return errors.PromotionImageUpdate.With(err)
	//}

	promotion, err := controller.promotionRepo.GetByCode(model.Code)
	if err != nil {
		return ErrorPromotionGetByCode()
	}

	if promotion == nil {
		return ErrorPromotionNotFound()
	}

	imagePath, err := controller.image.UploadPromotion(model.Code, model.Image.Name, model.Image.Buffer)
	if err != nil {
		return ErrorPromotionUpdateFileImages()
	}

	if err = controller.promotionRepo.UpdateImage(model.Code, imagePath, model.Lang); err != nil {
		return ErrorPromotionUpdateImages()
	}

	// Удаляю старую фотку
	oldPath := promotion.ImageRU
	if model.Lang == "kz" {
		oldPath = promotion.ImageKZ
	}

	if oldPath != nil {
		if err = controller.image.Delete(*oldPath); err != nil {
			return ErrorPromotionDelete()
		}
	}

	return nil
}

// _deleteAdmin удалить акцию
func (controller *Controller) _deleteAdmin(code *string) error {
	if err := controller.promotionRepo.DeleteByCode(code); err != nil {
		return ErrorPromotionUpdate()
	}

	if err := controller.image.DeletePromotionFolder(*code); err != nil {
		return ErrorPromotionDeleteFolder()
	}

	return nil
}

// _getListClient список промоакций (в админке)
func (controller *Controller) _getListClient() ([]models.PromotionClinetGet, error) {
	list, err := controller.promotionRepo.GetAll()

	if err != nil {
		return nil, ErrorPromotionGetList()
	}

	promotions := make([]models.PromotionClinetGet, 0, len(list))
	for _, promotion := range list {
		promotions = append(promotions, models.PromotionClinetGet{
			CODE:          promotion.Code,
			TitleRU:       promotion.TitleRU,
			TitleKZ:       promotion.TitleKZ,
			ImageRU:       promotion.ImageRU,
			ImageKZ:       promotion.ImageKZ,
			DescriptionRU: promotion.DescriptionRU,
			DescriptionKZ: promotion.DescriptionKZ,
		})
	}

	return promotions, nil
}

// _getSingleCodeAdmin промоакция по code
func (controller *Controller) _getSingleClient(code string) (*models.PromotionClinetGet, error) {
	rawPromotion, err := controller.promotionRepo.GetByCode(code)
	if err != nil {
		return nil, ErrorPromotionGetByCode()
	}

	// Если не нашелся
	if rawPromotion == nil {
		return nil, ErrorPromotionNotFound()
	}

	return &models.PromotionClinetGet{
		CODE:          rawPromotion.Code,
		TitleRU:       rawPromotion.TitleRU,
		TitleKZ:       rawPromotion.TitleKZ,
		ImageRU:       rawPromotion.ImageRU,
		ImageKZ:       rawPromotion.ImageKZ,
		DescriptionRU: rawPromotion.DescriptionRU,
		DescriptionKZ: rawPromotion.DescriptionKZ,
	}, nil
}
