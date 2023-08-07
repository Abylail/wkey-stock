package promotion_repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	connection *sqlx.DB
}

func New(connection *sqlx.DB) *Repository {
	return &Repository{
		connection: connection,
	}
}
