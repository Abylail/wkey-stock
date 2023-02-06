package main

import (
	"github.com/lowl11/lazylog/layers"
	"wkey-stock/src/api"
	"wkey-stock/src/controllers"
	"wkey-stock/src/definition"
	"wkey-stock/src/events"
	"wkey-stock/src/repositories"
)

func main() {
	definition.Init()
	logger := definition.Logger

	// инициализация ивентов
	apiEvents, err := events.Get()
	if err != nil {
		logger.Fatal(err, "Initializing events error", "Application")
	}

	// инициализация репозиториев
	apiRepositories, err := repositories.Get(apiEvents)
	if err != nil {
		logger.Fatal(err, "Initializing repositories error", layers.Database)
	}

	// контроллеры
	apiControllers := controllers.Get(apiEvents, apiRepositories)

	// запуск сервера
	api.StartServer(apiControllers, apiEvents)
}
