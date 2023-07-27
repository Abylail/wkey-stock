package client

import (
	"github.com/lowl11/boost"
	"wkey-stock/src/controllers/client/category_controller"
	"wkey-stock/src/controllers/client/product_controller"
	"wkey-stock/src/controllers/client/promotion_controller"
	"wkey-stock/src/events"
	"wkey-stock/src/repositories"
)

func Bind(router boost.Router, apiRepositories *repositories.ApiRepositories, apiEvents *events.ApiEvents) {
	setProduct(router, product_controller.New(apiRepositories, apiEvents))
	setCategory(router, category_controller.New(apiRepositories, apiEvents))
	setPromotion(router, promotion_controller.New(apiRepositories, apiEvents))
}

func setProduct(router boost.Router, controller *product_controller.Controller) {
	group := router.Group("/api/v1/stock/product")

	group.GET("/get", controller.GetREST)
}

func setCategory(router boost.Router, controller *category_controller.Controller) {
	// category
	categoryGroup := router.Group("/api/v1/stock/category")

	categoryGroup.GET("/get", controller.GetREST)
	categoryGroup.GET("/get/:code", controller.GetSingleREST)

	// sub category
	subCategoryGroup := router.Group("/api/v1/stock/category/:par_code/sub")

	subCategoryGroup.GET("/get", controller.GetSubREST)
	subCategoryGroup.GET("/get/:code", controller.GetSubSingleREST)
}

func setPromotion(router boost.Router, controller *promotion_controller.Controller) {
	group := router.Group("/api/v1/stock/promotion")

	group.GET("/get", controller.GetListREST)
	group.GET("/get/:code", controller.GetByCodeREST)
}
