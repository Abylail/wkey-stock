package repositories

import (
	"github.com/lowl11/lazylog/layers"
	"wkey-stock/src/definition"
	"wkey-stock/src/events"
	"wkey-stock/src/repositories/brand_repository"
	"wkey-stock/src/repositories/category_repository"
	"wkey-stock/src/repositories/product_repository"
	"wkey-stock/src/repositories/sub_category_repository"
	"wkey-stock/src/services/postgres_helper"
)

type ApiRepositories struct {
	Category    *category_repository.Repository
	SubCategory *sub_category_repository.Repository
	Product     *product_repository.Repository
	Brand       *brand_repository.Repository
}

func Get(apiEvents *events.ApiEvents) (*ApiRepositories, error) {
	logger := definition.Logger

	// подключение к Postgres
	connectionPostgres, err := postgres_helper.NewConnection()
	if err != nil {
		logger.Fatal(err, "Connect to Postgres database error", layers.Database)
	}

	return &ApiRepositories{
		Category:    category_repository.Create(connectionPostgres, apiEvents),
		SubCategory: sub_category_repository.Create(connectionPostgres, apiEvents),
		Product:     product_repository.Create(connectionPostgres, apiEvents),
		Brand:       brand_repository.Create(connectionPostgres, apiEvents),
	}, nil
}
