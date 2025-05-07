package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-with-mongo/dto"
	"golang-with-mongo/helper"
	"golang-with-mongo/model"
	"golang-with-mongo/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, userRequest dto.UserRequest) (primitive.ObjectID, *model.User) {

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

func (service *UserServiceImpl) Update(ctx context.Context, objectId primitive.ObjectID, userRequest dto.UserRequest) (bool, *model.User) {

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
