package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-with-mongo/helper"
	"golang-with-mongo/internal/model"
	"golang-with-mongo/internal/payload"
	"golang-with-mongo/internal/repository"
)

type UserService interface {
	Create(ctx context.Context, userRequest payload.UserRequest) (primitive.ObjectID, *model.User)
	Delete(ctx context.Context, id primitive.ObjectID) bool
	FindByID(ctx context.Context, id primitive.ObjectID) *model.User
	FindAll(ctx context.Context) []model.User
	Update(ctx context.Context, id primitive.ObjectID, userRequest payload.UserRequest) (bool, *model.User)
	FindTotalTransactions(ctx context.Context) []payload.UserTransaction
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, userRequest payload.UserRequest) (primitive.ObjectID, *model.User) {

	user := model.User{
		Name: userRequest.Name,
		Age:  userRequest.Age,
	}

	objectId, err := service.UserRepository.Save(ctx, user)
	helper.PanicIfError(err)

	user.ID = objectId

	return objectId, &user

}

func (service *UserServiceImpl) Delete(ctx context.Context, objectId primitive.ObjectID) bool {

	user := service.FindByID(ctx, objectId)
	if user == nil {
		return false
	}

	deleted, err := service.UserRepository.Delete(ctx, objectId)
	helper.PanicIfError(err)

	return deleted

}

func (service *UserServiceImpl) FindByID(ctx context.Context, objectId primitive.ObjectID) *model.User {

	user, err := service.UserRepository.FindByID(ctx, objectId)
	helper.PanicIfError(err)
	if user == nil {
		return nil
	}

	return user
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []model.User {
	users, err := service.UserRepository.FindAll(ctx)
	helper.PanicIfError(err)

	return users
}

func (service *UserServiceImpl) Update(ctx context.Context, objectId primitive.ObjectID, userRequest payload.UserRequest) (bool, *model.User) {

	user := service.FindByID(ctx, objectId)
	if user == nil {
		return false, nil
	}

	user.Age = userRequest.Age
	user.Name = userRequest.Name

	update, err := service.UserRepository.Update(ctx, objectId, user)
	helper.PanicIfError(err)

	return update, user

}

func (service *UserServiceImpl) FindTotalTransactions(ctx context.Context) []payload.UserTransaction {
	transactions, err := service.UserRepository.FindTotalTransactions(ctx)
	helper.PanicIfError(err)

	return transactions
}
