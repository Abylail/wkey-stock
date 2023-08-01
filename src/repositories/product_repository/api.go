package product_repository

import (
	"github.com/lowl11/lazy-entity/builders/delete_builder"
	"github.com/lowl11/lazy-entity/builders/select_builder"
	"wkey-stock/src/data/dtos"
	"wkey-stock/src/data/entities"
)

func (repo Repository) Create(product *dtos.Product) error {
	return repo.AddWithID(product.Entity())
}

func (repo Repository) All() ([]entities.Product, error) {
	return repo.GetList(func(builder *select_builder.Builder) {
		builder.Where(builder.Is("deleted", false))
	})
}

func (repo Repository) ByID(id string) (*entities.Product, error) {
	return repo.GetItem(func(builder *select_builder.Builder) {
		builder.Where(
			builder.Is("deleted", false),
			builder.Equal("id", id),
		)
	})
}

func (repo Repository) ByProsklad(proskladID int) (*entities.Product, error) {
	return repo.GetItem(func(builder *select_builder.Builder) {
		builder.Where(builder.Equal("prosklad_id", proskladID))
	})
}

func (repo Repository) UpdateProduct(product *dtos.Product) error {
	return repo.UpdateByID(product.ID().String(), product.Entity())
}

func (repo Repository) MakeDeleted(product *dtos.Product) error {
	entity := product.Entity()
	entity.Deleted = true
	return repo.UpdateByID(product.ID().String(), entity)
}

func (repo Repository) Remove(product *dtos.Product) error {
	return repo.DeleteByID(product.ID().String())
}

func (repo Repository) RemoveByProsklad(proskladID int) error {
	return repo.Delete(func(builder *delete_builder.Builder) {
		builder.Where(builder.Equal("prosklad_id", proskladID))
	})
}
