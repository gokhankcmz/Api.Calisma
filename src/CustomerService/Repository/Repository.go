package CustomerRepository

import (
	EntityModels "Api.Calisma/src/Common/Models/EntityModels"
	"Api.Calisma/src/Common/Models/ErrorModels"
	"Api.Calisma/src/Common/Models/RequestModels"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)



func GetCustomer(ID string) []byte {
	collection := GetMongoSingletonCollection()
	objID, err := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": objID}
	var result bson.M
	if err = collection.FindOne(context.Background(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			panic(ErrorModels.EntityNotFound.SetArgs(ID))
		}
	}
	response, _ := json.Marshal(result)
	return response
}

func GetAllCustomerIds() []byte {
	collection := GetMongoSingletonCollection()
	filter := bson.D{{}}
	res, _ := collection.Distinct(context.Background(),"_id", filter)
	response, _ := json.Marshal(res)
	return response
}
func GetAllCustomers() []byte{

	collection := GetMongoSingletonCollection()
	filter := bson.D{{}}
	var results []bson.M
	cur, _ := collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())
	cur.All(context.Background(), &results)
	response, _ := json.Marshal(results)
	return response
}

func CreateCustomer(c *EntityModels.Customer) []byte {
	collection := GetMongoSingletonCollection()
	res, _ := collection.InsertOne(context.Background(), c)
	response, _ := json.Marshal(res)
	return response
}

func UpdateCustomer(c *EntityModels.Customer) []byte {
	collection := GetMongoSingletonCollection()
	filter := bson.M{"_id": c.ID}
	res, _ := collection.ReplaceOne(context.Background(), filter, c)
	response, _ := json.Marshal(res)
	return response
}

func DeleteCustomer(ID string) []byte {
	collection := GetMongoSingletonCollection()
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": objID}
	res, _ := collection.DeleteOne(context.Background(), filter )
	response, _ := json.Marshal(res)
	return response
}

func CheckCustomerIfExist(tc *RequestModels.TokenCredentials){
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
