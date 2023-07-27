package sub_category_repository

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

func (repo *Repository) GetByParent(parentID int) ([]entities.SubCategory, error) {
	return repo.GetList(func(builder *select_builder.Builder) {
		builder.Where(builder.Equal("parent_id", parentID))
	})
}

func (repo *Repository) GetByQuery(parentID int, searchQuery string) ([]entities.SubCategory, error) {
	searchQuery = "%" + searchQuery + "%"

	return repo.GetList(func(builder *select_builder.Builder) {
		builder.Where(
			builder.Equal("parent_id", parentID),
			builder.Or(
				builder.ILike("title_ru", searchQuery),
				builder.ILike("title_kz", searchQuery),
				builder.ILike("code", searchQuery),
			),
		)
	})
}

func (repo *Repository) GetByCode(code string) (*entities.SubCategory, error) {
	return repo.GetItem(func(builder *select_builder.Builder) {
		builder.Where(builder.Equal("code", code))
	})
}

func (repo *Repository) Create(parentID int, model *models.SubCategoryAdd, categoryCode string) error {
	_, err := repo.Add(entities.SubCategory{
		Code:     categoryCode,
		TitleRU:  model.TitleRU,
		TitleKZ:  model.TitleKZ,
		ParentID: parentID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) UpdateByParent(parentID int, code string, model *models.SubCategoryUpdate) error {
	return repo.Update(func(builder *update_builder.Builder) {
		builder.Where(
			builder.Equal("code", code),
			builder.Equal("parent_id", parentID),
		)
	}, entities.SubCategory{
		TitleRU: model.TitleRU,
		TitleKZ: model.TitleKZ,
	})
}

func (repo *Repository) UpdateImage(parentID int, code, imagePath string) error {
	return repo.Update(func(builder *update_builder.Builder) {
		builder.Where(
			builder.Equal("code", code),
			builder.Equal("parent_id", parentID),
		)
	}, entities.SubCategory{
		Icon: &imagePath,
	})
}

func (repo *Repository) DeleteByParent(parentID int, code string) error {
	return repo.Delete(func(builder *delete_builder.Builder) {
		builder.Where(
			builder.Equal("code", code),
			builder.Equal("parent_id", parentID),
		)
	})
}

func (repo *Repository) CountByParent(parentID int) (int, error) {
	return repo.Count(func(builder *select_builder.Builder) {
		builder.Where(builder.Equal("parent_id", parentID))
	})
}

func (repo *Repository) Activate(parentID int, code string) error {
	return repo.Update(func(builder *update_builder.Builder) {
		builder.Where(
			builder.Equal("code", code),
			builder.Equal("parent_id", parentID),
		)
	}, entities.SubCategory{
		Status: statusInactive,
	})
}

func (repo *Repository) Deactivate(parentID int, code string) error {
	return repo.Update(func(builder *update_builder.Builder) {
		builder.Where(
			builder.Equal("code", code),
			builder.Equal("parent_id", parentID),
		)
	}, entities.SubCategory{
		Status: statusInactive,
	})
}
