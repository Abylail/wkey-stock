package postgres_helper

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/lazylog/layers"
	"time"
	"wkey-stock/src/definition"

	_ "github.com/lib/pq"
)

func NewConnection() (*sqlx.DB, error) {
	config := definition.Config.Database
	logger := definition.Logger

	// подключение к Postgres
	connection, err := sqlx.Open("postgres", config.Connection)
	if err != nil {
		return nil, err
	}

	connection.SetMaxOpenConns(maxConnections)
	connection.SetMaxIdleConns(maxConnections)
	connection.SetConnMaxIdleTime(time.Duration(maxLifetime) * time.Minute)

	logger.Info("Ping Postgres database...", layers.Database)
	if err = connection.Ping(); err != nil {
		return nil, err
	}
	logger.Info("Ping Postgres database done!", layers.Database)

	return connection, nil
}
