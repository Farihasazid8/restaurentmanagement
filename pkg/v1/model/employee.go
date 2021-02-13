package model

import (
	"fmt"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"reflect"
	"restaurentmanagement/common"
	"restaurentmanagement/pkg/db"
	"restaurentmanagement/pkg/db/collection"
	"restaurentmanagement/pkg/utils"
	"restaurentmanagement/pkg/utils/dto"
	"time"
)

type Employee struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"employeeName" bson:"employeeName,omitempty"`
	Dob      time.Time          `json:"employeeDob" bson:"employeeDob,omitempty"`
	Email    string             `json:"employeeEmail" bson:"employeeEmail,omitempty"`
	Salary   int                `json:"employeeSalary" bson:"employeeSalary,omitempty"`
	Position string             `json:"employeePosition" bson:"employeePosition,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
}

func EmployeeRouter(g *echo.Group) {
	employee := Employee{}
	g.POST("", employee.Save)
	g.GET("", employee.FindAll)
	g.GET("/:id", employee.GetById)
	g.DELETE("/:id", employee.Delete)
	g.PUT("", employee.Update)
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
	filter := bson.D{{"employeeEmail", formData.Email}}
	response := db.GetDmManager().FindOne(collection.Employee, filter, reflect.TypeOf(Employee{}))
	if response != nil {
		return common.GenerateErrorResponse(context, nil, "Email already exists")
	}

	var dobTime time.Time
	layout := "2006-01-02"
	dobTime, err := time.Parse(layout, formData.DobString)
	if err != nil {
		fmt.Println(err)
	}
	var payload = Employee{
		ID:       primitive.NewObjectID(),
		Name:     formData.Name,
		Dob:      dobTime,
		Email:    formData.Email,
		Salary:   formData.Salary,
		Position: formData.Position,
		Status:   "V",
	}
	insertData, err := db.GetDmManager().InsertSingleDocument(collection.Employee, payload)
	if err != nil {
		log.Println("[Error]:", err)
		return common.GenerateErrorResponse(context, nil, "Failed!")
	}
	return common.GenerateSuccessResponse(context, insertData, "Inserted")
}
func (employee Employee) FindAll(context echo.Context) error {
	findAllData, err := db.GetDmManager().FindAll(collection.Employee, reflect.TypeOf(Employee{}), bson.D{{"status", "V"}}, nil, 0, -1)

	if err != nil {
		log.Println("[Error]:", err)
		return common.GenerateErrorResponse(context, nil, "Failed!")
	}
	return common.GenerateSuccessResponse(context, findAllData, "Success")
}

func (employee Employee) GetById(context echo.Context) error {
	id := context.Param("id")
	getByIdData := db.GetDmManager().FindOneByStrId(collection.Employee, id, reflect.TypeOf(Employee{}))
	fmt.Println(getByIdData)
	if getByIdData == nil {
		return common.GenerateErrorResponse(context, nil, "Failed!")
	}
	return common.GenerateSuccessResponse(context, getByIdData, "Success")
}
func (employee Employee) Delete(context echo.Context) error {
	id := context.Param("id")
	//fmt.Println("id", id)
	//filter := bson.D{{"_id", id}}
	response := db.GetDmManager().FindOneByStrId(collection.Employee, id, reflect.TypeOf(Employee{}))
	if response != nil {
		DelData := db.GetDmManager().DeleteOneByStrId(collection.Employee, id)
		fmt.Println("Data", DelData)
		if DelData != nil {
			return common.GenerateErrorResponse(context, nil, "Failed!")
		}
		return common.GenerateSuccessResponse(context, DelData, "Success")
	} else {
		return common.GenerateErrorResponse(context, nil, "Data Not Found!")
	}
}

func (employee Employee) Update(context echo.Context) error {
	formData := dto.EmployeeDto{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	filter := bson.D{{"employeeEmail", formData.Email}}
	response := db.GetDmManager().FindOne(collection.Employee, filter, reflect.TypeOf(Employee{}))
	if response != nil {
		existingUser := *response.(*Employee)
		var t time.Time
		layout := "2006-01-02"
		t, err := time.Parse(layout, formData.DobString)
		if err != nil {
			fmt.Println(err)
		}
		existingUser.Dob = t
		existingUser.Name = formData.Name
		existingUser.Salary = formData.Salary
		existingUser.Position = formData.Position

		UpdateData := db.GetDmManager().UpdateOneByStrId(collection.Employee, existingUser.ID.Hex(), existingUser)
		fmt.Println(UpdateData)
		if UpdateData != nil {
			log.Println("[Error]:", UpdateData)
			return common.GenerateErrorResponse(context, nil, "Failed!")
		}
	} else {
		return common.GenerateErrorResponse(context, nil, "User does not exist")
	}
	return common.GenerateSuccessResponse(context, nil, "Updated")
}
