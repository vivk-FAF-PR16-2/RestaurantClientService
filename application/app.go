package application

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/vivk-FAF-PR16-2/RestaurantDinnerHall/internal/service"
	"github.com/vivk-FAF-PR16-2/RestaurantDinnerHall/internal/service/clientcontroller"
	"log"
	"net/http"
)

type IApp interface {
	Start()
	Shutdown()
}

type clientApp struct {
	server  *http.Server
	service service.IService
}

func New(ctx context.Context) IApp {
	router := gin.New()

	clientService := clientcontroller.NewService(ctx)

	return &clientApp{
		server: &http.Server{
			Addr:    ":56567",
			Handler: router,
		},
		service: clientService,
	}
}

func (app *clientApp) Start() {
	if err := app.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error while running `food ordering` server: %v\n", err)
	}
}

func (app *clientApp) Shutdown() {
	if err := app.server.Shutdown(context.Background()); err != nil {
		log.Fatalf("Unable to shutdown `food ordering` server: %v\n", err)
	}
}
