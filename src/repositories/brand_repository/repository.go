package brand_repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/lazy-entity/repo_config"
	"github.com/lowl11/lazy-entity/repository"
	"wkey-stock/src/data/entities"
)

type Repository struct {
	repository.ICrudRepository[entities.Brand, string]
}

func New(connection *sqlx.DB) *Repository {
	return &Repository{
		ICrudRepository: repository.NewCrud[entities.Brand, string](
			connection,
			"brands",
			repo_config.Crud{
				AliasName: "brand",
			},
		),
	}
}
