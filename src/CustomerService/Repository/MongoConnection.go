package CustomerRepository

import (
	"Api.Calisma/src/CustomerService/Constants"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"sync"
)

var once sync.Once

// type global

var MongoCollection *mongo.Collection

func GetMongoSingletonCollection() *mongo.Collection {
	once.Do(func() {
		collection, _ := GetMongoDbCollection(Constants.DBName, Constants.CollectionName)
		MongoCollection = collection
	})
	return MongoCollection
}


func GetMongoDbConnection() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Constants.MongoConnectionDuration)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(Constants.MongoConnectionString))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	return client, nil
}

func GetMongoDbCollection(DbName string, CollectionName string) (*mongo.Collection, error) {
	client, err := GetMongoDbConnection()
	if err != nil {
		return nil, err
	}
	collection := client.Database(DbName).Collection(CollectionName)
	return collection, nil
}
