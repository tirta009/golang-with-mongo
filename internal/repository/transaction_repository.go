package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-with-mongo/internal/model"
)

type TransactionRepository interface {
	SaveTransaction(ctx context.Context, transaction model.Transaction) (primitive.ObjectID, error)
}

type TransactionRepositoryImpl struct {
	transactionCollection *mongo.Collection
}

func NewTransactionRepository(database *mongo.Database) TransactionRepository {
	return &TransactionRepositoryImpl{
		transactionCollection: database.Collection("transaction"),
	}
}

func (transactionRepository *TransactionRepositoryImpl) SaveTransaction(ctx context.Context, transaction model.Transaction) (primitive.ObjectID, error) {

	result, err := transactionRepository.transactionCollection.InsertOne(ctx, transaction)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil

}
