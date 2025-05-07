package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-with-mongo/dto"
	"golang-with-mongo/model"
)

type UserService interface {
	Create(ctx context.Context, userRequest dto.UserRequest) (primitive.ObjectID, *model.User)
	Delete(ctx context.Context, id primitive.ObjectID) bool
	FindByID(ctx context.Context, id primitive.ObjectID) *model.User
	FindAll(ctx context.Context) []model.User
	Update(ctx context.Context, id primitive.ObjectID, userRequest dto.UserRequest) (bool, *model.User)
}
