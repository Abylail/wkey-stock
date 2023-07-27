package brand_controller

import (
	"wkey-stock/src/adaptors/brand_adaptor"
	"wkey-stock/src/data/dtos"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

func (controller Controller) _get(searchQuery string) ([]dtos.Brand, error) {
	var brands []entities.Brand
	var err error

	if len(searchQuery) == 0 {
		brands, err = controller.brandRepo.GetAll()
	} else {
		brands, err = controller.brandRepo.GetByQuery(searchQuery)
	}
	if err != nil {
		return nil, ErrorBrandGetList(err)
	}

	return brand_adaptor.EntityToDTO(brands), nil
}

func (controller Controller) _getSingle(id int) (*dtos.Brand, error) {
	brand, err := controller.brandRepo.GetByID(id)
	if err != nil {
		return nil, ErrorBrandGetByID(err)
	}

	return dtos.NewBrand(brand), nil
}

func (controller Controller) _add(model *models.BrandAdd) error {
	brand, err := controller.brandRepo.GetByTitle(model.Title)
	if err != nil {
		return ErrorBrandGetByTitle(model.Title)
	}

	if brand != nil {
		return ErrorBrandAlreadyExist(model.Title)
	}

	if err = controller.brandRepo.Create(model); err != nil {
		return ErrorBrandAdd(err)
	}

	return nil
}

func (controller Controller) _update(id int, model *models.BrandUpdate) error {
	if err := controller.brandRepo.Update(id, model); err != nil {
		return ErrorBrandUpdate(err)
	}

	return nil
}

func (controller Controller) _upload(brandID int, model *models.BrandUpload) (string, error) {
	brand, err := controller.brandRepo.GetByID(brandID)
	if err != nil {
		return "", ErrorBrandGetByID(err)
	}

	if brand == nil {
		return "", ErrorBrandNotFound(brandID)
	}

	imagePath, err := controller.image.UploadBrandIcon(brandID, model.Image.Name, model.Image.Buffer)
	if err != nil {
		return "", ErrorBrandUpdateFileIcon(err)
	}

	if err = controller.brandRepo.UpdateIcon(brandID, imagePath); err != nil {
		return "", ErrorBrandUpdateIcon(err)
	}

	return imagePath, nil
}

func (controller Controller) _delete(id int) error {
	if err := controller.brandRepo.DeleteByID(id); err != nil {
		return ErrorBrandDelete(err)
	}

	return nil
}
