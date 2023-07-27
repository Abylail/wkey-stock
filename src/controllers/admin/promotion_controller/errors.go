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

func ErrorPromotionBind(err error) error {
	return errors.
		New("Promotion bind error").
		SetType(typeErrorPromotionBind).
		SetHttpCode(http.StatusUnprocessableEntity).
		SetError(err)
}

func ErrorPromotionNotFoundByCode(promotionCode string) error {
	return errors.
		New("Promotion not found").
		SetType(typeErrorPromotionNotFound).
		SetHttpCode(http.StatusNotFound).
		AddContext("promotion_code", promotionCode)
}

func ErrorPromotionNotFoundByID(promotionID int) error {
	return errors.
		New("Promotion not found").
		SetType(typeErrorPromotionNotFound).
		SetHttpCode(http.StatusNotFound).
		AddContext("promotion_id", promotionID)
}

func ErrorPromotionGetList(err error) error {
	return errors.
		New("Get promotions list error").
		SetType(typeErrorPromotionGetList).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorPromotionGetByID(err error) error {
	return errors.
		New("Get promotion by id error").
		SetType(typeErrorPromotionGetByID).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorPromotionGetByCode(err error) error {
	return errors.
		New("Get promotion by code error").
		SetType(typeErrorPromotionGetByCode).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorPromotionAdd(err error) error {
	return errors.
		New("Add promotion error").
		SetType(typeErrorPromotionAdd).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorPromotionUpdate(err error) error {
	return errors.
		New("Update promotion error").
		SetType(typeErrorPromotionUpdate).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorPromotionUpdateImages(err error) error {
	return errors.
		New("Update promotion images error").
		SetType(typeErrorPromotionUpdateImages).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorPromotionUpdateFileImages(err error) error {
	return errors.
		New("Update promotion file images error").
		SetType(typeErrorPromotionUpdateFileImages).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorPromotionDelete(err error) error {
	return errors.
		New("Delete promotion error").
		SetType(typeErrorPromotionDelete).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}

func ErrorPromotionDeleteFolder(err error) error {
	return errors.
		New("Delete promotion folder error").
		SetType(typeErrorPromotionDeleteFolder).
		SetHttpCode(http.StatusInternalServerError).
		SetError(err)
}
