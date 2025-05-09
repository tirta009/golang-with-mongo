package payload

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserTransaction struct {
	UserId            primitive.ObjectID `json:"userId"`
	UserName          string             `json:"userName"`
	TotalTransactions int32              `json:"totalTransactions"`
}
