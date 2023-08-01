package category_controller

import (
	"github.com/lowl11/boost/pkg/errors"
	"net/http"
)

const (
	typeErrorCategoryIDRequired = "CategoryIDRequired"
	typeErrorProskladIDRequired = "CategoryIDRequired"
)

func ErrorCategoryIDRequired() error {
	return errors.
		New("Category ID is required").
		SetType(typeErrorCategoryIDRequired).
		SetHttpCode(http.StatusUnprocessableEntity)
}

func ErrorProskladIDRequired() error {
	return errors.
		New("Prosklad ID is required").
		SetType(typeErrorProskladIDRequired).
		SetHttpCode(http.StatusUnprocessableEntity)
}
