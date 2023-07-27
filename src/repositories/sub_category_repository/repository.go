package sub_category_repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/lazy-entity/repo_config"
	"github.com/lowl11/lazy-entity/repository"
	"wkey-stock/src/data/entities"
)

type Repository struct {
	repository.ICrudRepository[entities.SubCategory, int]
}

func New(connection *sqlx.DB) *Repository {
	return &Repository{
		ICrudRepository: repository.NewCrud[entities.SubCategory, int](
			connection,
			"sub_categories",
			repo_config.Crud{},
		),
	}
}
