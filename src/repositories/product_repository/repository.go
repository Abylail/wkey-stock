package product_repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/lazy-entity/repo_config"
	"github.com/lowl11/lazy-entity/repository"
	"wkey-stock/src/data/entities"
)

type Repository struct {
	repository.ICrudRepository[entities.Product, string]
}

func New(connection *sqlx.DB) *Repository {
	return &Repository{
		ICrudRepository: repository.NewCrud[entities.Product, string](
			connection,
			"products",
			repo_config.Crud{},
		),
	}
}
