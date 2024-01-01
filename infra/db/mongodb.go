package db

import (
	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb/gmongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"

	"realworld-go/conf"
)

func NewMongoDbClient() (*mongo.Client, *mongo.Database) {
	mongoConf := conf.EnvMongodb()

	clientOpt := options.Client()
	clientOpt.ApplyURI(mongoConf.URI())
	clientOpt.Monitor = otelmongo.NewMonitor()

	mClient, err := gmongodb.OpenConnMongoClient(clientOpt)
	gcommon.PanicIfError(err)

	return mClient, mClient.Database(mongoConf.Database)
}
