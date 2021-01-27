package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type DishesDto struct {
	ID                  primitive.ObjectID `json:"id" bson:"_id"`
	Name                string             `json:"dishName" bson:"dishName,omitempty"`
	Price               float32            `json:"dishPrice" bson:"dishPrice,omitempty"`
	RequiredIngredients []IngredientDto    `json:"requiredIngredients" bson:"requiredIngredients,omitempty"`
}