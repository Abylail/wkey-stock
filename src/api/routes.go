package api

import (
	"github.com/labstack/echo/v4"
	"wkey-stock/src/controllers"
	"wkey-stock/src/controllers/category_controller"
	"wkey-stock/src/controllers/product_controller"
	"wkey-stock/src/events"
)

func setRoutes(server *echo.Echo, apiControllers *controllers.ApiControllers, _ *events.ApiEvents) {
	// статичные методы
	server.GET("/health", apiControllers.Static.Health)
	server.RouteNotFound("*", apiControllers.Static.RouteNotFound)

	// эндпоинты
	setProduct(server, apiControllers.Product)
	setCategory(server, apiControllers.Category)
}

func setProduct(server *echo.Echo, controller *product_controller.Controller) {
	group := server.Group("/api/v1/stock/product")

	group.GET("/get", controller.Get)
}

func setCategory(server *echo.Echo, controller *category_controller.Controller) {
	clientGroup := server.Group("/api/v1/stock/category")

	clientGroup.GET("/get", controller.GetClientREST)

	adminGroup := server.Group("/admin/api/v1/stock/category")

	adminGroup.GET("/get", controller.GetAdminREST)
}
