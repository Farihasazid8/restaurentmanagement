package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bill struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	CustomerName string             `json:"customerName" bson:"customerName,omitempty"`
	Amount       float32            `json:"billAmount" bson:"billAmount,omitempty"`
	Status       string             `json:"status" bson:"status,omitempty"`
}
