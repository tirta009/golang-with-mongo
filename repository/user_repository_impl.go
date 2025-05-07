package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-with-mongo/model"
)

type UserRepositoryImpl struct {
	userCollection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &UserRepositoryImpl{
		userCollection: db.Collection("user"),
	}
}

func (userRepository *UserRepositoryImpl) Save(ctx context.Context, user model.User) (primitive.ObjectID, error) {

	result, err := userRepository.userCollection.InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (userRepository *UserRepositoryImpl) Delete(ctx context.Context, objectID primitive.ObjectID) (bool, error) {

	filter := bson.M{"_id": objectID}
	deleted, err := userRepository.userCollection.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}

	return deleted.DeletedCount > 0, nil
}

func (userRepository *UserRepositoryImpl) FindByID(ctx context.Context, id primitive.ObjectID) (*model.User, error) {
	user := &model.User{}
	filter := bson.M{"_id": id}
	result := userRepository.userCollection.FindOne(ctx, filter)
	if errors.Is(result.Err(), mongo.ErrNoDocuments) {
		return nil, nil
	}

	err := result.Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (userRepository *UserRepositoryImpl) FindAll(ctx context.Context) ([]model.User, error) {
	var users []model.User
	cursor, errFind := userRepository.userCollection.Find(ctx, bson.M{})
	if errFind != nil {
		return nil, errFind
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user model.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (userRepository *UserRepositoryImpl) Update(ctx context.Context, id primitive.ObjectID, user *model.User) (bool, error) {
	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": bson.M{
			"age":  user.Age,
			"name": user.Name,
		},
	}

	updated, err := userRepository.userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}
	return updated.ModifiedCount > 0, nil

}
