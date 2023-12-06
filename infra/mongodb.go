package infra

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"realworld-go/conf"
)

func OpenConnMongoClient(conf conf.MongodbConf) (*mongo.Client, error) {
	clientOpt := options.Client()
	clientOpt.ApplyURI(conf.URI())

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpt)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, nil
}

func OpenConnMongoDB(conf conf.MongodbConf) (*mongo.Database, error) {
	client, err := OpenConnMongoClient(conf)
	if err != nil {
		return nil, err
	}

	db := client.Database(conf.Database)

	return db, nil
}
