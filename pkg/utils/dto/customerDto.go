package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CustomerDto struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Name   string             `json:"username"`
	Dob    time.Time          `json:"userDob"`
	Email  string             `json:"empEmail"`
	Status string             `json:"status"`
}
