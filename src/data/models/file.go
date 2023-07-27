package models

type File struct {
	Name   string `json:"name" validate:"required"`
	Buffer string `json:"buffer" validate:"required"`
}
