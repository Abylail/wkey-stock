package repositories

import (
	"github.com/lowl11/boost"
	"github.com/lowl11/lazy-entity/services/connection_service"
	"github.com/lowl11/lazyconfig/config"
	"github.com/lowl11/lazylog/log"
	"wkey-stock/src/repositories/brand_repository"
	"wkey-stock/src/repositories/category_repository"
	"wkey-stock/src/repositories/product_repository"
	"wkey-stock/src/repositories/promotion_repository"
	"wkey-stock/src/repositories/sub_category_repository"
)

type ApiRepositories struct {
	Category    *category_repository.Repository
	SubCategory *sub_category_repository.Repository
	Product     *product_repository.Repository
	Brand       *brand_repository.Repository
	Promotion   *promotion_repository.Repository
}

func Get(app *boost.App) (*ApiRepositories, error) {
	// подключение к Postgres
	connection, err := connection_service.
		New(config.Get("database_connection")).
		ConnectionPool()
	if err != nil {
		return nil, err
	}

	app.Destroy(func() {
		if err = connection.Close(); err != nil {
			log.Error(err, "Database connection close error")
			return
		}
		log.Info("Database connection closed")
	})

	return &ApiRepositories{
		Category:    category_repository.New(connection),
		SubCategory: sub_category_repository.New(connection),
		Product:     product_repository.New(connection),
		Brand:       brand_repository.New(connection),
		Promotion:   promotion_repository.New(connection),
	}, nil
}
