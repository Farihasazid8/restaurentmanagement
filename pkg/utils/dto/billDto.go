package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type BillDto struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	OrderedDishes []string           `json:"orderedDishes" bson:"orderedDishes,omitempty"`
	BillingAmount float32            `json:"billingAmount" bson:"billingAmount,omitempty"`
}