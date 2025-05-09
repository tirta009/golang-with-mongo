package service

import "golang-with-mongo/internal/repository"

type Service struct {
	UserService        UserService
	TransactionService TransactionService
}

func NewService(repository *repository.Repository) *Service {

	userService := NewUserService(repository.UserRepository)
	transactionService := NewTransactionService(repository.TransactionRepository, userService)

	return &Service{
		UserService:        userService,
		TransactionService: transactionService,
	}
}
