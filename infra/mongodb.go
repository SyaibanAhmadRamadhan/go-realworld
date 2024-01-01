package infra

import (
	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb/gmongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"

	"realworld-go/conf"
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewMongoDbClient() *MongoDB {
	mongoConf := conf.LoadEnvMongodb()

	clientOpt := options.Client()
	clientOpt.ApplyURI(mongoConf.URI())
	clientOpt.Monitor = otelmongo.NewMonitor(otelmongo.WithCommandAttributeDisabled(false))

	mClient, err := gmongodb.OpenConnMongoClient(clientOpt)
	gcommon.PanicIfError(err)

	return &MongoDB{
		Client:   mClient,
		Database: mClient.Database(mongoConf.Database),
	}
}
