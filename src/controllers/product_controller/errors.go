package product_controller

import (
	"github.com/lowl11/boost/pkg/errors"
	"net/http"
)

const (
	// product types
	typeErrorProductBind             = "ProductParamBind"
	typeErrorProductParamRequired    = "ProductParamRequired"
	typeErrorProductGetImages        = "ProductGetImages"
	typeErrorProductUpdate           = "ProductUpdate"
	typeErrorProductUpdateImages     = "ProductUpdateImages"
	typeErrorProductUpdateFileImages = "ProductUpdateFileImages"
	typeErrorProductGetPairs         = "ProductGetPairs"

	// client products
	typeErrorClientProductGet      = "ClientProductGet"
	typeErrorClientProductGetCount = "ClientProductGetCount"

	// admin products
	typeErrorAdminProductNotFound = "AdminProductNotFound"
	typeErrorAdminProductGet      = "AdminProductGet"
	typeErrorAdminProductGetCount = "AdminProductGetCount"

	// brand types
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

func ErrorProductBind() error {
	return errors.
		New("Bind product object error").
		SetType(typeErrorProductBind).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorProductParamRequired() error {
	return errors.
		New("Product param required").
		SetType(typeErrorProductParamRequired).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorProductGetImages() error {
	return errors.
		New("Get product images error").
		SetType(typeErrorProductGetImages).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorProductUpdate() error {
	return errors.
		New("Update product error").
		SetType(typeErrorProductUpdate).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorProductGetPairs() error {
	return errors.
		New("Get product pairs error").
		SetType(typeErrorProductGetPairs).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorClientProductGet() error {
	return errors.
		New("Get client product error").
		SetType(typeErrorClientProductGet).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorClientProductGetCount() error {
	return errors.
		New("Get client product count error").
		SetType(typeErrorClientProductGetCount).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorAdminProductNotFound() error {
	return errors.
		New("Admin product not found").
		SetType(typeErrorAdminProductNotFound).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorAdminProductGet() error {
	return errors.
		New("Get admin product error").
		SetType(typeErrorAdminProductGet).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorAdminProductGetCount() error {
	return errors.
		New("Get admin product count error").
		SetType(typeErrorAdminProductGetCount).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorProductUpdateImages() error {
	return errors.
		New("Update product images error").
		SetType(typeErrorProductUpdateImages).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorProductUpdateFileImages() error {
	return errors.
		New("Update product file images error").
		SetType(typeErrorProductUpdateFileImages).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorBrandBind() error {
	return errors.
		New("Bind brand object error").
		SetType(typeErrorBrandBind).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorBrandParamRequired() error {
	return errors.
		New("Brand param required").
		SetType(typeErrorBrandParamRequired).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorBrandAlreadyExist() error {
	return errors.
		New("Brand already exist").
		SetType(typeErrorBrandAlreadyExist).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorBrandGetByID() error {
	return errors.
		New("Get brand by id error").
		SetType(typeErrorBrandGetByID).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorBrandGetByTitle() error {
	return errors.New("Get brand by title error").
		SetType(typeErrorBrandGetByTitle).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorBrandGetList() error {
	return errors.
		New("Get brand list error").
		SetType(typeErrorBrandGetList).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorBrandAdd() error {
	return errors.
		New("Add new brand error").
		SetType(typeErrorBrandAdd).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorBrandUpdate() error {
	return errors.
		New("Update brand error").
		SetType(typeErrorBrandUpdate).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorBrandUpdateIcon() error {
	return errors.
		New("Update brand icon error").
		SetType(typeErrorBrandUpdateIcon).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorBrandUpdateFileIcon() error {
	return errors.
		New("Update brand file icon error").
		SetType(typeErrorBrandUpdateFileIcon).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorBrandDelete() error {
	return errors.
		New("Delete brand error").
		SetType(typeErrorBrandDelete).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorBrandNotFound() error {
	return errors.
		New("Brand not found").
		SetType(typeErrorBrandNotFound).
		SetHttpCode(http.StatusNotFound)
}
