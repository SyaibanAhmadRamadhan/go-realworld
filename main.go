package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/SyaibanAhmadRamadhan/gocatch/genv"

	"realworld-go/conf"
	"realworld-go/infra"
)

func main() {
	err := genv.Initialize(genv.DefaultEnvLib, false)
	gcommon.PanicIfError(err)

	mongoConf := conf.EnvMongodb()

	db, err := infra.OpenConnMongoDB(*mongoConf)
	if err != nil {
		err := fmt.Errorf("failed to open connection to mongodb: %v", err)
		fmt.Println(err)
		os.Exit(1)
	}

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, os.Interrupt)
	go func() {
		<-exitSignal
		fmt.Println("Interrupt signal received, existing...")

		if err := db.Client().Disconnect(nil); err != nil {
			fmt.Printf("failed graceful shutdown: %v\n", err)
		}
		fmt.Println("graceful shutdown")
	}()

	time.Sleep(10 * time.Second)
}
