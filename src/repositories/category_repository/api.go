package category_repository

import (
	"github.com/lowl11/lazy-entity/builders/delete_builder"
	"github.com/lowl11/lazy-entity/builders/select_builder"
	"github.com/lowl11/lazy-entity/builders/update_builder"
	"wkey-stock/src/data/entities"
	"wkey-stock/src/data/models"
)

const (
	statusActive   = "active"
	statusInactive = "inactive"
)

func (repo *Repository) GetByQuery(searchQuery string) ([]entities.Category, error) {
	searchQuery = "%" + searchQuery + "%"

	return repo.GetList(func(builder *select_builder.Builder) {
		builder.Where(
			builder.Or(
				builder.Equal("title_ru", searchQuery),
				builder.Equal("title_kz", searchQuery),
				builder.Equal("code", searchQuery),
			),
		)
	})
}

func (repo *Repository) GetByCode(code string) (*entities.Category, error) {
	return repo.GetItem(func(builder *select_builder.Builder) {
		builder.Where(builder.Equal("code", code))
	})
}

func (repo *Repository) Create(model *models.CategoryAdd, categoryCode string) error {
	_, err := repo.Add(entities.Category{
		Code:    categoryCode,
		TitleRU: model.TitleRU,
		TitleKZ: model.TitleKZ,
	})
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) UpdateByCode(code string, model *models.CategoryUpdate) error {
	return repo.Update(func(builder *update_builder.Builder) {
		builder.Where(builder.Equal("code", code))
	}, entities.Category{
		TitleRU: model.TitleRU,
		TitleKZ: model.TitleKZ,
	})
}

func (repo *Repository) UpdateImage(code string, imagePath string) error {
	return repo.Update(func(builder *update_builder.Builder) {
		builder.Where(builder.Equal("code", code))
	}, entities.Category{
		Icon: &imagePath,
	})
}

func (repo *Repository) Activate(code string) error {
	return repo.Update(func(builder *update_builder.Builder) {
		builder.Where(builder.Equal("code", code))
	}, entities.Category{
		Status: statusActive,
	})
}

func (repo *Repository) Deactivate(code string) error {
	return repo.Update(func(builder *update_builder.Builder) {
		builder.Where(builder.Equal("code", code))
	}, entities.Category{
		Status: statusInactive,
	})
}

func (repo *Repository) DeleteByCode(code string) error {
	return repo.Delete(func(builder *delete_builder.Builder) {
		builder.Where(builder.Equal("code", code))
	})
}
