package category_repository

import (
	"github.com/lowl11/lazy-entity/builders/delete_builder"
	"github.com/lowl11/lazy-entity/builders/select_builder"
	"wkey-stock/src/data/dtos"
	"wkey-stock/src/data/entities"
)

func (repo Repository) Create(category *dtos.Category) error {
	return repo.AddWithID(category.Entity())
}

func (repo Repository) All() ([]entities.Category, error) {
	return repo.GetList(func(builder *select_builder.Builder) {
		builder.Where(builder.Is("deleted", false))
	})
}

func (repo Repository) ByID(id string) (*entities.Category, error) {
	return repo.GetItem(func(builder *select_builder.Builder) {
		builder.Where(
			builder.Is("deleted", false),
			builder.Equal("id", id),
		)
	})
}

func (repo Repository) ByProsklad(proskladID int) (*entities.Category, error) {
	return repo.GetItem(func(builder *select_builder.Builder) {
		builder.Where(builder.Equal("prosklad_id", proskladID))
	})
}

func (repo Repository) UpdateCategory(category *dtos.Category) error {
	return repo.UpdateByID(category.ID().String(), category.Entity())
}

func (repo Repository) MakeDeleted(category *dtos.Category) error {
	entity := category.Entity()
	entity.Deleted = true
	return repo.UpdateByID(category.ID().String(), entity)
}

func (repo Repository) Remove(category *dtos.Category) error {
	return repo.DeleteByID(category.ID().String())
}

func (repo Repository) RemoveByProsklad(proskladID int) error {
	return repo.Delete(func(builder *delete_builder.Builder) {
		builder.Where(builder.Equal("prosklad_id", proskladID))
	})
}
