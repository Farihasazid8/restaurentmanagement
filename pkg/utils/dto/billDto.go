package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BillDto struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	CustomerName  string             `json:"customerName" bson:"customerName,omitempty"`
	BillingAmount float32            `json:"billingAmount" bson:"billingAmount,omitempty"`
}
