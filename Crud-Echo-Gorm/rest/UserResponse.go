package rest

import "Crud-Echo-Gorm/model"

type UserDeleteResponse struct {
	Message string `json:"message"`
}

type UserCreateResponse struct {
	Message string `json:"message"`
	User    model.User
}
