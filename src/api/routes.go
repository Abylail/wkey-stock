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
	setSubCategory(server, apiControllers.Category)
	setBrand(server, apiControllers.Product)
}

func setProduct(server *echo.Echo, controller *product_controller.Controller) {
	adminGroup := server.Group("/admin/api/v1/stock/product")

	adminGroup.GET("/get", controller.GetAdminREST)
	adminGroup.GET("/get/:id", controller.GetAdminSingleREST)
	adminGroup.PUT("/update/:id", controller.UpdateProductREST)

	clientGroup := server.Group("/api/v1/stock/product")

	clientGroup.GET("/get", controller.GetClientREST)
}

func setCategory(server *echo.Echo, controller *category_controller.Controller) {
	adminGroup := server.Group("/admin/api/v1/stock/category")

	adminGroup.GET("/get", controller.GetAdminREST)
	adminGroup.GET("/get/:code", controller.GetAdminSingleREST)
	adminGroup.POST("/add", controller.AddREST)
	adminGroup.PUT("/update/:code", controller.UpdateREST)
	adminGroup.PUT("/upload/:code", controller.UploadREST)
	adminGroup.DELETE("/delete/:code", controller.DeleteREST)

	clientGroup := server.Group("/api/v1/stock/category")

	clientGroup.GET("/get", controller.GetClientREST)
}

func setSubCategory(server *echo.Echo, controller *category_controller.Controller) {
	adminGroup := server.Group("/admin/api/v1/stock/category/:parent_code/sub")

	adminGroup.GET("/get", controller.GetAdminSubREST)
	adminGroup.POST("/add", controller.AddSubREST)
	adminGroup.PUT("/update/:code", controller.UpdateSubREST)
	adminGroup.PUT("/upload/:code", controller.UploadSubREST)
	adminGroup.DELETE("/delete/:code", controller.DeleteSubREST)

	clientGroup := server.Group("/api/v1/stock/category/:par_code/sub")

	clientGroup.GET("/get", controller.GetClientSubREST)
}

func setBrand(server *echo.Echo, controller *product_controller.Controller) {
	adminGroup := server.Group("/admin/api/v1/stock/brand")

	adminGroup.GET("/get", controller.GetBrandREST)
	adminGroup.GET("/get/:id", controller.GetBrandSingleREST)
	adminGroup.POST("/add", controller.AddBrandREST)
	adminGroup.PUT("/update/:id", controller.UpdateBrandREST)
	adminGroup.PUT("/upload/:id", controller.UploadBrandREST)
	adminGroup.DELETE("/delete/:id", controller.DeleteBrandREST)
}
