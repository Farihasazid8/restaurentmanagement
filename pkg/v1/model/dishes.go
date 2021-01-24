package model

import (
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dishes struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Name   string             `json:"dishName" bson:"dishName,omitempty"`
	Price  float32            `json:"dishPrice" bson:"dishPrice,omitempty"`
	Status string             `json:"status" bson:"status,omitempty"`
}

func (dish Dishes) Save(context echo.Context) error {
	return nil
}
func (dish Dishes) FindAll(context echo.Context) error {
	return nil
}

func (dish Dishes) GetById(context echo.Context) error {
	return nil
}
func (dish Dishes) Delete(context echo.Context) error {
	return nil
}

func (dish Dishes) Update(context echo.Context) error {
	return nil
}
func (dish Dishes) EditQuantity(context echo.Context) error {
	return nil
}
