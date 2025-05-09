package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-with-mongo/helper"
	"golang-with-mongo/internal/model"
	"golang-with-mongo/internal/payload"
	"golang-with-mongo/internal/repository"
)

type TransactionService interface {
	SaveTransaction(ctx context.Context, transactionRequest payload.TransactionRequest) (primitive.ObjectID, *model.Transaction)
	FindTotalTransactions(ctx context.Context) []payload.UserTransaction
}

type TransactionServiceImpl struct {
	TransactionRepository repository.TransactionRepository
	UserService           UserService
}

func NewTransactionService(transactionRepository repository.TransactionRepository, userService UserService) TransactionService {
	return &TransactionServiceImpl{
		TransactionRepository: transactionRepository,
		UserService:           userService,
	}
}

func (transactionService *TransactionServiceImpl) SaveTransaction(ctx context.Context, transactionRequest payload.TransactionRequest) (primitive.ObjectID, *model.Transaction) {

	userId, err := primitive.ObjectIDFromHex(transactionRequest.UserId)
	helper.PanicIfError(err)

	user := transactionService.UserService.FindByID(ctx, userId)
	if user == nil {
		return primitive.NilObjectID, nil
	}

	var products []model.Product
	for _, productReq := range transactionRequest.ProductRequest {
		products = append(products, model.Product{
			ProductId: productReq.ProductId,
			Quantity:  productReq.Quantity,
			Price:     productReq.Price,
		})
	}

	shipment := model.Shipment{
		Province:   transactionRequest.ShipmentRequest.Province,
		City:       transactionRequest.ShipmentRequest.City,
		Address:    transactionRequest.ShipmentRequest.Address,
		PostalCode: transactionRequest.ShipmentRequest.PostalCode,
	}

	transaction := model.Transaction{
		UserId:      userId,
		TotalAmount: transactionRequest.TotalAmount,
		Product:     products,
		Shipment:    shipment,
	}

	objectIdTransaction, err := transactionService.TransactionRepository.SaveTransaction(ctx, transaction)
	helper.PanicIfError(err)

	transaction.Id = objectIdTransaction

	return objectIdTransaction, &transaction

}

func (transactionService *TransactionServiceImpl) FindTotalTransactions(ctx context.Context) []payload.UserTransaction {
	return transactionService.UserService.FindTotalTransactions(ctx)
}
