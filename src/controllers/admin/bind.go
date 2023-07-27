package admin

import (
	"github.com/lowl11/boost"
	"wkey-stock/src/controllers/admin/brand_controller"
	"wkey-stock/src/controllers/admin/category_controller"
	"wkey-stock/src/controllers/admin/product_controller"
	"wkey-stock/src/controllers/admin/promotion_controller"
	"wkey-stock/src/events"
	"wkey-stock/src/repositories"
)

func Bind(router boost.Router, apiRepositories *repositories.ApiRepositories, apiEvents *events.ApiEvents) {
	setProduct(router, product_controller.New(apiRepositories, apiEvents))
	setCategory(router, category_controller.New(apiRepositories, apiEvents))
	setPromotion(router, promotion_controller.New(apiRepositories, apiEvents))
	setBrand(router, brand_controller.New(apiRepositories, apiEvents))
}

func setProduct(router boost.Router, controller *product_controller.Controller) {
	group := router.Group("/admin/api/v1/stock/product")

	group.GET("/get", controller.GetREST)
	group.GET("/get/:id", controller.GetSingleREST)
	group.PUT("/update/:id", controller.UpdateProductREST)
	group.PUT("/upload/:id", controller.UploadProductREST)
}

func setCategory(router boost.Router, controller *category_controller.Controller) {
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

func setPromotion(router boost.Router, controller *promotion_controller.Controller) {
	group := router.Group("/admin/api/v1/stock/promotion")

	group.GET("/get", controller.GetListREST)
	group.GET("/get/:code", controller.GetSingleREST)
	group.GET("/id/:id", controller.GetSingleREST)
	group.POST("/create", controller.CreateREST)
	group.PUT("/update", controller.UpdateREST)
	group.PUT("/upload", controller.UploadREST)
	group.DELETE("/delete/:code", controller.DeleteREST)
}

func setBrand(router boost.Router, controller *brand_controller.Controller) {
	group := router.Group("/admin/api/v1/stock/brand")

	group.GET("/get", controller.GetREST)
	group.GET("/get/:id", controller.GetSingleREST)
	group.POST("/add", controller.AddREST)
	group.PUT("/update/:id", controller.UpdateREST)
	group.PUT("/upload/:id", controller.UploadREST)
	group.DELETE("/delete/:id", controller.DeleteREST)
}
