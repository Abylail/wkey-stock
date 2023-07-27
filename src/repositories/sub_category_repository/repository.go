package sub_category_repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/lazy-entity/repository"
)

type Repository struct {
	repository.IScriptRepository
	connection *sqlx.DB
}

func New(connection *sqlx.DB) *Repository {
	return &Repository{
		IScriptRepository: repository.NewScript(),
		connection:        connection,
	}
}
