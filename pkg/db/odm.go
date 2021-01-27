package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"reflect"
	"restaurentManagement/config"
	_ "restaurentManagement/pkg/type"
	_type "restaurentManagement/pkg/type"
	"sync"
	"time"
)

type Status string

//Dummy Interface
const (
	Valid   = Status("V")
	Deleted = Status("D")
	Hidden  = Status("H")
)

type dmManager struct {
	connected bool
	ctx       context.Context
	client    *mongo.Client
	db        *mongo.Database
}
type DbFindObj interface{}

// Implementing Singleton
var singletonDmManager *dmManager
var onceDmManager sync.Once

func GetDmManager() *dmManager {
	onceDmManager.Do(func() {
		log.Println("[INFO] Starting Initializing Singleton DB Manager")
		singletonDmManager = &dmManager{}
		singletonDmManager.connected = singletonDmManager.initConnection()
	})
	return singletonDmManager
}

func (dm *dmManager) initConnection() bool {
	// Base context.
	ctx := context.Background()
	dm.ctx = ctx
	//clientOpts := options.Client().ApplyURI(config.DatabaseConnectionString)
	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&ssl=false")

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Println("[ERROR] DB Connection error:", err.Error())
		return false
	}

	dm.client = client
	if !dm.ping() {
		return false
	}

	db := client.Database(config.DatabaseName)
	dm.db = db

	log.Println("[INFO] Initialized Singleton DB Manager")

	go DbHealthChecker()

	return true
}
func DbHealthChecker() {
	log.Println("[INFO] Starting db health checker")
	dm := GetDmManager()
	for {
		if !dm.ping() {
			dm.connected = false
			break
		}
		time.Sleep(time.Duration(5) * time.Second)
	}
	dm.reconnect()
}

func (dm *dmManager) ping() bool {
	if dm.client != nil {
		err := dm.client.Ping(context.TODO(), nil)
		if err != nil {
			log.Println("[ERROR] DB ping failed", err.Error())
			return false
		}
		return true
	} else {
		log.Println("[ERROR] DB Client doesn't exist")
	}
	return false
}

func (dm *dmManager) IsConnected() bool {
	return dm.connected
}

func (dm *dmManager) reconnect() {
	log.Println("[INFO] Trying to reconnect to DB")
	for !dm.initConnection() {
		time.Sleep(time.Duration(1) * time.Second)
		log.Println("[INFO] Trying to reconnect to DB")
	}
	dm.connected = true
	log.Println("[INFO] DB reconnected")
}

func (dm *dmManager) InsertSingleDocument(collectionName string, document interface{}) (primitive.ObjectID, error) {
	coll := dm.db.Collection(collectionName)
	result, err := coll.InsertOne(dm.ctx, document)
	if err != nil {
		log.Println("[ERROR] Insert document:", err.Error())
		return primitive.ObjectID{}, err
	}

	// ID of the inserted document.
	objectID := result.InsertedID.(primitive.ObjectID)
	return objectID, nil
}

func (dm *dmManager) InsertMultipleDocument(collectionName string, documents []interface{}) ([]interface{}, error) {
	coll := dm.db.Collection(collectionName)
	results, err := coll.InsertMany(dm.ctx, documents)
	if err != nil {
		log.Println("[ERROR] Insert multiple documents:", err.Error())
		return nil, err
	}

	return results.InsertedIDs, nil
}

func (dm *dmManager) FindOne(collectionName string, filter interface{}, objType reflect.Type) interface{} {
	coll := dm.db.Collection(collectionName)

	findResult := coll.FindOne(dm.ctx, filter)
	if err := findResult.Err(); err != nil {
		return nil
	}

	objValue := reflect.New(objType)
	obj := objValue.Interface()
	err := findResult.Decode(obj)
	if err != nil {
		log.Println("[ERROR] Find document decoding:", err.Error())
		return nil
	}
	return obj
}

func (dm *dmManager) FindAll(collectionName string, objType reflect.Type, filter interface{}, sortParam *_type.SortParam, skip int64, limit int64) ([]DbFindObj, []DbFindObj) {
	// Pass these options to the Find method
	findOptions := options.Find()
	if sortParam != nil {
		findOptions.SetSort(bson.D{{sortParam.SortBy, sortParam.Type}})
	}
	if skip > 0 {
		findOptions.SetSkip(skip)
	}
	if limit > -1 {
		findOptions.SetLimit(limit)
	}
	coll := dm.db.Collection(collectionName)
	reflect.New(objType)
	var results []DbFindObj

	// Passing bson.D{{}} as the filter matches all documents in the collection
	if filter == nil {
		filter = bson.M{"status": "V"}
	}
	cur, err := coll.Find(dm.ctx, filter, findOptions)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		elemValue := reflect.New(objType)
		elem := elemValue.Interface()
		err := cur.Decode(elem)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, elem)
	}
	if err := cur.Err(); err != nil {
		log.Println("[ERROR]", err)
	} else {
		// Close the cursor once finished
		cur.Close(context.TODO())
	}
	return results, nil
}

