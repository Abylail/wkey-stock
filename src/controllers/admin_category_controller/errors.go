package admin_category_controller

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

func ErrorCategoryBind(err error) error {
	return errors.New("Bind category error").
		SetType(typeErrorCategoryBind).
		SetHttpCode(http.StatusUnprocessableEntity).
		SetError(err)
}

func ErrorCategoryParamRequired(paramName string) error {
	return errors.New("AdminCategory param required error").
		SetType(typeErrorCategoryParamRequired).
		SetHttpCode(http.StatusUnprocessableEntity).
		AddContext("param_name", paramName)
}

func ErrorSubCategoryBind(err error) error {
	return errors.
		New("Bind sub-category error").
		SetType(typeErrorSubCategoryBind).
		SetHttpCode(http.StatusUnprocessableEntity).
		SetError(err)
}

func ErrorCategoryAlreadyExist(code string) error {
	return errors.
		New("AdminCategory already exist").
		SetType(typeErrorCategoryAlreadyExist).
		SetHttpCode(http.StatusNotAcceptable).
		AddContext("code", code)
}

func ErrorCategoryNotFound(code string) error {
	return errors.
		New("AdminCategory not found").
		SetType(typeErrorCategoryNotFound).
		SetHttpCode(http.StatusNotAcceptable).
		AddContext("code", code)
}

func ErrorCategoryHasChildren(categoryID int) error {
	return errors.
		New("Cannot delete. AdminCategory has children").
		SetType(typeErrorCategoryHasChildren).
		SetHttpCode(http.StatusInternalServerError).
		AddContext("category_id", categoryID)
}

func ErrorCategoryNoChildren(subCategoryID int) error {
	return errors.
		New("AdminCategory has no children").
		SetType(typeErrorCategoryNoChildren).
		SetHttpCode(http.StatusInternalServerError).
		AddContext("sub_category_id", subCategoryID)
}

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

func ErrorCategoryAdd(err error) error {
	return errors.
		New("Add category error").
		SetType(typeErrorCategoryAdd).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorCategoryUpdate(err error) error {
	return errors.
		New("Update category error").
		SetType(typeErrorCategoryUpdate).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorCategoryUpdateStatus(err error) error {
	return errors.
		New("Update category status error").
		SetType(typeErrorCategoryUpdateStatus).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorCategoryUpdateImage(err error) error {
	return errors.
		New("Update category image error").
		SetType(typeErrorCategoryUpdateImage).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorCategoryUpdateFileImage(err error) error {
	return errors.
		New("Update category file image error").
		SetType(typeErrorCategoryUpdateFileImage).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorCategoryDelete(err error) error {
	return errors.
		New("Delete category error").
		SetType(typeErrorCategoryDelete).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorSubCategoryAlreadyExist(code string) error {
	return errors.
		New("Sub-category already exist").
		SetType(typeErrorSubCategoryAlreadyExist).
		SetHttpCode(http.StatusNotAcceptable).
		AddContext("code", code)
}

func ErrorSubCategoryNotFound(code string) error {
	return errors.
		New("Sub-category not found").
		SetType(typeErrorSubCategoryNotFound).
		SetHttpCode(http.StatusNotAcceptable).
		AddContext("code", code)
}

func ErrorSubCategoryGetList(err error) error {
	return errors.
		New("Get sub-category list error").
		SetType(typeErrorSubCategoryGetList).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorSubCategoryGetByCode(err error) error {
	return errors.
		New("Get sub-category by code error").
		SetType(typeErrorSubCategoryGetByCode).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorSubCategoryGetCount(err error) error {
	return errors.
		New("Get sub-category count error").
		SetType(typeErrorSubCategoryGetCount).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorSubCategoryAdd(err error) error {
	return errors.
		New("Add sub-category error").
		SetType(typeErrorSubCategoryAdd).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorSubCategoryUpdate(err error) error {
	return errors.
		New("Update sub-category error").
		SetType(typeErrorSubCategoryUpdate).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorSubCategoryUpdateStatus(err error) error {
	return errors.
		New("Update sub-category status error").
		SetType(typeErrorSubCategoryUpdateStatus).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorSubCategoryUpdateImage(err error) error {
	return errors.
		New("Update sub-category image error").
		SetType(typeErrorSubCategoryUpdateImage).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorSubCategoryUpdateFileImage(err error) error {
	return errors.
		New("Update sub-category file image error").
		SetType(typeErrorSubCategoryUpdateFileImage).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorSubCategoryBindSub(err error) error {
	return errors.
		New("Bind sub-category error").
		SetType(typeErrorCategoryBindSub).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorSubCategoryUnbindSub(err error) error {
	return errors.
		New("Unbind sub-category error").
		SetType(typeErrorCategoryUnbindSub).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorSubCategoryDelete(err error) error {
	return errors.
		New("Delete sub-category error").
		SetType(typeErrorSubCategoryDelete).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}
