package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bill struct {
	ID           primitive.ObjectID `json:"id"`
	CustomerName string             `json:"customerName"`
	Amount       float32            `json:"billAmount"`
	Status       string             `json:"status"`
}