func (dm *dmManager) FindOneByObjId(collectionName string, objId primitive.ObjectID, objType reflect.Type) interface{} {
	coll := dm.db.Collection(collectionName)
	fmt.Println("Hello from odm findbyObjid", objId, objType)
	findResult := coll.FindOne(dm.ctx, bson.M{"_id": objId, "status": "V"})
	//findResult := coll.FindOne(dm.ctx, bson.D{{"_id", objId}})
	fmt.Println("Hello from odm findbyObjid--find result", findResult)
	if err := findResult.Err(); err != nil {
		return nil
	}

	objValue := reflect.New(objType)
	obj := objValue.Interface()
	err := findResult.Decode(obj)
	if err != nil {
		log.Println("[ERROR] Find document decoding:", err.Error())
		return nil
	}
	return obj
}

func (dm *dmManager) FindOneByStrId(collectionName string, strId string, objType reflect.Type) interface{} {
	objId, err := primitive.ObjectIDFromHex(strId)
	if err != nil {
		log.Println("[ERROR] Converting Str ID to Obj ID:", err.Error())
		return nil
	}
	return dm.FindOneByObjId(collectionName, objId, objType)
}

func (dm *dmManager) UpdateOneByObjId(collectionName string, objId primitive.ObjectID, document interface{}) error {
	coll := dm.db.Collection(collectionName)
	fmt.Println("Document-", document)
	//filter := bson.M{"_id": objId, "status": Valid}
	filter := bson.M{"_id": objId, "status": "V"}
	update := bson.M{"$set": document}

	// Call the driver's UpdateOne() method and pass filter and update to it
	_, err := coll.UpdateOne(
		dm.ctx,
		filter,
		update,
	)

	return err
}

func (dm *dmManager) UpdateOneByStrId(collectionName string, strId string, document interface{}) error {
	objId, err := primitive.ObjectIDFromHex(strId)
	if err != nil {
		log.Println("[ERROR]: Converting Str ID to Obj ID:", err.Error())
		return nil
	}
	return dm.UpdateOneByObjId(collectionName, objId, document)
}

func (dm *dmManager) DeleteOneByObjId(collectionName string, objId primitive.ObjectID) error {
	fmt.Println("odm2, objId", objId)
	coll := dm.db.Collection(collectionName)
	//filter := bson.M{"_id": objId, "status": Valid}
	filter := bson.M{"_id": objId, "status": "V"}
	update := bson.M{"$set": bson.M{"status": Deleted}}
	// Call the driver's UpdateOne() method and pass filter and update to it
	_, err := coll.UpdateOne(
		dm.ctx,
		filter,
		update,
	)
	return err
}

func (dm *dmManager) DeleteOneByStrId(collectionName string, strId string) error {
	objId, err := primitive.ObjectIDFromHex(strId)
	fmt.Println("odm, objId ", objId)
	if err != nil {
		log.Println("[ERROR] Converting Str ID to Obj ID:", err.Error())
		return nil
	}
	return dm.DeleteOneByObjId(collectionName, objId)
}

func (dm *dmManager) RestoreOneByObjId(collectionName string, objId primitive.ObjectID) error {
	coll := dm.db.Collection(collectionName)
	filter := bson.M{"_id": objId, "status": Deleted}
	update := bson.M{"$set": bson.M{"status": Valid}}
	// Call the driver's UpdateOne() method and pass filter and update to it
	_, err := coll.UpdateOne(
		dm.ctx,
		filter,
		update,
	)
	return err
}

func (dm *dmManager) PermanentDeleteOneByObjId(collectionName string, objId primitive.ObjectID) error {
	coll := dm.db.Collection(collectionName)

	filter := bson.M{"_id": objId}

	_, err := coll.DeleteOne(dm.ctx, filter)
	return err
}

func (dm *dmManager) PermanentDeleteOneByStrId(collectionName string, strId string) error {
	objId, err := primitive.ObjectIDFromHex(strId)
	if err != nil {
		log.Println("[ERROR] Converting Str ID to Obj ID:", err.Error())
		return nil
	}
	return dm.PermanentDeleteOneByObjId(collectionName, objId)
}

func (dm *dmManager) DropCollection(collectionName string) error {
	coll := dm.db.Collection(collectionName)
	err := coll.Drop(dm.ctx)
	return err
}
