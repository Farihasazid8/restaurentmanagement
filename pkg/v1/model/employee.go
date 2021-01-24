package model

import (
	"fmt"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"reflect"
	"restaurentManagement/common"
	"restaurentManagement/pkg/db"
	"restaurentManagement/pkg/db/collecion"
	"restaurentManagement/pkg/utils"
	"restaurentManagement/pkg/utils/dto"
	"time"
)

type Employee struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"username" bson:"username,omitempty"`
	Dob      time.Time          `json:"userDob" bson:"userDob,omitempty"`
	Email    string             `json:"empEmail" bson:"empEmail,omitempty"`
	Salary   int                `json:"empSalary" bson:"empSalary,omitempty"`
	Position string             `json:"empPosition" bson:"empPosition,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
}

func (employee Employee) Save(context echo.Context) error {
	formData := dto.EmployeeDto{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	if !utils.IsEmailValid(formData.Email) {
		fmt.Println("Invalid Email")
		return common.GenerateErrorResponse(context, nil, "Invalid Email")
	}
	filter := bson.D{{"userEmail", formData.Email}}
	response := db.GetDmManager().FindOne(collecion.Employee, filter, reflect.TypeOf(Employee{}))
	if response != nil {
		return common.GenerateErrorResponse(context, nil, "Email already exists")
	}

	var t time.Time
	layout := "2006-01-02"
	t, err := time.Parse(layout, formData.DobStr)
	if err != nil {
		fmt.Println(err)
	}
	var payload = Employee{
		ID:       primitive.NewObjectID(),
		Name:     formData.Name,
		Dob:      t,
		Email:    formData.Email,
		Salary:   formData.Salary,
		Position: formData.Position,
		Status:   "V",
	}
	insertData, err := db.GetDmManager().InsertSingleDocument(collecion.Employee, payload)
	if err != nil {
		log.Println("[Error]:", err)
		return common.GenerateErrorResponse(context, nil, "Failed!")
	}
	return common.GenerateSuccessResponse(context, insertData, "Inserted")
	return nil
}
func (employee Employee) FindAll(context echo.Context) error {
	return nil
}

func (employee Employee) GetById(context echo.Context) error {
	return nil
}
func (employee Employee) Delete(context echo.Context) error {
	return nil
}

func (employee Employee) Update(context echo.Context) error {
	return nil
}
