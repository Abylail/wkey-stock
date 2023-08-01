package product_controller

import (
	"github.com/lowl11/boost/pkg/errors"
	"net/http"
)

const (
	typeErrorProductIDRequired = "ProductIDRequired"
)

func ErrorProductIDRequired() error {
	return errors.
		New("Product ID is required").
		SetType(typeErrorProductIDRequired).
		SetHttpCode(http.StatusUnprocessableEntity)
}
