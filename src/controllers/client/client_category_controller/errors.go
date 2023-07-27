package client_category_controller

import (
	"github.com/lowl11/boost/pkg/errors"
	"net/http"
)

const (
	// category
	typeErrorCategoryNotFound  = "CategoryNotFound"
	typeErrorCategoryGetList   = "CategoryGetList"
	typeErrorCategoryGetByCode = "CategoryGetByCode"

	// sub category
	typeErrorSubCategoryNotFound  = "SubCategoryNotFound"
	typeErrorSubCategoryGetByCode = "SubCategoryGetByCode"
)

func ErrorCategoryGetList(err error) error {
	return errors.
		New("Get category list error").
		SetType(typeErrorCategoryGetList).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorCategoryGetByCode(code string) error {
	return errors.
		New("Get category by code error").
		SetType(typeErrorCategoryGetByCode).
		SetHttpCode(http.StatusInternalServerError).
		AddContext("code", code)
}

func ErrorCategoryNotFound(code string) error {
	return errors.
		New("AdminCategory not found").
		SetType(typeErrorCategoryNotFound).
		SetHttpCode(http.StatusNotAcceptable).
		AddContext("code", code)
}

func ErrorSubCategoryGetByCode(err error) error {
	return errors.
		New("Get sub-category by code error").
		SetType(typeErrorSubCategoryGetByCode).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorSubCategoryNotFound(code string) error {
	return errors.
		New("Sub-category not found").
		SetType(typeErrorSubCategoryNotFound).
		SetHttpCode(http.StatusNotAcceptable).
		AddContext("code", code)
}
