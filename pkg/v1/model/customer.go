package model

import (
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Customer struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Name   string             `json:"customerName" bson:"customerName,omitempty"`
	Dob    time.Time          `json:"customerDob" bson:"customerDob,omitempty"`
	Email  string             `json:"customerEmail" bson:"customerEmail,omitempty"`
	Status string             `json:"status" bson:"status,omitempty"`
}

func (customer Customer) Save(context echo.Context) error {
	return nil
}
func (customer Customer) FindAll(context echo.Context) error {
	return nil
}

func (customer Customer) GetById(context echo.Context) error {
	return nil
}
func (customer Customer) Delete(context echo.Context) error {
	return nil
}

func (customer Customer) Update(context echo.Context) error {
	return nil
}
