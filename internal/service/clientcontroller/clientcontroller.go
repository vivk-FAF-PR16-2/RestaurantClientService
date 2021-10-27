package clientcontroller

import (
	"context"
	"github.com/vivk-FAF-PR16-2/RestaurantDinnerHall/internal/infrastucture/client"
	"github.com/vivk-FAF-PR16-2/RestaurantDinnerHall/internal/service"
	"log"
)

type clientControllerService struct {
	clients []client.IClient
}

func NewService(ctx context.Context) service.IService {
	controller := &clientControllerService{}

	controller.Start(ctx)
	return controller
}

func (s *clientControllerService) Start(ctx context.Context) {

	count := 3 // TODO: Add value in config file

	s.clients = make([]client.IClient, count)
	for i := 0; i < count; i++ {
		s.clients[i] = client.NewClient(i)
	}

	go s.update(ctx)
}

func (s *clientControllerService) update(ctx context.Context) {
	for {

		select {
		case <-ctx.Done():
			log.Println("Stopping service listening")
			return
		}
	}
}
