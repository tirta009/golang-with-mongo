package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-with-mongo/dto"
)

type TransactionService interface {
	SaveTransaction(ctx context.Context, transactionRequest dto.TransactionRequest) primitive.ObjectID
}
