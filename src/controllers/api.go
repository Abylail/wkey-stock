package controllers

import (
	"github.com/lowl11/boost"
	"wkey-stock/src/controllers/category_controller"
	"wkey-stock/src/controllers/product_controller"
	"wkey-stock/src/gateways"
)

type Dependencies struct {
	Gateways *gateways.Gateways
}

func BindAPI(router boost.Router, deps Dependencies) {
	setProduct(router, product_controller.New(deps.Gateways))
	setCategory(router, category_controller.New(deps.Gateways))
}

func setProduct(router boost.Router, controller *product_controller.Controller) {
	group := router.Group("/api/v1/stock/product")

	group.GET("/get", controller.Get)
	group.GET("/get/:product-id", controller.GetByID)
	group.POST("/add", controller.Add)
	group.PUT("/update-prosklad/:prosklad-id", controller.UpdateProsklad)
	group.PUT("/update-description/:product-id", controller.UpdateDescription)
	group.PUT("/update-count/:product-id", controller.UpdateCount)
	group.DELETE("/delete/:product-id", controller.Delete)
}

func setCategory(router boost.Router, controller *category_controller.Controller) {
	group := router.Group("/api/v1/stock/category")

	group.GET("/get", controller.Get)
	group.GET("/get/:category-id", controller.GetByID)
	group.POST("/add", controller.Add)
	group.PUT("/update-prosklad/:category-id", controller.UpdateProsklad)
	group.PUT("/update-count/:category-id", controller.UpdateCount)
	group.DELETE("/delete/:category-id", controller.Delete)
}
