package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ingredient struct {
	ID       primitive.ObjectID `json:"id"`
	Name     string             `json:"username"`
	Quantity int                `json:"empSalary"`
	Status   string             `json:"status"`
}
