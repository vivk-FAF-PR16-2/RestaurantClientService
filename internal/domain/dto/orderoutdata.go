package dto

type OrderOutData struct {
	RestaurantId         int    `json:"restaurant_id"`
	RestaurantAddress    string `json:"restaurant_address"`
	OrderId              int    `json:"order_id"`
	EstimatedWaitingTime int    `json:"estimated_waiting_time"`
	CreatedTime          int64  `json:"created_time"`
	RegisteredTime       int64  `json:"registered_time"`
}
