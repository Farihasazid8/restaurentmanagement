package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerDto struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"customerName" bson:"customerName,omitempty"`
	DobString string             `json:"customerDob" bson:"customerDob,omitempty"`
	Email     string             `json:"customerEmail" bson:"customerEmail,omitempty"`
}
