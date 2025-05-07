package main

import (
	"github.com/julienschmidt/httprouter"
	"golang-with-mongo/config"
	"golang-with-mongo/helper"
	"golang-with-mongo/internal/handler"
	"golang-with-mongo/internal/repository"
	"golang-with-mongo/internal/service"
	"log"
	"net/http"
)

func main() {
	config.InitDatabase()

	newRepository := repository.NewRepository(config.DB)
	newService := service.NewService(newRepository)
	newHandler := handler.NewHandler(newService)

	router := httprouter.New()
	RegisterUserRoutes(router, newHandler.UserHandler)

	server := http.Server{
		Addr:    "localhost:9090",
		Handler: router,
	}

	log.Println("Server started at :9090")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

func RegisterUserRoutes(router *httprouter.Router, controller handler.UserHandler) {
	const prefix = "/api/users"

	router.POST(prefix, controller.Create)
	router.GET(prefix+"/:id", controller.FindById)
	router.GET(prefix, controller.FindAll)
	router.DELETE(prefix+"/:id", controller.Delete)
	router.PUT(prefix+"/:id", controller.Update)
}
