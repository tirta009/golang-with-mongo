package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-with-mongo/model"
)

type UserRepository interface {
	Save(ctx context.Context, user model.User) (primitive.ObjectID, error)
	Delete(ctx context.Context, id primitive.ObjectID) (bool, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*model.User, error)
	FindAll(ctx context.Context) ([]model.User, error)
	Update(ctx context.Context, id primitive.ObjectID, user *model.User) (bool, error)
}
