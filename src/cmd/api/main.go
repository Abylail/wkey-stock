package main

import (
	"github.com/lowl11/boost"
	"github.com/lowl11/boost/pkg/middlewares"
	"github.com/lowl11/lazyconfig/config"
	"github.com/lowl11/lazylog/log"
	"time"
	"wkey-stock/src/controllers/admin"
	"wkey-stock/src/controllers/client"

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

	// привязка роутов
	admin.Bind(app, apiRepositories, apiEvents)
	client.Bind(app, apiRepositories, apiEvents)

	// запуск сервера
	app.Run(config.Get("port"))
}
