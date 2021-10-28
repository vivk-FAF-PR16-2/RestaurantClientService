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
	addr := viper.GetString("food_ordering_addr")
	addrMenu := fmt.Sprintf("http://%s/menu", addr)
	responseMenu := sendrequest.Get(addrMenu, nil)

	var menu dto.MenuData
	err := json.NewDecoder(responseMenu.Body).Decode(&menu)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	clientData := GenerateOrder(c.id, menu)
	jsonData, err := json.Marshal(clientData)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	addrOrder := fmt.Sprintf("http://%s/order", addr)
	responseOrder := sendrequest.Post(addrOrder, jsonData)

	var responseData dto.ClientOutData
	err = json.NewDecoder(responseOrder.Body).Decode(&responseData)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	c.threads = make([]IThread, len(responseData.Orders))
	for i, order := range responseData.Orders {
		c.threads[i] = NewThread(order)
	}
}
