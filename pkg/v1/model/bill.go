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
	"time"
)
func BillRouter(g *echo.Group) {
	bill := Bill{}
	g.POST("", bill.Save)
	g.GET("", bill.FindAll)
	g.GET("/:id", bill.GetById)
}
type Bill struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	BillingTime   time.Time          `json:"billingTime" bson:"billingTime,omitempty"`
	OrderedDishes []string           `json:"orderedDishes" bson:"orderedDishes,omitempty"`
	BillingAmount float32            `json:"billingAmount" bson:"billingAmount,omitempty"`
	Status        string             `json:"status" bson:"status,omitempty"`
}

func (bill Bill) Save(context echo.Context) error {
	formData := dto.BillDto{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	var totalBillAmount float32 = 0.0
	for _, value := range formData.OrderedDishes {
		var objId, _ = primitive.ObjectIDFromHex(value)
		filter := bson.D{{"_id", objId}}
		response := db.GetDmManager().FindOne(collection.Dishes, filter, reflect.TypeOf(Dishes{}))
		if response != nil {
			existingDish := *response.(*Dishes)
			totalBillAmount += existingDish.Price
			//update ingredient quantity
			fmt.Println(existingDish.RequiredIngredients)
			fmt.Println("length" ,len(existingDish.RequiredIngredients))
			for key, value := range existingDish.RequiredIngredients {
				fmt.Println("Key:", key, "Value:", value)
				//ingredients.UpdateQuantityByName(context, key, value)
				filter := bson.D{{"ingredientName", key}}
				response := db.GetDmManager().FindOne(collection.Ingredients, filter, reflect.TypeOf(Ingredient{}))
				if response != nil {
					existingIngredient := *response.(*Ingredient)
					existingIngredient.Quantity -= value

					UpdateData := db.GetDmManager().UpdateOneByStrId(collection.Ingredients, existingIngredient.ID.Hex(), existingIngredient)
					fmt.Println(UpdateData)
					if UpdateData != nil {
						log.Println("[Error]:", UpdateData)
						return common.GenerateErrorResponse(context, nil, "Failed!")
					}
				} else {
					return common.GenerateErrorResponse(context, nil, "Ingredient does not exist")
				}
			}
		}


	}
	var payload = Bill{
		ID:            primitive.NewObjectID(),
		BillingAmount: totalBillAmount,
		BillingTime:   time.Now(),
		OrderedDishes: formData.OrderedDishes,
		Status:        "V",
	}
	insertData, err := db.GetDmManager().InsertSingleDocument(collection.Bill, payload)
	if err != nil {
		log.Println("[Error]:", err)
		return common.GenerateErrorResponse(context, nil, "Failed!")
	}
	return common.GenerateSuccessResponse(context, insertData, "Inserted")
}
func (bill Bill) FindAll(context echo.Context) error {
	findAllData, err := db.GetDmManager().FindAll(collection.Bill, reflect.TypeOf(Bill{}), bson.D{{"status", "V"}}, nil, 0, -1)

	if err != nil {
		log.Println("[Error]:", err)
		return common.GenerateErrorResponse(context, nil, "Failed!")
	}
	return common.GenerateSuccessResponse(context, findAllData, "Success")
}

func (bill Bill) GetById(context echo.Context) error {
	id := context.Param("id")
	getByIdData := db.GetDmManager().FindOneByStrId(collection.Bill, id, reflect.TypeOf(Bill{}))
	fmt.Println(getByIdData)
	if getByIdData == nil {
		return common.GenerateErrorResponse(context, nil, "Failed!")
	}
	return common.GenerateSuccessResponse(context, getByIdData, "Success")
}