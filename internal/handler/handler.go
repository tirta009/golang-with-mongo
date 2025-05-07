package handler

import (
	"golang-with-mongo/internal/service"
)

type Handler struct {
	UserHandler UserHandler
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		UserHandler: NewUserHandler(service.UserService),
	}
}
