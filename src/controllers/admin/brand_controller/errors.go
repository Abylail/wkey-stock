package brand_controller

import (
	"github.com/lowl11/boost/pkg/errors"
	"net/http"
)

const (
	typeErrorBrandBind           = "BrandParamBind"
	typeErrorBrandParamRequired  = "BrandParamRequired"
	typeErrorBrandAlreadyExist   = "BrandAlreadyExist"
	typeErrorBrandGetByID        = "BrandGetByID"
	typeErrorBrandGetByTitle     = "BrandGetByTitle"
	typeErrorBrandGetList        = "BrandGetList"
	typeErrorBrandAdd            = "BrandAdd"
	typeErrorBrandUpdate         = "BrandUpdate"
	typeErrorBrandUpdateIcon     = "BrandUpdateIcon"
	typeErrorBrandUpdateFileIcon = "BrandUpdateFileIcon"
	typeErrorBrandDelete         = "BrandDelete"
	typeErrorBrandNotFound       = "BrandNotFound"
)

func ErrorBrandBind(err error) error {
	return errors.
		New("Bind brand object error").
		SetType(typeErrorBrandBind).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorBrandParamRequired(paramName string) error {
	return errors.
		New("Brand param required").
		SetType(typeErrorBrandParamRequired).
		SetHttpCode(http.StatusInternalServerError).
		AddContext("param_name", paramName)
}

func ErrorBrandAlreadyExist(title string) error {
	return errors.
		New("Brand already exist").
		SetType(typeErrorBrandAlreadyExist).
		SetHttpCode(http.StatusInternalServerError).
		AddContext("title", title)
}

func ErrorBrandGetByID(err error) error {
	return errors.
		New("Get brand by id error").
		SetType(typeErrorBrandGetByID).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorBrandGetByTitle(title string) error {
	return errors.
		New("Get brand by title error").
		SetType(typeErrorBrandGetByTitle).
		SetHttpCode(http.StatusInternalServerError).
		AddContext("title", title)
}

func ErrorBrandGetList(err error) error {
	return errors.
		New("Get brand list error").
		SetType(typeErrorBrandGetList).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorBrandAdd(err error) error {
	return errors.
		New("Add new brand error").
		SetType(typeErrorBrandAdd).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorBrandUpdate(err error) error {
	return errors.
		New("Update brand error").
		SetType(typeErrorBrandUpdate).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorBrandUpdateIcon(err error) error {
	return errors.
		New("Update brand icon error").
		SetType(typeErrorBrandUpdateIcon).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorBrandUpdateFileIcon(err error) error {
	return errors.
		New("Update brand file icon error").
		SetType(typeErrorBrandUpdateFileIcon).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorBrandDelete(err error) error {
	return errors.
		New("Delete brand error").
		SetType(typeErrorBrandDelete).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorBrandNotFound(brandID int) error {
	return errors.
		New("Brand not found").
		SetType(typeErrorBrandNotFound).
		SetHttpCode(http.StatusNotFound).
		AddContext("brand_id", brandID)
}
