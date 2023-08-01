package category_gateway

import (
	"github.com/lowl11/boost/pkg/errors"
	"net/http"
)

const (
	typeErrorCategoryAlreadyExist = "CategoryAlreadyExist"
	typeErrorCategoryNotFound     = "CategoryNotFound"
	typeErrorGetAllCategories     = "GetAllCategories"
	typeErrorGetByID              = "GetByID"
	typeErrorGetByProsklad        = "GetByProsklad"
	typeErrorAddCategory          = "AddCategory"
	typeErrorUpdateCategory       = "UpdateCategory"
	typeErrorRemoveCategory       = "RemoveCategory"
)

func ErrorCategoryAlreadyExist(proskladID int) error {
	return errors.
		New("Get category by prosklad error").
		SetType(typeErrorCategoryAlreadyExist).
		SetHttpCode(http.StatusConflict).
		AddContext("prosklad_id", proskladID)
}

func ErrorCategoryNotFound(id any) error {
	return errors.
		New("Category not found by id").
		SetType(typeErrorCategoryNotFound).
		SetHttpCode(http.StatusNotFound).
		AddContext("id", id)
}

func ErrorGetAllCategories(err error) error {
	return errors.
		New("Get all categories error").
		SetType(typeErrorGetAllCategories).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorGetByProsklad(proskladID int, err error) error {
	return errors.
		New("Get category by id error").
		SetType(typeErrorGetByID).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err).
		AddContext("prosklad_id", proskladID)
}

func ErrorGetByID(id string, err error) error {
	return errors.
		New("Get category by prosklad error").
		SetType(typeErrorGetByProsklad).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err).
		AddContext("id", id)
}

func ErrorAddCategory(err error) error {
	return errors.
		New("Add new category error").
		SetType(typeErrorAddCategory).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorUpdateCategory(id any, err error) error {
	return errors.
		New("Update category error").
		SetType(typeErrorUpdateCategory).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err).
		AddContext("id", id)
}

func ErrorRemoveCategory(id any, err error) error {
	return errors.
		New("Remove category by id error").
		SetType(typeErrorRemoveCategory).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err).
		AddContext("id", id)
}
