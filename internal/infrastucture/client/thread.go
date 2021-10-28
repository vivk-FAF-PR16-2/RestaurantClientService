package client

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"github.com/vivk-FAF-PR16-2/RestaurantDinnerHall/internal/domain/dto"
	"github.com/vivk-FAF-PR16-2/RestaurantDinnerHall/internal/http/sendrequest"
	"log"
	"time"
)

type IThread interface {
	GetReadyStatus() bool
	Update()
}

type orderThread struct {
	ready bool

	addr string
	id   int

	order dto.OrderStatusData
	timer <-chan time.Time
}

func NewThread(data dto.OrderOutData) IThread {
	order := &orderThread{
		addr: data.RestaurantAddress,
		id:   data.OrderId,

		ready: false,
	}

	order.Reset()
	return order
}

func (t *orderThread) GetReadyStatus() bool {
	return t.ready
}

func (t *orderThread) Reset() {
	var order dto.OrderStatusData

	addrOrderStatus := fmt.Sprintf("http://%s/v2/order/%d", t.addr, t.id)
	responseOrderStatus := sendrequest.Get(addrOrderStatus, nil)
	err := json.NewDecoder(responseOrderStatus.Body).Decode(&order)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if !order.IsReady {
		t.ready = false
		t.order = order
		t.Calculate()
		return
	}

	t.ready = true
}

func (t *orderThread) Calculate() {
	mult := viper.GetInt("time_unit_mult")
	timeUnit := time.Millisecond * time.Duration(mult)

	timeValue := time.Duration(t.order.EstimatedWaitingTime)

	t.timer = time.After(timeValue * timeUnit)
}

func (t *orderThread) Update() {
	if t.ready {
		return
	}

	select {
	case <-t.timer:
		t.Reset()
		return
	default:
		return
	}
}
