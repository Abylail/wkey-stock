package repositories

import (
	"github.com/lowl11/boost"
	"github.com/lowl11/lazy-entity/services/connection_service"
	"github.com/lowl11/lazyconfig/config"
	"github.com/lowl11/lazylog/log"
	"wkey-stock/src/repositories/category_repository"
	"wkey-stock/src/repositories/product_repository"
)

type ApiRepositories struct {
	Product  *product_repository.Repository
	Category *category_repository.Repository
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
		Product:  product_repository.New(connection),
		Category: category_repository.New(connection),
	}, nil
}
