package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeDto struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"username"`
	DobStr   string             `json:"userDob"`
	Email    string             `json:"empEmail"`
	Salary   int                `json:"empSalary"`
	Position string             `json:"empPosition"`
	Status   string             `json:"status"`
}
