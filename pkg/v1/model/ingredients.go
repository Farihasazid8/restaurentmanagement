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

type Ingredient struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"ingredientName" bson:"ingredientName,omitempty"`
	Quantity int                `json:"quantity" bson:"quantity,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
}

func IngredientRouter(g *echo.Group) {
	ingredient := Ingredient{}
	g.POST("", ingredient.Save)
	g.GET("", ingredient.FindAll)
	g.GET("/:id", ingredient.GetById)
	g.DELETE("/:id", ingredient.Delete)
	g.PUT("", ingredient.Update)
}
func (ingredient Ingredient) Save(context echo.Context) error {
	formData := dto.IngredientDto{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	filter := bson.D{{"ingredientName", formData.Name}}
	response := db.GetDmManager().FindOne(collection.Ingredients, filter, reflect.TypeOf(Ingredient{}))
	if response != nil {
		return common.GenerateErrorResponse(context, nil, "Ingredient already exists")
	}
	var payload = Ingredient{
		ID:       primitive.NewObjectID(),
		Name:     formData.Name,
		Quantity: formData.Quantity,
		Status:   "V",
	}
	insertData, err := db.GetDmManager().InsertSingleDocument(collection.Ingredients, payload)
	if err != nil {
		log.Println("[Error]:", err)
		return common.GenerateErrorResponse(context, nil, "Failed!")
	}
	return common.GenerateSuccessResponse(context, insertData, "Inserted")
}
func (ingredient Ingredient) FindAll(context echo.Context) error {
	findAllData, err := db.GetDmManager().FindAll(collection.Ingredients, reflect.TypeOf(Ingredient{}), bson.D{{"status", "V"}}, nil, 0, -1)
	if err != nil {
		log.Println("[Error]:", err)
		return common.GenerateErrorResponse(context, nil, "Failed!")
	}
	return common.GenerateSuccessResponse(context, findAllData, "Success")
}

func (ingredient Ingredient) GetById(context echo.Context) error {
	id := context.Param("id")
	getByIdData := db.GetDmManager().FindOneByStrId(collection.Ingredients, id, reflect.TypeOf(Ingredient{}))
	fmt.Println(getByIdData)
	if getByIdData == nil {
		return common.GenerateErrorResponse(context, nil, "Failed!")
	}
	return common.GenerateSuccessResponse(context, getByIdData, "Success")
}
func (ingredient Ingredient) Delete(context echo.Context) error {
	id := context.Param("id")
	//fmt.Println("id", id)
	//filter := bson.D{{"_id", id}}
	response := db.GetDmManager().FindOneByStrId(collection.Ingredients, id, reflect.TypeOf(Ingredient{}))
	if response != nil {
		DelData := db.GetDmManager().DeleteOneByStrId(collection.Ingredients, id)
		fmt.Println("Data", DelData)
		if DelData != nil {
			return common.GenerateErrorResponse(context, nil, "Failed!")
		}
		return common.GenerateSuccessResponse(context, DelData, "Success")
	} else {
		return common.GenerateErrorResponse(context, nil, "Data Not Found!")
	}
}

func (ingredient Ingredient) Update(context echo.Context) error {
	formData := dto.IngredientDto{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	filter := bson.D{{"ingredientName", formData.Name}}
	response := db.GetDmManager().FindOne(collection.Ingredients, filter, reflect.TypeOf(Ingredient{}))
	if response != nil {
		existingUser := *response.(*Ingredient)
		existingUser.Name = formData.Name
		existingUser.Quantity = formData.Quantity

		UpdateData := db.GetDmManager().UpdateOneByStrId(collection.Ingredients, existingUser.ID.Hex(), existingUser)
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
