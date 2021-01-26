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
	"restaurentManagement/pkg/utils/dto"
)

type Dishes struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Name   string             `json:"dishName" bson:"dishName,omitempty"`
	Price  float32            `json:"dishPrice" bson:"dishPrice,omitempty"`
	Status string             `json:"status" bson:"status,omitempty"`
}

func DishRouter(g *echo.Group) {
	dish := Dishes{}
	g.POST("", dish.Save)
	g.GET("", dish.FindAll)
	g.GET("/:id", dish.GetById)
	g.DELETE("/:id", dish.Delete)
	g.PUT("", dish.Update)
}
func (dish Dishes) Save(context echo.Context) error {
	formData := dto.DishesDto{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}

	var payload = Dishes{
		ID:     primitive.NewObjectID(),
		Name:   formData.Name,
		Price:  formData.Price,
		Status: "V",
	}
	insertData, err := db.GetDmManager().InsertSingleDocument(collection.Dishes, payload)
	if err != nil {
		log.Println("[Error]:", err)
		return common.GenerateErrorResponse(context, nil, "Failed!")
	}
	return common.GenerateSuccessResponse(context, insertData, "Inserted")
}
func (dish Dishes) FindAll(context echo.Context) error {
	findAllData, err := db.GetDmManager().FindAll(collection.Dishes, reflect.TypeOf(Dishes{}), bson.D{{"status", "V"}}, nil, 0, -1)
	if err != nil {
		log.Println("[Error]:", err)
		return common.GenerateErrorResponse(context, nil, "Failed!")
	}
	return common.GenerateSuccessResponse(context, findAllData, "Success")
}

func (dish Dishes) GetById(context echo.Context) error {
	id := context.Param("id")
	getByIdData := db.GetDmManager().FindOneByStrId(collection.Dishes, id, reflect.TypeOf(Dishes{}))
	fmt.Println(getByIdData)
	if getByIdData == nil {
		return common.GenerateErrorResponse(context, nil, "Failed!")
	}
	return common.GenerateSuccessResponse(context, getByIdData, "Success")
}
func (dish Dishes) Delete(context echo.Context) error {
	id := context.Param("id")
	//fmt.Println("id", id)
	//filter := bson.D{{"_id", id}}
	response := db.GetDmManager().FindOneByStrId(collection.Dishes, id, reflect.TypeOf(Dishes{}))
	if response != nil {
		DelData := db.GetDmManager().DeleteOneByStrId(collection.Dishes, id)
		fmt.Println("Data", DelData)
		if DelData != nil {
			return common.GenerateErrorResponse(context, nil, "Failed!")
		}
		return common.GenerateSuccessResponse(context, DelData, "Success")
	} else {
		return common.GenerateErrorResponse(context, nil, "Data Not Found!")
	}
}

func (dish Dishes) Update(context echo.Context) error {
	formData := dto.DishesDto{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	filter := bson.D{{"dishName", formData.Name}}
	response := db.GetDmManager().FindOne(collection.Dishes, filter, reflect.TypeOf(Dishes{}))
	if response != nil {
		existingUser := *response.(*Dishes)
		existingUser.Name = formData.Name
		existingUser.Price = formData.Price

		UpdateData := db.GetDmManager().UpdateOneByStrId(collection.Dishes, existingUser.ID.Hex(), existingUser)
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
