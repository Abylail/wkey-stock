package promotion_controller

import (
	"github.com/lowl11/lazy-collection/type_list"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/errors"
	"wkey-stock/src/data/models"
)

// _getListAdmin список промоакций (в админке)
func (controller *Controller) _getListAdmin() ([]models.PromotionAdminGet, *models.Error) {
	list, err := controller.promotionRepo.GetAll()

	if err != nil {
		return nil, errors.PromotionGetList.With(err)
	}

	promotions := type_list.NewWithList[entities.AdminPromotion, models.PromotionAdminGet](list...).
		Select(func(item entities.AdminPromotion) models.PromotionAdminGet {
			return models.PromotionAdminGet{
				ID:            item.ID,
				CODE:          item.CODE,
				TitleRU:       item.TitleRU,
				TitleKZ:       item.TitleKZ,
				ImageRU:       item.ImageRU,
				ImageKZ:       item.ImageKZ,
				DescriptionRU: item.DescriptionRU,
				DescriptionKZ: item.DescriptionKZ,
			}
		}).
		Slice()

	return promotions, nil
}

// _getSingleAdmin промоакция по id
func (controller *Controller) _getSingleAdmin(id int) (*models.PromotionAdminGet, *models.Error) {
	rawPromotion, err := controller.promotionRepo.GetById(id)
	if err != nil {
		return nil, errors.PromotionGetById.With(err)
	}

	// Если не нашелся
	if rawPromotion == nil {
		return nil, errors.PromotionNotFound
	}

	return &models.PromotionAdminGet{
		ID:            rawPromotion.ID,
		CODE:          rawPromotion.CODE,
		TitleRU:       rawPromotion.TitleRU,
		TitleKZ:       rawPromotion.TitleKZ,
		ImageRU:       rawPromotion.ImageRU,
		ImageKZ:       rawPromotion.ImageKZ,
		DescriptionRU: rawPromotion.DescriptionRU,
		DescriptionKZ: rawPromotion.DescriptionKZ,
	}, nil
}

// _getSingleCodeAdmin промоакция по code
func (controller *Controller) _getSingleCodeAdmin(code string) (*models.PromotionAdminGet, *models.Error) {
	rawPromotion, err := controller.promotionRepo.GetByCode(code)
	if err != nil {
		return nil, errors.PromotionGetByCode.With(err)
	}

	// Если не нашелся
	if rawPromotion == nil {
		return nil, errors.PromotionNotFound
	}

	return &models.PromotionAdminGet{
		ID:            rawPromotion.ID,
		CODE:          rawPromotion.CODE,
		TitleRU:       rawPromotion.TitleRU,
		TitleKZ:       rawPromotion.TitleKZ,
		ImageRU:       rawPromotion.ImageRU,
		ImageKZ:       rawPromotion.ImageKZ,
		DescriptionRU: rawPromotion.DescriptionRU,
		DescriptionKZ: rawPromotion.DescriptionKZ,
	}, nil
}

// _createAdmin создание промоации
func (controller *Controller) _createAdmin(model *models.PromotionAdminCreate) (*string, *models.Error) {
	code, err := controller.promotionRepo.Create(model)
	if err != nil {
		return nil, errors.PromotionCreate.With(err)
	}
	return code, nil
}

// _updateAdmin обновление промоакции
func (controller *Controller) _updateAdmin(model *models.PromotionAdminUpdate) *models.Error {
	if err := controller.promotionRepo.Update(model); err != nil {
		return errors.PromotionUpdate.With(err)
	}
	return nil
}

// _uploadAdmin загрузка фотографий
func (controller *Controller) _uploadAdmin(model *models.PromotionAdminUpload) *models.Error {

	//if err := controller.promotionRepo.UpdateImage(model.Code, "imagepattext", model.Lang); err != nil {
	//	return errors.PromotionImageUpdate.With(err)
	//}

	promotion, err := controller.promotionRepo.GetByCode(model.Code)
	if err != nil {
		return errors.PromotionGetByCode.With(err)
	}

	if promotion == nil {
		return errors.PromotionNotFound
	}

	imagePath, err := controller.image.UploadPromotion(model.Code, model.Image.Name, model.Image.Buffer)
	if err != nil {
		return errors.PromotionImageUpload.With(err)
	}

	if err := controller.promotionRepo.UpdateImage(model.Code, imagePath, model.Lang); err != nil {
		return errors.PromotionImageUpdate.With(err)
	}

	// Удаляю старую фотку
	oldPath := promotion.ImageRU
	if model.Lang == "kz" {
		oldPath = promotion.ImageKZ
	}

	if oldPath != nil {
		if err := controller.image.Delete(*oldPath); err != nil {
			return errors.PromotionImageDelete.With(err)
		}
	}

	return nil
}

// _deleteAdmin удалить акцию
func (controller *Controller) _deleteAdmin(code *string) *models.Error {
	if err := controller.promotionRepo.Delete(code); err != nil {
		return errors.PromotionUpdate.With(err)
	}
	return nil
}
