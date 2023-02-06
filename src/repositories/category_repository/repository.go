package category_repository

import (
	"github.com/jmoiron/sqlx"
	"wkey-stock/src/events"
	"wkey-stock/src/repositories/repository"
)

type Repository struct {
	repository.Base
	connection *sqlx.DB
}

func Create(connection *sqlx.DB, apiEvents *events.ApiEvents) *Repository {
	return &Repository{
		Base:       repository.CreateBase(apiEvents.Script),
		connection: connection,
	}
}
