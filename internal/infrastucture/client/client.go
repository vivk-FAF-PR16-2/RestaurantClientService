package client

import (
	"github.com/vivk-FAF-PR16-2/RestaurantDinnerHall/internal/infrastucture/random"
	"time"
)

type IClient interface {
	Update()
}

type client struct {
	id     int
	status bool

	timer <-chan time.Time
}

func NewClient(id int) IClient {
	c := &client{
		id: id,
	}

	c.changeState(false)
	return c
}

func (c *client) Update() {
	select {
	case <-c.timer:
		if c.status {

		} else {
			c.changeState(!c.status)
		}
		return
	}
}

func (c *client) changeState(status bool) {
	c.status = status
	if c.status {
		c.free()
	} else {
		c.order()
	}
}

func (c *client) free() {
	// TODO: Change source of `min` and `max` and `timeUnit` values to config
	min := 2
	max := 4

	timeUnit := time.Millisecond * 100

	timeValue := time.Duration(random.Range(min, max)) * timeUnit
	c.timer = time.After(timeValue)
}

func (c *client) order() {
	// TODO: Get menu from food ordering

	// TODO: Generate and send `order` to food ordering

	// TODO: Get `order` info and wait
}
