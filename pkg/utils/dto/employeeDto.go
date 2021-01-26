package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeDto struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"employeeName" bson:"employeeName,omitempty"`
	DobString string             `json:"employeeDob" bson:"employeeDob,omitempty"`
	Email     string             `json:"employeeEmail" bson:"employeeEmail,omitempty"`
	Salary    int                `json:"employeeSalary" bson:"employeeSalary,omitempty"`
	Position  string             `json:"employeePosition" bson:"employeePosition,omitempty"`
}
