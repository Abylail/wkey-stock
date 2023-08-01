package controllers

import (
	"github.com/lowl11/boost"
	"wkey-stock/src/controllers/product_controller"
	"wkey-stock/src/gateways"
)

type Dependencies struct {
	Gateways *gateways.Gateways
}

func BindAPI(router boost.Router, deps Dependencies) {
	setProduct(router, product_controller.New(deps.Gateways))
}

func setProduct(router boost.Router, controller *product_controller.Controller) {
	group := router.Group("/api/v1/stock/product")

	group.GET("/get", controller.Get)
	group.GET("/get/:product-id", controller.GetByID)
	group.POST("/add", controller.Add)
	group.PUT("/update-description/:product-id", controller.UpdateDescription)
	group.PUT("/update-count/:product-id", controller.UpdateCount)
	group.DELETE("/delete/:product-id", controller.Delete)
}
