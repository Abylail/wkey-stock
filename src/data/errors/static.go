package errors

import "wkey-stock/src/data/models"

const (
	defaultMessage = "Произошла ошибка"
)

var (
	RouteNotFound = &models.Error{
		TechMessage:     "Route not found",
		BusinessMessage: "Путь не найден",
	}

	Timeout = &models.Error{
		TechMessage:     "Request reached timed out",
		BusinessMessage: "Время работы истекло",
	}
)
