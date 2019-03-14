package db

import (
	"bingwall/internal/entity"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

const (
	databaseName   = "bingwall"
	collectionName = "history"
)

var (
	m *sync.Mutex
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
	if err := client.Disconnect(context.Background()); err != nil {
		return err
	}

	m = &sync.Mutex{}
	mongoClient = client
	return nil
}

func InsertHistory(history entity.History) error {
	m.Lock()
	defer m.Unlock()
	if err := mongoClient.Connect(context.Background()); err != nil {
		return err
	}
	defer mongoClient.Disconnect(context.Background())

	collect := mongoClient.Database(databaseName).Collection(collectionName)
	_, err := collect.InsertOne(context.Background(), history)
	return err
}

func ExistHistory(id string) (bool, error) {
	m.Lock()
	defer m.Unlock()
	if err := mongoClient.Connect(context.Background()); err != nil {
		return false, err
	}
	defer mongoClient.Disconnect(context.Background())

	collect := mongoClient.Database(databaseName).Collection(collectionName)
	err := collect.FindOne(context.Background(), bson.M{
		"_id": id,
	}).Err()
	if err == nil {
		return true, nil
	}
	if err == mongo.ErrNoDocuments {
		return false, nil
	}
	return false, err
}
