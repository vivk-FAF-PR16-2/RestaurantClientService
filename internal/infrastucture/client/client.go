package client

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"github.com/vivk-FAF-PR16-2/RestaurantDinnerHall/internal/domain/dto"
	"github.com/vivk-FAF-PR16-2/RestaurantDinnerHall/internal/http/sendrequest"
	"log"
)

type IClient interface {
	Update()
}

type client struct {
	id     int
	status bool

	threads []IThread
}

func NewClient(id int) IClient {
	c := &client{
		id: id,
	}

	c.order()
	return c
}

func (c *client) Update() {
	max := len(c.threads)
	count := 0
	for _, thread := range c.threads {
		thread.Update()
		if thread.GetReadyStatus() {
			count++
		}
	}

	if max == count {
		// TODO: Finish...
	}
}

func (c *client) order() {
	var menu dto.MenuData
	// TODO: Get menu from food ordering
	menu = dto.MenuData{}

	clientData := GenerateOrder(c.id, menu)
	jsonData, err := json.Marshal(clientData)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	addr := fmt.Sprintf("http://%s/order", viper.GetString("food_ordering_addr"))
	response := sendrequest.Post(addr, jsonData)

	var responseData dto.ClientOutData
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	c.threads = make([]IThread, len(responseData.Orders))
	for i, order := range responseData.Orders {
		c.threads[i] = NewThread(order)
	}
}
