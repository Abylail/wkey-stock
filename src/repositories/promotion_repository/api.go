package promotion_repository

import (
	"github.com/lowl11/lazy-entity/builders/delete_builder"
	"github.com/lowl11/lazy-entity/builders/select_builder"
	"github.com/lowl11/lazy-entity/builders/update_builder"
	"github.com/mehanizm/iuliia-go"
	"strings"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

func (repo *Repository) GetByCode(code string) (*entities.Promotion, error) {
	return repo.GetItem(func(builder *select_builder.Builder) {
		builder.Where(builder.Equal("code", code))
	})
}

func (repo *Repository) Create(model *models.PromotionAdminCreate) (*string, error) {
	// генерируем код категории
	code := strings.TrimSpace(strings.ToLower(iuliia.Wikipedia.Translate(model.TitleRU)))
	code = strings.ReplaceAll(code, " ", "_")

	_, err := repo.Add(entities.Promotion{
		Code:          code,
		TitleRU:       model.TitleRU,
		TitleKZ:       model.TitleKZ,
		DescriptionRU: model.DescriptionRU,
		DescriptionKZ: model.DescriptionKZ,
	})
	if err != nil {
		return nil, err
	}

	return &code, nil
}

func (repo *Repository) UpdateByCode(model *models.PromotionAdminUpdate) error {
	return repo.Update(func(builder *update_builder.Builder) {
		builder.Where(builder.Equal("code", model.Code))
	}, entities.Promotion{
		TitleRU:       model.TitleRU,
		TitleKZ:       model.TitleKZ,
		DescriptionRU: model.DescriptionRU,
		DescriptionKZ: model.DescriptionKZ,
	})
}

func (repo *Repository) UpdateImage(code string, imagePath string, lang string) error {
	var entity entities.Promotion
	if lang == "ru" {
		entity.ImageRU = &imagePath
	} else if lang == "kz" {
		entity.ImageKZ = &imagePath
	} else {
		return nil
	}

	return repo.Update(func(builder *update_builder.Builder) {
		builder.Where(builder.Equal("code", code))
	}, entity)
}

func (repo *Repository) DeleteByCode(code *string) error {
	return repo.Delete(func(builder *delete_builder.Builder) {
		builder.Equal("code", code)
	})
}
