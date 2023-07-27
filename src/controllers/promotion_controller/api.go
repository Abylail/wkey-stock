package promotion_controller

import (
	"wkey-stock/src/adaptors/promotion_adaptor"
	"wkey-stock/src/data/dtos"
	"wkey-stock/src/data/models"
)

// _getListAdmin список промоакций (в админке)
func (controller *Controller) _getListAdmin() ([]dtos.Promotion, error) {
	list, err := controller.promotionRepo.GetAll()
	if err != nil {
		return nil, ErrorPromotionGetList(err)
	}

	return promotion_adaptor.EntityToDTO(list), nil
}

// _getSingleAdmin промоакция по id
func (controller *Controller) _getSingleAdmin(id int) (*dtos.Promotion, error) {
	promotion, err := controller.promotionRepo.GetByID(id)
	if err != nil {
		return nil, ErrorPromotionGetByID(err)
	}

	// Если не нашелся
	if promotion == nil {
		return nil, ErrorPromotionNotFoundByID(id)
	}

	return dtos.NewPromotion(promotion), nil
}

// _getSingleCodeAdmin промоакция по code
func (controller *Controller) _getSingleCodeAdmin(code string) (*models.PromotionGet, error) {
	rawPromotion, err := controller.promotionRepo.GetByCode(code)
	if err != nil {
		return nil, ErrorPromotionGetByCode(err)
	}

	// Если не нашелся
	if rawPromotion == nil {
		return nil, ErrorPromotionNotFoundByCode(code)
	}

	return &models.PromotionGet{
		ID:            rawPromotion.ID,
		Code:          rawPromotion.Code,
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
		return nil, ErrorPromotionAdd(err)
	}

	return code, nil
}

// _updateAdmin обновление промоакции
func (controller *Controller) _updateAdmin(model *models.PromotionAdminUpdate) error {
	if err := controller.promotionRepo.UpdateByCode(model); err != nil {
		return ErrorPromotionUpdate(err)
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
		return ErrorPromotionGetByCode(err)
	}

	if promotion == nil {
		return ErrorPromotionNotFoundByCode(model.Code)
	}

	imagePath, err := controller.image.UploadPromotion(model.Code, model.Image.Name, model.Image.Buffer)
	if err != nil {
		return ErrorPromotionUpdateFileImages(err)
	}

	if err = controller.promotionRepo.UpdateImage(model.Code, imagePath, model.Lang); err != nil {
		return ErrorPromotionUpdateImages(err)
	}

	// Удаляю старую фотку
	oldPath := promotion.ImageRU
	if model.Lang == "kz" {
		oldPath = promotion.ImageKZ
	}

	if oldPath != nil {
		if err = controller.image.Delete(*oldPath); err != nil {
			return ErrorPromotionDelete(err)
		}
	}

	return nil
}

// _deleteAdmin удалить акцию
func (controller *Controller) _deleteAdmin(code *string) error {
	if err := controller.promotionRepo.DeleteByCode(code); err != nil {
		return ErrorPromotionUpdate(err)
	}

	if err := controller.image.DeletePromotionFolder(*code); err != nil {
		return ErrorPromotionDeleteFolder(err)
	}

	return nil
}

// _getListClient список промоакций (в админке)
func (controller *Controller) _getListClient() ([]models.PromotionGet, error) {
	list, err := controller.promotionRepo.GetAll()
	if err != nil {
		return nil, ErrorPromotionGetList(err)
	}

	promotions := make([]models.PromotionGet, 0, len(list))
	for _, promotion := range list {
		promotions = append(promotions, models.PromotionGet{
			Code:          promotion.Code,
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
func (controller *Controller) _getSingleClient(code string) (*models.PromotionGet, error) {
	rawPromotion, err := controller.promotionRepo.GetByCode(code)
	if err != nil {
		return nil, ErrorPromotionGetByCode(err)
	}

	// Если не нашелся
	if rawPromotion == nil {
		return nil, ErrorPromotionNotFoundByCode(code)
	}

	return &models.PromotionGet{
		Code:          rawPromotion.Code,
		TitleRU:       rawPromotion.TitleRU,
		TitleKZ:       rawPromotion.TitleKZ,
		ImageRU:       rawPromotion.ImageRU,
		ImageKZ:       rawPromotion.ImageKZ,
		DescriptionRU: rawPromotion.DescriptionRU,
		DescriptionKZ: rawPromotion.DescriptionKZ,
	}, nil
}
