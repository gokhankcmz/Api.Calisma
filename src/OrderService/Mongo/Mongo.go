package Mongo

import (
	"Api.Calisma/src/OrderService/Constants"
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"
)

var once sync.Once

// type global

var Collection *mongo.Collection

func GetMongoSingletonCollection() *mongo.Collection {
	once.Do(func() {
		collection, _ := GetMongoDbCollection(Constants.DBName, Constants.CollectionName)
		Collection = collection
	})
	return Collection
}

func GetMongoDbConnection() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Constants.MongoConnectionDuration)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(Constants.MongoConnectionString))
	if err != nil {
		log.Fatal(err) //todo: Logger'ı dışarıdan al
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
