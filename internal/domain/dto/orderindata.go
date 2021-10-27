package dto

type OrderInData struct {
	RestaurantId int   `json:"restaurant_id"`
	Items        []int `json:"items"`
	Priority     int   `json:"priority"`
	MaxWait      int   `json:"max_wait"`
	CreatedTime  int64 `json:"created_time"`
}
