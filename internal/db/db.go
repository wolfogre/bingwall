package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"bingwall/internal/entity"
)

const (
	databaseName   = "bingwall"
	collectionName = "history"
)

var (
	mongoClient *mongo.Client
)

func Init(mongoUrl string) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl))
	if err != nil {
		return err
	}
	if err := client.Connect(context.Background()); err != nil {
		return err
	}
	if err := client.Ping(context.Background(), nil); err != nil {
		return err
	}

	mongoClient = client
	return nil
}

func InsertHistory(history entity.History) error {
	collect := mongoClient.Database(databaseName).Collection(collectionName)
	_, err := collect.InsertOne(context.Background(), history)
	return err
}

func FindHistory(id string) (*entity.History, error) {
	collect := mongoClient.Database(databaseName).Collection(collectionName)
	ret := entity.History{}
	err := collect.FindOne(context.Background(), bson.M{
		"_id": id,
	}).Decode(&ret)
	return &ret, err
}

func ExistHistory(id string) (bool, error) {
	collect := mongoClient.Database(databaseName).Collection(collectionName)
	count, err := collect.CountDocuments(context.Background(), bson.M{
		"_id": id,
	})
	return count > 0, err
}
