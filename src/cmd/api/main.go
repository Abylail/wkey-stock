package main

import (
	"github.com/lowl11/boost"
	"github.com/lowl11/boost/pkg/middlewares"
	"github.com/lowl11/lazyconfig/config"
	"github.com/lowl11/lazylog/log"
	"time"
	"wkey-stock/src/controllers"
	"wkey-stock/src/gateways"
	"wkey-stock/src/repositories"
)

func main() {
	app := boost.New()

	app.Use(middlewares.Timeout(time.Second * 10))

	// инициализация репозиториев
	apiRepositories, err := repositories.Get(app)
	if err != nil {
		log.Fatal(err, "Initializing repositories error")
	}

	// привязка роутов
	controllers.BindAPI(app, controllers.Dependencies{
		Gateways: gateways.Get(apiRepositories),
	})

	// запуск сервера
	app.Run(config.Get("port"))
}
