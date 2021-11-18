package OrderRepository

import (
	EntityModels "Api.Calisma/src/Common/Models/EntityModels"
	"Api.Calisma/src/Common/Models/ErrorModels"
	"Api.Calisma/src/Common/Models/RequestModels"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func GetOrder(ID string) []byte {
	collection := GetMongoSingletonCollection()
	objID, err := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": objID}
	var result bson.M
	if err = collection.FindOne(context.Background(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			panic(ErrorModels.EntityNotFound.SetArgs(ID))
		}
	}
	response, err := json.Marshal(result)
	if err != nil{
		fmt.Println(err)
	}
	return response
}

func GetAllOrderIds() []byte {
	collection := GetMongoSingletonCollection()
	filter := bson.D{{}}
	res, _ := collection.Distinct(context.Background(),"_id", filter)
	response, _ := json.Marshal(res)
	return response
}
func GetAllOrders() []byte{

	collection := GetMongoSingletonCollection()
	filter := bson.D{{}}
	var results []bson.M
	cur, _ := collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())
	cur.All(context.Background(), &results)
	response, _ := json.Marshal(results)
	return response
}

func CreateOrder(c *EntityModels.Order) []byte {
	collection := GetMongoSingletonCollection()
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	res, _ := collection.InsertOne(context.Background(), c)
	response, _ := json.Marshal(res)
	return response
}

func UpdateOrder(c *EntityModels.Order) []byte {
	collection := GetMongoSingletonCollection()
	filter := bson.M{"_id": c.ID}
	c.UpdatedAt = time.Now()
	res, err := collection.ReplaceOne(context.Background(), filter, c)
	fmt.Println(err)
	response, _ := json.Marshal(res)
	return response
}

func DeleteOrder(ID string) []byte {
	collection := GetMongoSingletonCollection()
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": objID}
	res, _ := collection.DeleteOne(context.Background(), filter )
	response, _ := json.Marshal(res)
	return response
}

func CheckOrderIfExist(tc *RequestModels.TokenCredentials){
	collection := GetMongoSingletonCollection()
	objID, err := primitive.ObjectIDFromHex(tc.ID)
	filter := bson.M{"_id": objID, "email": tc.Email}
	var result bson.M
	if err = collection.FindOne(context.Background(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			panic(ErrorModels.InvalidCredentials.SetPublicDetail("Wrong e-mail or id."))
		}
	}

}

