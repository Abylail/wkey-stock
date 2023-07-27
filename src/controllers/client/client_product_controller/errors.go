package client_product_controller

import (
	"github.com/lowl11/boost/pkg/errors"
	"net/http"
)

const (
	typeErrorProductGetImages = "ProductGetImages"
	typeErrorProductGet       = "ProductGet"
	typeErrorProductGetCount  = "ProductGetCount"
)

func ErrorProductGet(err error) error {
	return errors.
		New("Get client product error").
		SetType(typeErrorProductGet).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorProductGetCount(err error) error {
	return errors.
		New("Get client product count error").
		SetType(typeErrorProductGetCount).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorProductGetImages(err error) error {
	return errors.
		New("Get product images error").
		SetType(typeErrorProductGetImages).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}
