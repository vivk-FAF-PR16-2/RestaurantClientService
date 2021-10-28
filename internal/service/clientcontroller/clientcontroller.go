package clientcontroller

import (
	"context"
	"github.com/spf13/viper"
	"github.com/vivk-FAF-PR16-2/RestaurantDinnerHall/internal/infrastucture/client"
	"github.com/vivk-FAF-PR16-2/RestaurantDinnerHall/internal/infrastucture/idprovider"
	"github.com/vivk-FAF-PR16-2/RestaurantDinnerHall/internal/service"
	"log"
)

type clientControllerService struct {
	clients  []client.IClient
	provider idprovider.IProvider
}

func NewService(ctx context.Context) service.IService {
	controller := &clientControllerService{}

	controller.Start(ctx)
	return controller
}

func (s *clientControllerService) Start(ctx context.Context) {
	count := viper.GetInt("client_count")

	s.clients = make([]client.IClient, count)
	for i := 0; i < count; i++ {
		s.provider = idprovider.NewProvider()
		id := s.provider.Get()
		s.clients[i] = client.NewClient(id)
	}

	loop := func() {
		for {
			s.update(ctx)
		}
	}

	go loop()
}

func (s *clientControllerService) update(ctx context.Context) {
	for i := range s.clients {
		s.clients[i].Update()

		if s.clients[i].GetReadyStatus() {
			id := s.provider.Get()
			s.clients[i] = client.NewClient(id)
		}
	}

	select {
	case <-ctx.Done():
		log.Println("Stopping service listening")
		return
	}
}
