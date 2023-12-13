package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/brianvoe/gofakeit/v6"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	// err := genv.Initialize(genv.DefaultEnvLib, false)
	// gcommon.PanicIfError(err)
	//
	// mongoConf := conf.EnvMongodb()
	// mClient, err := gmongodb.OpenConnMongoClient(mongoConf.URI())
	// gcommon.PanicIfError(err)
	//
	// go graceFullShutdown(mClient)
	//
	// time.Sleep(10 * time.Second)
	fmt.Println(gofakeit.Paragraph(3, 2, 10, "\n"))
}

func graceFullShutdown(mClient *mongo.Client) {
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, os.Interrupt)
	go func() {
		<-exitSignal
		fmt.Println("Interrupt signal received, existing...")

		if err := mClient.Disconnect(context.Background()); err != nil {
			fmt.Printf("failed graceful shutdown: %v\n", err)
		}
		fmt.Println("graceful shutdown")
	}()
}
