package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserId      primitive.ObjectID `bson:"user_id" json:"user_id"`
	TotalAmount float64            `bson:"total_amount" json:"total_amount"`
	Product     []Product          `bson:"product" json:"product"`
	Shipment    Shipment           `bson:"shipment" json:"shipment"`
}

type Product struct {
	ProductId int `bson:"product_id" json:"product_id"`
	Quantity  int `bson:"quantity" json:"quantity"`
	Price     int `bson:"price" json:"price"`
}

type Shipment struct {
	Province   string `bson:"province" json:"province"`
	City       string `bson:"city" json:"city"`
	Address    string `bson:"address" json:"address"`
	PostalCode string `bson:"postal_code" json:"postal_code"`
}
