package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/SyaibanAhmadRamadhan/gocatch/genv"

	"realworld-go/infra/db"
	"realworld-go/internal"
	"realworld-go/presentation/rapi"
)

func main() {
	err := genv.Initialize(genv.DefaultEnvLib, false)
	gcommon.PanicIfError(err)

	mClient, mdb := db.NewMongoDbClient()

	dependecy := internal.DependencyMongodb(mdb, mClient)
	api := rapi.NewPresenter(dependecy)

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, os.Interrupt)
	go func() {
		<-exitSignal
		fmt.Println("Interrupt signal received, existing...")

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		if err := mClient.Disconnect(ctx); err != nil {
			fmt.Printf("failed graceful shutdown: %v\n", err)
		}

		api.Closed(ctx)
	}()

	api.InitProviderAndStart("realworld-services")
}
