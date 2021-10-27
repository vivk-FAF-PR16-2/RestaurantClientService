package client

import (
	"github.com/spf13/viper"
	"github.com/vivk-FAF-PR16-2/RestaurantDinnerHall/internal/domain/dto"
	"github.com/vivk-FAF-PR16-2/RestaurantDinnerHall/internal/infrastucture/random"
	"time"
)

func GenerateOrder(id int, menu dto.MenuData) *dto.ClientInData {
	clientData := dto.ClientInData{
		ClientId: id,
		Orders:   make([]dto.OrderInData, 0),
	}

	for _, restaurant := range menu.RestaurantsData {
		min := viper.GetInt("min_order_items")
		max := viper.GetInt("min_order_items")
		count := random.Range(min, max)

		orderInData := dto.OrderInData{
			RestaurantId: restaurant.RestaurantId,
			Items:        make([]int, count),
			Priority:     random.Range(0, viper.GetInt("max_priority")-1),
		}

		items := restaurant.Menu
		itemsCount := len(items)
		maxWait := -1
		for i := 0; i < count; i++ {
			valueIndex := random.Range(0, itemsCount-1)
			orderInData.Items[i] = valueIndex

			item := items[valueIndex]
			if maxWait < item.PreparationTime {
				maxWait = item.PreparationTime
			}
		}

		orderInData.MaxWait = int(float64(maxWait) * viper.GetFloat64("max_wait_mult"))
		orderInData.CreatedTime = time.Now().Unix()

		clientData.Orders = append(clientData.Orders, orderInData)
	}

	return &clientData
}
