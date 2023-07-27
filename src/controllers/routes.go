package controllers

import (
	"github.com/lowl11/boost"
	"wkey-stock/src/controllers/category_controller"
	"wkey-stock/src/controllers/product_controller"
	"wkey-stock/src/controllers/promotion_controller"
	"wkey-stock/src/events"
	"wkey-stock/src/repositories"
)

func Bind(app *boost.App, apiEvents *events.ApiEvents, apiRepositories *repositories.ApiRepositories) {
	setRoutes(app, Get(apiEvents, apiRepositories))
}

func setRoutes(router boost.Router, apiControllers *ApiControllers) {
	setProduct(router, apiControllers.Product)
	setCategory(router, apiControllers.Category)
	setSubCategory(router, apiControllers.Category)
	setBrand(router, apiControllers.Product)
	setPromotion(router, apiControllers.Promotion)
}

func setProduct(router boost.Router, controller *product_controller.Controller) {
	adminGroup := router.Group("/admin/api/v1/stock/product")

	adminGroup.GET("/get", controller.GetAdminREST)
	adminGroup.GET("/get/:id", controller.GetAdminSingleREST)
	adminGroup.PUT("/update/:id", controller.UpdateProductREST)
	adminGroup.PUT("/upload/:id", controller.UploadProductREST)

	clientGroup := router.Group("/api/v1/stock/product")
	clientGroup.GET("/get", controller.GetClientREST)
}

func setCategory(router boost.Router, controller *category_controller.Controller) {
	adminGroup := router.Group("/admin/api/v1/stock/category")

	adminGroup.GET("/get", controller.GetAdminREST)
	adminGroup.GET("/get/:code", controller.GetAdminSingleREST)
	adminGroup.POST("/add", controller.AddREST)
	adminGroup.PUT("/update/:code", controller.UpdateREST)
	adminGroup.PUT("/upload/:code", controller.UploadREST)
	adminGroup.DELETE("/delete/:code", controller.DeleteREST)

	adminGroup.POST("/activate/:code", controller.ActivateREST)
	adminGroup.POST("/deactivate/:code", controller.DeactivateREST)

	clientGroup := router.Group("/api/v1/stock/category")

	clientGroup.GET("/get", controller.GetClientREST)
	clientGroup.GET("/get/:code", controller.GetClientSingleREST)
}

func setSubCategory(router boost.Router, controller *category_controller.Controller) {
	adminGroup := router.Group("/admin/api/v1/stock/category/:parent_code/sub")

	adminGroup.GET("/get", controller.GetAdminSubREST)
	adminGroup.GET("/get/:code", controller.GetAdminSingleSubREST)
	adminGroup.POST("/add", controller.AddSubREST)
	adminGroup.PUT("/update/:code", controller.UpdateSubREST)
	adminGroup.PUT("/upload/:code", controller.UploadSubREST)
	adminGroup.DELETE("/delete/:code", controller.DeleteSubREST)

	adminGroup.POST("/activate/:code", controller.ActivateSubREST)
	adminGroup.POST("/deactivate/:code", controller.DeactivateSubREST)

	adminGroup.POST("/bind/:code", controller.BindProductListREST)
	adminGroup.POST("/unbind/:code/product/:product_id", controller.UnbindProductItemREST)

	clientGroup := router.Group("/api/v1/stock/category/:par_code/sub")

	clientGroup.GET("/get", controller.GetClientSubREST)
	clientGroup.GET("/get/:code", controller.GetClientSubSingleREST)
}

func setBrand(router boost.Router, controller *product_controller.Controller) {
	adminGroup := router.Group("/admin/api/v1/stock/brand")

	adminGroup.GET("/get", controller.GetBrandREST)
	adminGroup.GET("/get/:id", controller.GetBrandSingleREST)
	adminGroup.POST("/add", controller.AddBrandREST)
	adminGroup.PUT("/update/:id", controller.UpdateBrandREST)
	adminGroup.PUT("/upload/:id", controller.UploadBrandREST)
	adminGroup.DELETE("/delete/:id", controller.DeleteBrandREST)
}

// Промоакции
func setPromotion(router boost.Router, controller *promotion_controller.Controller) {
	adminGroup := router.Group("/admin/api/v1/stock/promotion")

	adminGroup.GET("/get", controller.GetListAdmin)
	adminGroup.GET("/get/:code", controller.GetSingleCodeAdmin)
	adminGroup.GET("/id/:id", controller.GetSingleAdmin)
	adminGroup.POST("/create", controller.CreateAdmin)
	adminGroup.PUT("/update", controller.UpdateAdmin)
	adminGroup.PUT("/upload", controller.UploadAdmin)
	adminGroup.DELETE("/delete/:code", controller.DeleteAdmin)

	clientGroup := router.Group("/api/v1/stock/promotion")
	clientGroup.GET("/get", controller.GetListClient)
	clientGroup.GET("/get/:code", controller.GetSingleClient)
}
