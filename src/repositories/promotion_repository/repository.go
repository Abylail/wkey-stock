package promotion_repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/lazy-entity/repo_config"
	"github.com/lowl11/lazy-entity/repository"
	"wkey-stock/src/data/entities"
)

type Repository struct {
	repository.ICrudRepository[entities.Promotion, string]
}

func New(connection *sqlx.DB) *Repository {
	return &Repository{
		ICrudRepository: repository.NewCrud[entities.Promotion, string](
			connection,
			"promotions",
			repo_config.Crud{
				AliasName: "prom",
			},
		),
	}
}
