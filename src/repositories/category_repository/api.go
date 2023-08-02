package category_repository

import (
	"context"
	"github.com/lowl11/lazy-entity/builders/delete_builder"
	"github.com/lowl11/lazy-entity/builders/select_builder"
	"wkey-stock/src/data/dtos"
	"wkey-stock/src/data/entities"
)

func (repo Repository) Create(ctx context.Context, category *dtos.Category) error {
	return repo.AddWithID(ctx, category.Entity())
}

func (repo Repository) All(ctx context.Context) ([]entities.Category, error) {
	return repo.GetList(ctx, func(builder *select_builder.Builder) {
		builder.Where(builder.Is("deleted", false))
	})
}

func (repo Repository) ByID(ctx context.Context, id string) (*entities.Category, error) {
	return repo.GetItem(ctx, func(builder *select_builder.Builder) {
		builder.Where(
			builder.Is("deleted", false),
			builder.Equal("id", id),
		)
	})
}

func (repo Repository) ByProsklad(ctx context.Context, proskladID int) (*entities.Category, error) {
	return repo.GetItem(ctx, func(builder *select_builder.Builder) {
		builder.Where(builder.Equal("prosklad_id", proskladID))
	})
}

func (repo Repository) UpdateCategory(ctx context.Context, category *dtos.Category) error {
	return repo.UpdateByID(ctx, category.ID().String(), category.Entity())
}

func (repo Repository) MakeDeleted(ctx context.Context, category *dtos.Category) error {
	entity := category.Entity()
	entity.Deleted = true
	return repo.UpdateByID(ctx, category.ID().String(), entity)
}

func (repo Repository) Remove(ctx context.Context, category *dtos.Category) error {
	return repo.DeleteByID(ctx, category.ID().String())
}

func (repo Repository) RemoveByProsklad(ctx context.Context, proskladID int) error {
	return repo.Delete(ctx, func(builder *delete_builder.Builder) {
		builder.Where(builder.Equal("prosklad_id", proskladID))
	})
}
