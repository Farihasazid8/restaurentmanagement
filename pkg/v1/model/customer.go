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
	"restaurentManagement/pkg/db/collection"
	"restaurentManagement/pkg/utils"
	"restaurentManagement/pkg/utils/dto"
	"time"
)

type Customer struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Name   string             `json:"customerName" bson:"customerName,omitempty"`
	Dob    time.Time          `json:"customerDob" bson:"customerDob,omitempty"`
	Email  string             `json:"customerEmail" bson:"customerEmail,omitempty"`
	Status string             `json:"status" bson:"status,omitempty"`
}

func CustomerRouter(g *echo.Group) {
	customer := Customer{}
	g.POST("", customer.Save)
	g.GET("", customer.FindAll)
	g.GET("/:id", customer.GetById)
	g.DELETE("/:id", customer.Delete)
	g.PUT("", customer.Update)
}

func (customer Customer) Save(context echo.Context) error {
	formData := dto.CustomerDto{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	if !utils.IsEmailValid(formData.Email) {
		fmt.Println("Invalid Email")
		return common.GenerateErrorResponse(context, nil, "Invalid Email")
	}
	filter := bson.D{{"customerEmail", formData.Email}}
	response := db.GetDmManager().FindOne(collection.Customer, filter, reflect.TypeOf(Customer{}))
	if response != nil {
		return common.GenerateErrorResponse(context, nil, "Email already exists")
	}

	var t time.Time
	layout := "2006-01-02"
	t, err := time.Parse(layout, formData.DobString)
	if err != nil {
		fmt.Println(err)
	}
	var payload = Customer{
		ID:     primitive.NewObjectID(),
		Name:   formData.Name,
		Dob:    t,
		Email:  formData.Email,
		Status: "V",
	}
	insertData, err := db.GetDmManager().InsertSingleDocument(collection.Customer, payload)
	if err != nil {
		log.Println("[Error]:", err)
		return common.GenerateErrorResponse(context, nil, "Failed!")
	}
	return common.GenerateSuccessResponse(context, insertData, "Inserted")
}
func (customer Customer) FindAll(context echo.Context) error {
	findAllData, err := db.GetDmManager().FindAll(collection.Customer, reflect.TypeOf(Customer{}), bson.D{{"status", "V"}}, nil, 0, -1)

	if err != nil {
		log.Println("[Error]:", err)
		return common.GenerateErrorResponse(context, nil, "Failed!")
	}
	return common.GenerateSuccessResponse(context, findAllData, "Success")
}

func (customer Customer) GetById(context echo.Context) error {
	id := context.Param("id")
	getByIdData := db.GetDmManager().FindOneByStrId(collection.Customer, id, reflect.TypeOf(Customer{}))
	fmt.Println(getByIdData)
	if getByIdData == nil {
		return common.GenerateErrorResponse(context, nil, "Failed!")
	}
	return common.GenerateSuccessResponse(context, getByIdData, "Success")
}
func (customer Customer) Delete(context echo.Context) error {
	id := context.Param("id")
	//fmt.Println("id", id)
	//filter := bson.D{{"_id", id}}
	response := db.GetDmManager().FindOneByStrId(collection.Customer, id, reflect.TypeOf(Customer{}))
	if response != nil {
		DelData := db.GetDmManager().DeleteOneByStrId(collection.Customer, id)
		fmt.Println("Data", DelData)
		if DelData != nil {
			return common.GenerateErrorResponse(context, nil, "Failed!")
		}
		return common.GenerateSuccessResponse(context, DelData, "Success")
	} else {
		return common.GenerateErrorResponse(context, nil, "Data Not Found!")
	}
}

func (customer Customer) Update(context echo.Context) error {
	formData := dto.CustomerDto{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	filter := bson.D{{"customerEmail", formData.Email}}
	response := db.GetDmManager().FindOne(collection.Customer, filter, reflect.TypeOf(Customer{}))
	if response != nil {
		existingUser := *response.(*Customer)
		var dobTime time.Time
		layout := "2006-01-02"
		dobTime, err := time.Parse(layout, formData.DobString)
		if err != nil {
			fmt.Println(err)
		}
		existingUser.Dob = dobTime
		existingUser.Name = formData.Name

		UpdateData := db.GetDmManager().UpdateOneByStrId(collection.Customer, existingUser.ID.Hex(), existingUser)
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
