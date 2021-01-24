package model

import (
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ingredient struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"username" bson:"username,omitempty"`
	Quantity int                `json:"empSalary" bson:"empSalary,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
}

func (ingredient Ingredient) Save(context echo.Context) error {
	return nil
}
func (ingredient Ingredient) FindAll(context echo.Context) error {
	return nil
}

func (ingredient Ingredient) GetById(context echo.Context) error {
	return nil
}
func (ingredient Ingredient) Delete(context echo.Context) error {
	return nil
}

func (ingredient Ingredient) Update(context echo.Context) error {
	return nil
}
func (ingredient Ingredient) EditQuantity(context echo.Context) error {
	return nil
}
