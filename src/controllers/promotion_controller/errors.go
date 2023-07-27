package promotion_controller

import (
	"github.com/lowl11/boost/pkg/errors"
	"net/http"
)

const (
	typeErrorPromotionBind             = "PromotionBind"
	typeErrorPromotionNotFound         = "PromotionNotFound"
	typeErrorPromotionGetList          = "PromotionGetList"
	typeErrorPromotionGetByID          = "PromotionGetByID"
	typeErrorPromotionGetByCode        = "PromotionGetByCode"
	typeErrorPromotionAdd              = "PromotionAdd"
	typeErrorPromotionUpdate           = "PromotionUpdate"
	typeErrorPromotionUpdateImages     = "PromotionUpdateImages"
	typeErrorPromotionUpdateFileImages = "PromotionUpdateFileImages"
	typeErrorPromotionDelete           = "PromotionDelete"
	typeErrorPromotionDeleteFolder     = "PromotionDeleteFolder"
)

func ErrorPromotionBind() error {
	return errors.
		New("Promotion bind error").
		SetType(typeErrorPromotionBind).
		SetHttpCode(http.StatusUnprocessableEntity)
}

func ErrorPromotionNotFound() error {
	return errors.
		New("Promotion not found").
		SetType(typeErrorPromotionNotFound).
		SetHttpCode(http.StatusNotFound)
}

func ErrorPromotionGetList() error {
	return errors.
		New("Get promotions list error").
		SetType(typeErrorPromotionGetList).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorPromotionGetByID() error {
	return errors.
		New("Get promotion by id error").
		SetType(typeErrorPromotionGetByID).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorPromotionGetByCode() error {
	return errors.
		New("Get promotion by code error").
		SetType(typeErrorPromotionGetByCode).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorPromotionAdd() error {
	return errors.
		New("Add promotion error").
		SetType(typeErrorPromotionAdd).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorPromotionUpdate() error {
	return errors.
		New("Update promotion error").
		SetType(typeErrorPromotionUpdate).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorPromotionUpdateImages() error {
	return errors.
		New("Update promotion images error").
		SetType(typeErrorPromotionUpdateImages).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorPromotionUpdateFileImages() error {
	return errors.
		New("Update promotion file images error").
		SetType(typeErrorPromotionUpdateFileImages).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorPromotionDelete() error {
	return errors.
		New("Delete promotion error").
		SetType(typeErrorPromotionDelete).
		SetHttpCode(http.StatusInternalServerError)
}

func ErrorPromotionDeleteFolder() error {
	return errors.
		New("Delete promotion folder error").
		SetType(typeErrorPromotionDeleteFolder).
		SetHttpCode(http.StatusInternalServerError)
}
