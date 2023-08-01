package product_controller

import (
	"github.com/lowl11/boost/pkg/errors"
	"net/http"
)

const (
	typeErrorProductIDRequired  = "ProductIDRequired"
	typeErrorProskladIDRequired = "ProskladIDRequired"
)

func ErrorProductIDRequired() error {
	return errors.
		New("Product ID is required").
		SetType(typeErrorProductIDRequired).
		SetHttpCode(http.StatusUnprocessableEntity)
}

func ErrorProskladIDRequired() error {
	return errors.
		New("Prosklad ID is required").
		SetType(typeErrorProskladIDRequired).
		SetHttpCode(http.StatusUnprocessableEntity)
}
