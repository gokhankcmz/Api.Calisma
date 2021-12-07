package OrderRepository

import (
	EntityModels "Api.Calisma/src/Common/Models/EntityModels"
	"Api.Calisma/src/Common/Models/ErrorModels"
	"Api.Calisma/src/Common/Models/RequestModels"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type IOrderRepository interface {
	GetOrder(ID string) EntityModels.Order
	GetAllOrderIds() ([]string, int)
	GetAllOrders() []EntityModels.Order
	CreateOrder(c *EntityModels.Order) string
	UpdateOrder(c *EntityModels.Order) string
	CheckOrderIfExist(tc *RequestModels.TokenCredentials)
	DeleteOrder(ID string) int64
}

type Repository struct {
	mc *mongo.Collection
}

func NewRepository(mc *mongo.Collection) *Repository {
	return &Repository{mc: mc}
}

func (r Repository) GetOrder(ID string) EntityModels.Order {
	collection := GetMongoSingletonCollection()
	objID, err := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": objID}
	var result EntityModels.Order
	if err = collection.FindOne(context.Background(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			panic(ErrorModels.EntityNotFound.SetArgs(ID))
		}
	}
	return result
}

func (r Repository) GetAllOrderIds() ([]string, int) {
	collection := GetMongoSingletonCollection()
	filter := bson.D{{}}
	res, _ := collection.Distinct(context.Background(), "_id", filter)
	response, _ := json.Marshal(res)
	var resp []string
	json.Unmarshal(response, &resp)
	return resp, len(resp)
}
func (r Repository) GetAllOrders() []EntityModels.Order {
	collection := GetMongoSingletonCollection()
	filter := bson.D{{}}
	var results []EntityModels.Order
	cur, _ := collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())
	return results
}

func (r Repository) CreateOrder(c *EntityModels.Order) string {
	collection := GetMongoSingletonCollection()
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	res, _ := collection.InsertOne(context.Background(), c)
	return res.InsertedID.(primitive.ObjectID).Hex()
}

func (r Repository) UpdateOrder(c *EntityModels.Order) string {
	collection := GetMongoSingletonCollection()
	filter := bson.M{"_id": c.ID}
	c.UpdatedAt = time.Now()
	res, _ := collection.ReplaceOne(context.Background(), filter, c)
	return res.UpsertedID.(primitive.ObjectID).Hex()
}

func (r Repository) DeleteOrder(ID string) int64 {
	collection := GetMongoSingletonCollection()
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": objID}
	res, _ := collection.DeleteOne(context.Background(), filter)
	return res.DeletedCount
}

func (r Repository) CheckOrderIfExist(tc *RequestModels.TokenCredentials) {
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
