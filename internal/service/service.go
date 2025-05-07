package service

import "golang-with-mongo/internal/repository"

type Service struct {
	UserService        UserService
	TransactionService TransactionService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		UserService:        NewUserService(repository.UserRepository),
		TransactionService: NewTransactionService(repository.TransactionRepository),
	}
}
