package main

import (
	"github.com/lowl11/boost"
	"github.com/lowl11/boost/pkg/middlewares"
	"github.com/lowl11/lazyconfig/config"
	"github.com/lowl11/lazylog/log"
	"time"

	"wkey-stock/src/controllers"
	"wkey-stock/src/events"
	"wkey-stock/src/repositories"
)

func main() {
	app := boost.New()

	app.Use(middlewares.Timeout(time.Second * 10))

	// инициализация ивентов
	apiEvents, err := events.Get()
	if err != nil {
		log.Fatal(err, "Initializing events error")
	}

	// инициализация репозиториев
	apiRepositories, err := repositories.Get(app)
	if err != nil {
		log.Fatal(err, "Initializing repositories error")
	}

	controllers.Bind(app, apiEvents, apiRepositories)

	// запуск сервера
	app.Run(config.Get("port"))
}
