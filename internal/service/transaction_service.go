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
	SaveTransaction(ctx context.Context, transactionRequest payload.TransactionRequest) primitive.ObjectID
}

type TransactionServiceImpl struct {
	TransactionRepository repository.TransactionRepository
}

func NewTransactionService(transactionRepository repository.TransactionRepository) TransactionService {
	return &TransactionServiceImpl{
		TransactionRepository: transactionRepository,
	}
}

func (transactionService *TransactionServiceImpl) SaveTransaction(ctx context.Context, transactionRequest payload.TransactionRequest) primitive.ObjectID {

	userId, err := primitive.ObjectIDFromHex(transactionRequest.UserId)
	helper.PanicIfError(err)

	transaction := model.Transaction{
		UserId: userId,
	}

	objectIdTransaction, err := transactionService.TransactionRepository.SaveTransaction(ctx, transaction)
	helper.PanicIfError(err)

	return objectIdTransaction

}
