package api

import (
	"wkey-stock/src/controllers"
	"wkey-stock/src/definition"
	"wkey-stock/src/events"
)

func StartServer(apiControllers *controllers.ApiControllers, apiEvents *events.ApiEvents) {
	server := definition.Server
	config := definition.Config.Server

	// проставлять роуты
	setRoutes(server, apiControllers, apiEvents)

	// проставлять миддлвейры
	setMiddlewares(server)

	// запуск сервера
	server.Logger.Fatal(server.Start(config.Port))
}
