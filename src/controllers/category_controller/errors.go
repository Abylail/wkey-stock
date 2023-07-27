package category_controller

import (
	"github.com/lowl11/boost/pkg/errors"
	"net/http"
)

const (
	// category
	typeErrorCategoryBind            = "CategoryBind"
	typeErrorCategoryParamRequired   = "CategoryParamRequired"
	typeErrorCategoryAlreadyExist    = "CategoryAlreadyExist"
	typeErrorCategoryNotFound        = "CategoryNotFound"
	typeErrorCategoryHasChildren     = "CategoryHasChildren"
	typeErrorCategoryNoChildren      = "CategoryNoChildren"
	typeErrorCategoryGetList         = "CategoryGetList"
	typeErrorCategoryGetByCode       = "CategoryGetByCode"
	typeErrorCategoryGetCount        = "CategoryGetCount"
	typeErrorCategoryAdd             = "CategoryAdd"
	typeErrorCategoryUpdate          = "CategoryUpdate"
	typeErrorCategoryUpdateStatus    = "CategoryUpdateStatus"
	typeErrorCategoryUpdateImage     = "CategoryUpdateImage"
	typeErrorCategoryUpdateFileImage = "CategoryUpdateFileImage"
	typeErrorCategoryBindSub         = "CategoryBindSub"
	typeErrorCategoryUnbindSub       = "CategoryUnbindSub"
	typeErrorCategoryDelete          = "CategoryDelete"

	// sub category
	typeErrorSubCategoryBind            = "SubCategoryBind"
	typeErrorSubCategoryAlreadyExist    = "SubCategoryAlreadyExist"
	typeErrorSubCategoryNotFound        = "SubCategoryNotFound"
	typeErrorSubCategoryGetList         = "SubCategoryGetList"
	typeErrorSubCategoryGetByCode       = "SubCategoryGetByCode"
	typeErrorSubCategoryGetCount        = "SubCategoryGetCount"
	typeErrorSubCategoryAdd             = "SubCategoryAdd"
	typeErrorSubCategoryUpdate          = "SubCategoryUpdate"
	typeErrorSubCategoryUpdateStatus    = "SubCategoryUpdateStatus"
	typeErrorSubCategoryUpdateImage     = "SubCategoryUpdateImage"
	typeErrorSubCategoryUpdateFileImage = "SubCategoryUpdateFileImage"
	typeErrorSubCategoryDelete          = "SubCategoryDelete"
)

func ErrorCategoryBind() error {
	return errors.New("Bind category error").
		SetType(typeErrorCategoryBind).
		SetHttpCode(http.StatusUnprocessableEntity)
}

func ErrorCategoryParamRequired() error {
	return errors.New("Category param required error").
		SetType(typeErrorCategoryParamRequired).
		SetHttpCode(http.StatusUnprocessableEntity)
}

func ErrorSubCategoryBind() error {
	return errors.New("Bind sub-category error").
		SetType(typeErrorSubCategoryBind).
		SetHttpCode(http.StatusUnprocessableEntity)
}

func ErrorCategoryAlreadyExist() error {
	return errors.
		New("Category already exist").
		SetType(typeErrorCategoryAlreadyExist).
		SetHttpCode(http.StatusNotAcceptable)
}

func ErrorCategoryNotFound() error {
	return errors.
		New("Category not found").
		SetType(typeErrorCategoryNotFound).
		SetHttpCode(http.StatusNotAcceptable)
}

func ErrorCategoryHasChildren() error {
	return errors.
		New("Cannot delete. Category has children").
		SetType(typeErrorCategoryHasChildren).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorCategoryNoChildren() error {
	return errors.
		New("Category has no children").
		SetType(typeErrorCategoryNoChildren).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorCategoryGetList() error {
	return errors.
		New("Get category list error").
		SetType(typeErrorCategoryGetList).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorCategoryGetByCode() error {
	return errors.
		New("Get category by code error").
		SetType(typeErrorCategoryGetByCode).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorCategoryGetCount() error {
	return errors.
		New("Get category count error").
		SetType(typeErrorCategoryGetCount).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorCategoryAdd() error {
	return errors.
		New("Add category error").
		SetType(typeErrorCategoryAdd).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorCategoryUpdate() error {
	return errors.
		New("Update category error").
		SetType(typeErrorCategoryUpdate).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorCategoryUpdateStatus() error {
	return errors.
		New("Update category status error").
		SetType(typeErrorCategoryUpdateStatus).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorCategoryUpdateImage() error {
	return errors.
		New("Update category image error").
		SetType(typeErrorCategoryUpdateImage).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorCategoryUpdateFileImage() error {
	return errors.
		New("Update category file image error").
		SetType(typeErrorCategoryUpdateFileImage).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorCategoryDelete() error {
	return errors.
		New("Delete category error").
		SetType(typeErrorCategoryDelete).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorSubCategoryAlreadyExist() error {
	return errors.
		New("Sub-category already exist").
		SetType(typeErrorSubCategoryAlreadyExist).
		SetHttpCode(http.StatusNotAcceptable)
}

func ErrorSubCategoryNotFound() error {
	return errors.
		New("Sub-category not found").
		SetType(typeErrorSubCategoryNotFound).
		SetHttpCode(http.StatusNotAcceptable)
}

func ErrorSubCategoryGetList() error {
	return errors.
		New("Get sub-category list error").
		SetType(typeErrorSubCategoryGetList).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorSubCategoryGetByCode() error {
	return errors.
		New("Get sub-category by code error").
		SetType(typeErrorSubCategoryGetByCode).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorSubCategoryGetCount() error {
	return errors.
		New("Get sub-category count error").
		SetType(typeErrorSubCategoryGetCount).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorSubCategoryAdd() error {
	return errors.
		New("Add sub-category error").
		SetType(typeErrorSubCategoryAdd).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorSubCategoryUpdate() error {
	return errors.
		New("Update sub-category error").
		SetType(typeErrorSubCategoryUpdate).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorSubCategoryUpdateStatus() error {
	return errors.
		New("Update sub-category status error").
		SetType(typeErrorSubCategoryUpdateStatus).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorSubCategoryUpdateImage() error {
	return errors.
		New("Update sub-category image error").
		SetType(typeErrorSubCategoryUpdateImage).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorSubCategoryUpdateFileImage() error {
	return errors.
		New("Update sub-category file image error").
		SetType(typeErrorSubCategoryUpdateFileImage).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorSubCategoryBindSub() error {
	return errors.
		New("Bind sub-category error").
		SetType(typeErrorCategoryBindSub).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorSubCategoryUnbindSub() error {
	return errors.
		New("Unbind sub-category error").
		SetType(typeErrorCategoryUnbindSub).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorSubCategoryDelete() error {
	return errors.
		New("Delete sub-category error").
		SetType(typeErrorSubCategoryDelete).
		SetHttpCode(http.StatusInternalServerError)
}
