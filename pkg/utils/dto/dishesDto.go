package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dishes struct {
	ID     primitive.ObjectID `json:"id"`
	Name   string             `json:"dishName"`
	Price  float32            `json:"dishPrice"`
	Status string             `json:"status"`
}
