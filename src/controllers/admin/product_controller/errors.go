package product_controller

import (
	"github.com/lowl11/boost/pkg/errors"
	"net/http"
)

const (
	typeErrorProductParamRequired    = "ProductParamRequired"
	typeErrorProductGetImages        = "ProductGetImages"
	typeErrorProductUpdate           = "ProductUpdate"
	typeErrorProductUpdateImages     = "ProductUpdateImages"
	typeErrorProductUpdateFileImages = "ProductUpdateFileImages"
	typeErrorProductGetPairs         = "ProductGetPairs"
	typeErrorProductNotFound         = "ProductNotFound"
	typeErrorProductGet              = "ProductGet"
	typeErrorProductGetCount         = "ProductGetCount"
)

func ErrorProductParamRequired(paramName string) error {
	return errors.
		New("AdminProduct param required").
		SetType(typeErrorProductParamRequired).
		SetHttpCode(http.StatusInternalServerError).
		AddContext("param_name", paramName)
}

func ErrorProductGetImages(err error) error {
	return errors.
		New("Get product images error").
		SetType(typeErrorProductGetImages).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorProductUpdate(err error) error {
	return errors.
		New("Update product error").
		SetType(typeErrorProductUpdate).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorProductGetPairs(err error) error {
	return errors.
		New("Get product pairs error").
		SetType(typeErrorProductGetPairs).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorAdminProductNotFound(productID int) error {
	return errors.
		New("Admin product not found").
		SetType(typeErrorProductNotFound).
		SetHttpCode(http.StatusInternalServerError).
		AddContext("product_id", productID)
}

func ErrorAdminProductGet(err error) error {
	return errors.
		New("Get admin product error").
		SetType(typeErrorProductGet).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorAdminProductGetCount(err error) error {
	return errors.
		New("Get admin product count error").
		SetType(typeErrorProductGetCount).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorProductUpdateImages(err error) error {
	return errors.
		New("Update product images error").
		SetType(typeErrorProductUpdateImages).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorProductUpdateFileImages(err error) error {
	return errors.
		New("Update product file images error").
		SetType(typeErrorProductUpdateFileImages).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}
