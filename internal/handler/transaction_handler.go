package handler

import (
	"github.com/julienschmidt/httprouter"
	"golang-with-mongo/helper"
	"golang-with-mongo/internal/payload"
	"golang-with-mongo/internal/service"
	"net/http"
)

type TransactionHandler interface {
	CreateTransaction(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindTotalTransactions(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type TransactionHandlerImpl struct {
	TransactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) TransactionHandler {
	return &TransactionHandlerImpl{
		TransactionService: transactionService,
	}
}

func (transactionHandler *TransactionHandlerImpl) CreateTransaction(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	transactionRequest := payload.TransactionRequest{}
	helper.ReadFromRequestBody(request, &transactionRequest)

	objectId, transaction := transactionHandler.TransactionService.SaveTransaction(request.Context(), transactionRequest)
	if transaction == nil {
		helper.WriteErrorResponse(writer, http.StatusNotFound, "User not found")
		return
	}

	helper.WriteSuccessResponse(writer, http.StatusOK, "Successfully created with Id : "+objectId.Hex(), transaction)

}

func (transactionHandler *TransactionHandlerImpl) FindTotalTransactions(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userTransactions := transactionHandler.TransactionService.FindTotalTransactions(request.Context())

	helper.WriteSuccessResponse(writer, http.StatusOK, "", userTransactions)
}
