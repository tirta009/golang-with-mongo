package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-with-mongo/dto"
	"golang-with-mongo/helper"
	"golang-with-mongo/model"
	"golang-with-mongo/repository"
)

type TransactionServiceImpl struct {
	TransactionRepository repository.TransactionRepository
}

func NewTransactionServiceImpl(transactionRepository repository.TransactionRepository) TransactionService {
	return &TransactionServiceImpl{
		TransactionRepository: transactionRepository,
	}
}

func (transactionService *TransactionServiceImpl) SaveTransaction(ctx context.Context, transactionRequest dto.TransactionRequest) primitive.ObjectID {

	userId, err := primitive.ObjectIDFromHex(transactionRequest.UserId)
	helper.PanicIfError(err)

	transaction := model.Transaction{
		UserId: userId,
	}

	objectIdTransaction, err := transactionService.TransactionRepository.SaveTransaction(ctx, transaction)
	helper.PanicIfError(err)

	return objectIdTransaction

}
