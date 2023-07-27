package controllers

import (
	"github.com/lowl11/boost"
	"wkey-stock/src/controllers/admin/admin_category_controller"
	"wkey-stock/src/controllers/admin/admin_product_controller"
	"wkey-stock/src/controllers/admin/admin_promotion_controller"
	"wkey-stock/src/controllers/client/client_category_controller"
	"wkey-stock/src/controllers/client/client_product_controller"
	"wkey-stock/src/controllers/client/client_promotion_controller"
	"wkey-stock/src/events"
	"wkey-stock/src/repositories"
)

func Bind(app *boost.App, apiEvents *events.ApiEvents, apiRepositories *repositories.ApiRepositories) {
	setRoutes(app, Get(apiEvents, apiRepositories))
}

func setRoutes(router boost.Router, apiControllers *ApiControllers) {
	setBrand(router, apiControllers.AdminProduct)

	setAdminProduct(router, apiControllers.AdminProduct)
	setClientProduct(router, apiControllers.ClientProduct)

	setAdminCategory(router, apiControllers.AdminCategory)
	setClientCategory(router, apiControllers.ClientCategory)

	setPromotionAdmin(router, apiControllers.AdminPromotion)
	setPromotionClient(router, apiControllers.ClientPromotion)
}

func setAdminProduct(router boost.Router, controller *admin_product_controller.Controller) {
	group := router.Group("/admin/api/v1/stock/product")

	group.GET("/get", controller.GetREST)
	group.GET("/get/:id", controller.GetSingleREST)
	group.PUT("/update/:id", controller.UpdateProductREST)
	group.PUT("/upload/:id", controller.UploadProductREST)
}

func setClientProduct(router boost.Router, controller *client_product_controller.Controller) {
	group := router.Group("/api/v1/stock/product")
	group.GET("/get", controller.GetREST)
}

func setAdminCategory(router boost.Router, controller *admin_category_controller.Controller) {
	// category
	categoryGroup := router.Group("/admin/api/v1/stock/category")

	categoryGroup.GET("/get", controller.GetREST)
	categoryGroup.GET("/get/:code", controller.GetSingleREST)
	categoryGroup.POST("/add", controller.AddREST)
	categoryGroup.PUT("/update/:code", controller.UpdateREST)
	categoryGroup.PUT("/upload/:code", controller.UploadREST)
	categoryGroup.DELETE("/delete/:code", controller.DeleteREST)

	categoryGroup.POST("/activate/:code", controller.ActivateREST)
	categoryGroup.POST("/deactivate/:code", controller.DeactivateREST)

	// sub category
	subCategoryGroup := router.Group("/admin/api/v1/stock/category/:parent_code/sub")

	subCategoryGroup.GET("/get", controller.GetSubREST)
	subCategoryGroup.GET("/get/:code", controller.GetSingleSubREST)
	subCategoryGroup.POST("/add", controller.AddSubREST)
	subCategoryGroup.PUT("/update/:code", controller.UpdateSubREST)
	subCategoryGroup.PUT("/upload/:code", controller.UploadSubREST)
	subCategoryGroup.DELETE("/delete/:code", controller.DeleteSubREST)

	subCategoryGroup.POST("/activate/:code", controller.ActivateSubREST)
	subCategoryGroup.POST("/deactivate/:code", controller.DeactivateSubREST)

	subCategoryGroup.POST("/bind/:code", controller.BindProductListREST)
	subCategoryGroup.POST("/unbind/:code/product/:product_id", controller.UnbindProductItemREST)
}

func setClientCategory(router boost.Router, controller *client_category_controller.Controller) {
	// category
	categoryGroup := router.Group("/api/v1/stock/category")

	categoryGroup.GET("/get", controller.GetREST)
	categoryGroup.GET("/get/:code", controller.GetSingleREST)

	// sub category
	subCategoryGroup := router.Group("/api/v1/stock/category/:par_code/sub")

	subCategoryGroup.GET("/get", controller.GetSubREST)
	subCategoryGroup.GET("/get/:code", controller.GetSubSingleREST)
}

func setBrand(router boost.Router, controller *admin_product_controller.Controller) {
	group := router.Group("/admin/api/v1/stock/brand")

	group.GET("/get", controller.GetBrandREST)
	group.GET("/get/:id", controller.GetBrandSingleREST)
	group.POST("/add", controller.AddBrandREST)
	group.PUT("/update/:id", controller.UpdateBrandREST)
	group.PUT("/upload/:id", controller.UploadBrandREST)
	group.DELETE("/delete/:id", controller.DeleteBrandREST)
}

func setPromotionAdmin(router boost.Router, controller *admin_promotion_controller.Controller) {
	group := router.Group("/admin/api/v1/stock/promotion")

	group.GET("/get", controller.GetListREST)
	group.GET("/get/:code", controller.GetSingleREST)
	group.GET("/id/:id", controller.GetSingleREST)
	group.POST("/create", controller.CreateREST)
	group.PUT("/update", controller.UpdateREST)
	group.PUT("/upload", controller.UploadREST)
	group.DELETE("/delete/:code", controller.DeleteREST)
}

func setPromotionClient(router boost.Router, controller *client_promotion_controller.Controller) {
	group := router.Group("/api/v1/stock/promotion")
	group.GET("/get", controller.GetListREST)
	group.GET("/get/:code", controller.GetByCodeREST)
}
