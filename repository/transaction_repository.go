package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-with-mongo/model"
)

type TransactionRepository interface {
	SaveTransaction(ctx context.Context, transaction model.Transaction) (primitive.ObjectID, error)
}
