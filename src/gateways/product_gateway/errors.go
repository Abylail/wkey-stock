package product_gateway

import (
	"github.com/lowl11/boost/pkg/errors"
	"net/http"
)

const (
	typErrorGetAllProducts = "GetAllProducts"
)

func ErrorGetAllProducts(err error) error {
	return errors.
		New("Get all products error").
		SetType(typErrorGetAllProducts).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}
