package dto

import "github.com/vivk-FAF-PR16-2/RestaurantDinnerHall/internal/domain/entity"

type OrderStatusData struct {
	OrderId              int                    `json:"order_id"`
	IsReady              bool                   `json:"is_ready"`
	EstimatedWaitingTime int                    `json:"estimated_waiting_time"`
	Priority             int                    `json:"priority"`
	MaxWait              int                    `json:"max_wait"`
	CreatedTime          int64                  `json:"created_time"`
	RegisteredTime       int64                  `json:"registered_time"`
	PreparedTime         int64                  `json:"prepared_time"`
	CookingTime          int                    `json:"cooking_time"`
	CookingDetails       []entity.CookingDetail `json:"cooking_details"`
}
