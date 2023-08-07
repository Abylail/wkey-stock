package product_gateway

import (
	"context"
	"wkey-stock/src/adaptors/product_adaptor"
	"wkey-stock/src/data/dtos"
	"wkey-stock/src/data/models"
	"wkey-stock/src/enums/languages"
)

func (gateway Gateway) GetAll(ctx context.Context) ([]dtos.Product, error) {
	products, err := gateway.productRepo.All(ctx)
	if err != nil {
		return nil, ErrorGetAllProducts(err)
	}

	return product_adaptor.EntityToDTO(products), nil
}

func (gateway Gateway) GetByID(ctx context.Context, id string) (*dtos.Product, error) {
	productEntity, err := gateway.productRepo.ByID(ctx, id)
	if err != nil {
		return nil, ErrorGetByID(id, err)
	}

	if productEntity == nil {
		return nil, ErrorProductNotFoundID(id)
	}

	return dtos.NewProduct(productEntity), nil
}

func (gateway Gateway) GetByProsklad(ctx context.Context, proskladID int) (*dtos.Product, error) {
	productEntity, err := gateway.productRepo.ByProsklad(ctx, proskladID)
	if err != nil {
		return nil, ErrorGetByProsklad(proskladID, err)
	}

	if productEntity == nil {
		return nil, ErrorProductNotFoundProsklad(proskladID)
	}

	return dtos.NewProduct(productEntity), nil
}

func (gateway Gateway) Add(ctx context.Context, model *models.ProductProsklad) (string, error) {
	// пытаемся найти продукт по Prosklad ID
	productEntity, err := gateway.productRepo.ByProsklad(ctx, model.ID)
	if err != nil {
		return "", err
	}

	if productEntity != nil {
		return "", ErrorProductAlreadyExist(model.ID)
	}

	// создаем dto продукта
	product := dtos.NewProductAdd(model)

	// создаем запись в БД
	if err = gateway.productRepo.Create(ctx, product); err != nil {
		return "", ErrorAddProduct(err)
	}

	// возвращаем ID сгенерированного продукта
	return product.ID().String(), nil
}

func (gateway Gateway) UpdateProsklad(ctx context.Context, proskladID int, model *models.ProductProsklad) error {
	product, err := gateway.GetByProsklad(ctx, proskladID)
	if err != nil {
		return err
	}

	product.EditProsklad(model)

	if err = gateway.productRepo.Change(ctx, product); err != nil {
		return ErrorUpdateProduct(proskladID, err)
	}

	return nil
}

func (gateway Gateway) UpdateDescription(ctx context.Context, id string, model *models.ProductUpdateDescription) error {
	product, err := gateway.GetByID(ctx, id)
	if err != nil {
		return err
	}

	updateDTO := dtos.NewProductUpdateDescription(model)

	product.EditDescription(updateDTO.DescriptionRU(), languages.RU)
	product.EditDescription(updateDTO.DescriptionKZ(), languages.KZ)

	if err = gateway.productRepo.Change(ctx, product); err != nil {
		return ErrorUpdateProduct(id, err)
	}

	return nil
}

func (gateway Gateway) UpdateCount(ctx context.Context, id string, model *models.ProductUpdateCount) error {
	product, err := gateway.GetByID(ctx, id)
	if err != nil {
		return err
	}

	updateDTO := dtos.NewProductUpdateCount(model)

	product.EditCount(updateDTO.Count())

	if err = gateway.productRepo.Change(ctx, product); err != nil {
		return ErrorUpdateProduct(id, err)
	}

	return nil
}

func (gateway Gateway) UpdateImages(ctx context.Context, id string, model *models.ProductUpdateImages) error {
	product, err := gateway.GetByID(ctx, id)
	if err != nil {
		return err
	}

	product.EditImages(model.Primary, model.Secondary)

	if err = gateway.productRepo.Change(ctx, product); err != nil {
		return ErrorUpdateProduct(id, err)
	}

	return nil
}

func (gateway Gateway) RemoveByID(ctx context.Context, id string) error {
	product, err := gateway.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if err = gateway.productRepo.MakeDeleted(ctx, product); err != nil {
		return ErrorRemoveProductID(id, err)
	}

	return nil
}

func (gateway Gateway) RemoveByProsklad(ctx context.Context, proskladID int) error {
	product, err := gateway.GetByProsklad(ctx, proskladID)
	if err != nil {
		return ErrorGetByProsklad(proskladID, err)
	}

	if err = gateway.productRepo.Remove(ctx, product); err != nil {
		return ErrorRemoveProductProsklad(proskladID, err)
	}

	return nil
}
