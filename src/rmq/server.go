package rmq

import (
	"wkey-stock/src/controllers"
	"wkey-stock/src/controllers/product_controller"
)

type Server struct {
	Product *product_controller.Controller
}

func New(apiControllers *controllers.ApiControllers) *Server {
	return &Server{
		Product: apiControllers.Product,
	}
}
