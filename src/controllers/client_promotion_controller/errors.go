package client_promotion_controller

import (
	"github.com/lowl11/boost/pkg/errors"
	"net/http"
)

const (
	typeErrorPromotionNotFound  = "PromotionNotFound"
	typeErrorPromotionGetList   = "PromotionGetList"
	typeErrorPromotionGetByCode = "PromotionGetByCode"
)

func ErrorPromotionGetList(err error) error {
	return errors.
		New("Get promotions list error").
		SetType(typeErrorPromotionGetList).
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

func ErrorPromotionNotFoundByCode(promotionCode string) error {
	return errors.
		New("Promotion not found").
		SetType(typeErrorPromotionNotFound).
		SetHttpCode(http.StatusNotFound).
		AddContext("promotion_code", promotionCode)
}
