package product_repository

import (
	"context"
	"github.com/lowl11/lazy-entity/builders/delete_builder"
	"github.com/lowl11/lazy-entity/builders/select_builder"
	"wkey-stock/src/data/dtos"
	"wkey-stock/src/data/entities"
)

func (repo Repository) Create(ctx context.Context, product *dtos.Product) error {
	return repo.AddWithID(ctx, product.Entity())
}

func (repo Repository) All(ctx context.Context) ([]entities.Product, error) {
	return repo.GetList(ctx, func(builder *select_builder.Builder) {
		builder.Where(builder.Is("deleted", false))
	})
}

func (repo Repository) ByID(ctx context.Context, id string) (*entities.Product, error) {
	return repo.GetItem(ctx, func(builder *select_builder.Builder) {
		builder.Where(
			builder.Is("deleted", false),
			builder.Equal("id", id),
		)
	})
}

func (repo Repository) ByProsklad(ctx context.Context, proskladID int) (*entities.Product, error) {
	return repo.GetItem(ctx, func(builder *select_builder.Builder) {
		builder.Where(builder.Equal("prosklad_id", proskladID))
	})
}

func (repo Repository) Change(ctx context.Context, product *dtos.Product) error {
	return repo.UpdateByID(ctx, product.ID().String(), product.Entity())
}

func (repo Repository) MakeDeleted(ctx context.Context, product *dtos.Product) error {
	entity := product.Entity()
	entity.Deleted = true
	return repo.UpdateByID(ctx, product.ID().String(), entity)
}

func (repo Repository) Remove(ctx context.Context, product *dtos.Product) error {
	return repo.DeleteByID(ctx, product.ID().String())
}

func (repo Repository) RemoveByProsklad(ctx context.Context, proskladID int) error {
	return repo.Delete(ctx, func(builder *delete_builder.Builder) {
		builder.Where(builder.Equal("prosklad_id", proskladID))
	})
}
