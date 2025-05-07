package main

import (
	"github.com/julienschmidt/httprouter"
	"golang-with-mongo/controller"
	"golang-with-mongo/database"
	"golang-with-mongo/helper"
	"golang-with-mongo/repository"
	"golang-with-mongo/service"
	"log"
	"net/http"
)

func main() {
	database.InitClient()

	userRepository := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	router := httprouter.New()
	RegisterUserRoutes(router, userController)

	server := http.Server{
		Addr:    "localhost:9090",
		Handler: router,
	}

	log.Println("Server started at :9090")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

func RegisterUserRoutes(router *httprouter.Router, controller controller.UserController) {
	const prefix = "/api/users"

	router.POST(prefix, controller.Create)
	router.GET(prefix+"/:id", controller.FindById)
	router.GET(prefix, controller.FindAll)
	router.DELETE(prefix+"/:id", controller.Delete)
	router.PUT(prefix+"/:id", controller.Update)
}
