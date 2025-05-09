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
	RegisterTransactionRoutes(router, newHandler.TransactionHandler)

	server := http.Server{
		Addr:    "localhost:9090",
		Handler: router,
	}

	log.Println("Server started at :9090")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

func RegisterUserRoutes(router *httprouter.Router, userHandler handler.UserHandler) {
	const prefix = "/api/users"

	router.POST(prefix, userHandler.Create)
	router.GET(prefix+"/:id", userHandler.FindById)
	router.GET(prefix, userHandler.FindAll)
	router.DELETE(prefix+"/:id", userHandler.Delete)
	router.PUT(prefix+"/:id", userHandler.Update)

}

func RegisterTransactionRoutes(router *httprouter.Router, transactionHandler handler.TransactionHandler) {
	const prefix = "/api/transactions"

	router.POST(prefix, transactionHandler.CreateTransaction)
	router.GET(prefix+"/users", transactionHandler.FindTotalTransactions)
}
