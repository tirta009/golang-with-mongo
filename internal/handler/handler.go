package handler

import (
	"golang-with-mongo/internal/service"
)

type Handler struct {
	UserHandler        UserHandler
	TransactionHandler TransactionHandler
}

func NewHandler(service *service.Service) *Handler {

	return &Handler{
		UserHandler:        NewUserHandler(service.UserService),
		TransactionHandler: NewTransactionHandler(service.TransactionService),
	}
}
