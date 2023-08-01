package product_gateway

import (
	"github.com/lowl11/boost/pkg/errors"
	"net/http"
)

const (
	typeErrorProductAlreadyExist = "ProductAlreadyExist"
	typeErrorProductNotFound     = "ProductNotFound"
	typeErrorGetAllProducts      = "GetAllProducts"
	typeErrorGetByID             = "GetByID"
	typeErrorGetByProsklad       = "GetByProsklad"
	typeErrorAddProduct          = "AddProduct"
	typeErrorUpdateProduct       = "UpdateProduct"
	typeErrorRemoveProduct       = "RemoveProduct"
)

func ErrorProductAlreadyExist(proskladID int) error {
	return errors.
		New("Get product by prosklad error").
		SetType(typeErrorProductAlreadyExist).
		SetHttpCode(http.StatusConflict).
		AddContext("prosklad_id", proskladID)
}

func ErrorProductNotFoundID(id string) error {
	return errors.
		New("Product not found by id").
		SetType(typeErrorProductNotFound).
		SetHttpCode(http.StatusNotFound).
		AddContext("id", id)
}

func ErrorProductNotFoundProsklad(proskladID int) error {
	return errors.
		New("Product not found by prosklad id").
		SetType(typeErrorProductNotFound).
		SetHttpCode(http.StatusNotFound).
		AddContext("prosklad_id", proskladID)
}

func ErrorGetAllProducts(err error) error {
	return errors.
		New("Get all products error").
		SetType(typeErrorGetAllProducts).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorGetByProsklad(proskladID int, err error) error {
	return errors.
		New("Get product by id error").
		SetType(typeErrorGetByID).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err).
		AddContext("prosklad_id", proskladID)
}

func ErrorGetByID(id string, err error) error {
	return errors.
		New("Get product by prosklad error").
		SetType(typeErrorGetByProsklad).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err).
		AddContext("id", id)
}

func ErrorAddProduct(err error) error {
	return errors.
		New("Add new product error").
		SetType(typeErrorAddProduct).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorUpdateProduct(id any, err error) error {
	return errors.
		New("Update product error").
		SetType(typeErrorUpdateProduct).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err).
		AddContext("id", id)
}

func ErrorRemoveProductID(id string, err error) error {
	return errors.
		New("Remove product by id error").
		SetType(typeErrorRemoveProduct).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err).
		AddContext("id", id)
}

func ErrorRemoveProductProsklad(id int, err error) error {
	return errors.
		New("Remove product by prosklad error").
		SetType(typeErrorRemoveProduct).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err).
		AddContext("prosklad_id", id)
}
