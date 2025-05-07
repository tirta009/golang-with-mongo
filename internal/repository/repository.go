package repository

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
	UserRepository        UserRepository
	TransactionRepository TransactionRepository
}

func NewRepository(database *mongo.Database) *Repository {
	return &Repository{
		UserRepository:        NewUserRepository(database),
		TransactionRepository: NewTransactionRepository(database),
	}
}
