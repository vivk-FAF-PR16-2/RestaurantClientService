package main

import (
	"context"
	"github.com/vivk-FAF-PR16-2/RestaurantDinnerHall/application"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	isDone := make(chan os.Signal)
	signal.Notify(isDone, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	mainApp := application.New(ctx)
	go mainApp.Start()

	<-isDone
	cancel()
	mainApp.Shutdown()
}
