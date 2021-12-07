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
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

//TODO: SOYUTLA
// TODO: ECHO BIND ARASTIR

type IRepository interface {
	GetCustomer(ID string) EntityModels.Customer
	GetAllCustomerIds() ([]string, int)
	GetCustomers(options *options.FindOptions, filter *bson.M) ([]EntityModels.Customer, int)
	CreateCustomer(c *EntityModels.Customer) string
	UpdateCustomer(c *EntityModels.Customer) *mongo.UpdateResult
	DeleteCustomer(ID string) *mongo.DeleteResult
	CheckCustomerIfExist(tc *RequestModels.TokenCredentials)
}

type Repository struct {
	mc *mongo.Collection
}

func NewRepository(mc *mongo.Collection) *Repository {
	return &Repository{mc: mc}
}

func (r Repository) GetCustomer(ID string) EntityModels.Customer {
	objID, err := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": objID}
	var result EntityModels.Customer
	if err = r.mc.FindOne(context.Background(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			panic(ErrorModels.EntityNotFound.SetArgs(ID))
		}
	}
	return result
}

func (r Repository) GetAllCustomerIds() ([]string, int) {
	filter := bson.D{{}}
	res, _ := r.mc.Distinct(context.Background(), "_id", filter)
	response, _ := json.Marshal(res)
	var resp []string
	json.Unmarshal(response, &resp)
	return resp, len(resp)

}
func (r Repository) GetCustomers(options *options.FindOptions, filter *bson.M) ([]EntityModels.Customer, int) {
	var results []EntityModels.Customer
	cur, _ := r.mc.Find(context.Background(), filter, options)
	defer cur.Close(context.Background())
	cur.All(context.Background(), &results)
	return results, len(results)
}

func (r Repository) CreateCustomer(c *EntityModels.Customer) string {
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	res, _ := r.mc.InsertOne(context.Background(), c)
	return res.InsertedID.(primitive.ObjectID).Hex()
}

func (r Repository) UpdateCustomer(c *EntityModels.Customer) *mongo.UpdateResult {
	c.UpdatedAt = time.Now()
	filter := bson.M{"_id": c.ID}
	res, _ := r.mc.ReplaceOne(context.Background(), filter, c)
	return res
}

func (r Repository) DeleteCustomer(ID string) *mongo.DeleteResult {
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": objID}
	res, _ := r.mc.DeleteOne(context.Background(), filter)
	return res
}

func (r Repository) CheckCustomerIfExist(tc *RequestModels.TokenCredentials) {
	objID, err := primitive.ObjectIDFromHex(tc.ID)
	filter := bson.M{"_id": objID, "email": tc.Email}
	var result bson.M
	if err = r.mc.FindOne(context.Background(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			panic(ErrorModels.InvalidCredentials.SetPublicDetail("Wrong e-mail or id."))
		}
	}

}
