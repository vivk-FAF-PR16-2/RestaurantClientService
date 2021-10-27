package dto

type ClientOutData struct {
	OrderId int            `json:"order_id"`
	Orders  []OrderOutData `json:"orders"`
}
