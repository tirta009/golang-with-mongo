package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-with-mongo/internal/model"
	"golang-with-mongo/internal/payload"
)

type UserRepository interface {
	Save(ctx context.Context, user model.User) (primitive.ObjectID, error)
	Delete(ctx context.Context, id primitive.ObjectID) (bool, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*model.User, error)
	FindAll(ctx context.Context) ([]model.User, error)
	Update(ctx context.Context, id primitive.ObjectID, user *model.User) (bool, error)
	FindTotalTransactions(ctx context.Context) ([]payload.UserTransaction, error)
}

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

func (userRepository *UserRepositoryImpl) FindTotalTransactions(ctx context.Context) ([]payload.UserTransaction, error) {

	/*
		db.user.aggregate([
		    {
		        $lookup : {
		            from : "transaction",
		            localField : "_id",
		            foreignField: "user_id",
		            as : "transactions"
		        }
		    },
		    {
		        $project : {
		            _id : 0,
		            "userId" : "$_id",
		            "userName" : "$name",
		            totalTransaction : {$size : "$transactions"}
		        }
		    }
		])
	*/

	var transactions []payload.UserTransaction

	lookupStage := bson.D{
		{"$lookup", bson.D{
			{"from", "transaction"},
			{"localField", "_id"},
			{"foreignField", "user_id"},
			{"as", "transactions"},
		}},
	}

	projectStage := bson.D{
		{"$project", bson.D{
			{"_id", 0},
			{"userId", "$_id"},
			{"userName", "$name"},
			{"totalTransactions", bson.D{
				{"$size", "$transactions"},
			}},
		}},
	}

	cursor, errAgg := userRepository.userCollection.Aggregate(ctx, mongo.Pipeline{lookupStage, projectStage})
	if errAgg != nil {
		return nil, errAgg
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var request payload.UserTransaction
		err := cursor.Decode(&request)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, request)
	}

	return transactions, nil
}
