package promotion_repository

import (
	"context"
	"github.com/lowl11/lazy-entity/builders/select_builder"
	"wkey-stock/src/data/dtos"
	"wkey-stock/src/data/entities"
)

func (repo Repository) All(ctx context.Context) ([]entities.Promotion, error) {
	return repo.GetList(ctx, func(builder *select_builder.Builder) {
		builder.Where(
			builder.Is("deleted", false),
		)
	})
}

func (repo Repository) ByID(ctx context.Context, id string) (*entities.Promotion, error) {
	return repo.GetItem(ctx, func(builder *select_builder.Builder) {
		builder.Where(
			builder.Is("deleted", false),
			builder.Equal("id", id),
		)
	})
}

func (repo Repository) Create(ctx context.Context, subCategory *dtos.Promotion) error {
	return repo.AddWithID(ctx, subCategory.Entity())
}

func (repo Repository) Change(ctx context.Context, subCategory *dtos.Promotion) error {
	return repo.UpdateByID(ctx, subCategory.ID().String(), subCategory.Entity())
}

func (repo Repository) Remove(ctx context.Context, subCategory *dtos.Promotion) error {
	return repo.DeleteByID(ctx, subCategory.ID().String())
}
