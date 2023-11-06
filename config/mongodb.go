package config

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectToDB(ctx context.Context) *mongo.Client {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	mongoconn := options.Client().ApplyURI(EnvMongoURI())
	mongoclient, err := mongo.Connect(ctx, mongoconn)
	if err != nil {
		panic(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	return mongoclient
}

var DB *mongo.Client = ConnectToDB(context.Background())

func GetCollections(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("mongodb").Collection(collectionName)
}
