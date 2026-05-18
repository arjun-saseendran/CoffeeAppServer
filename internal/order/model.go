package order

import "time"

type CoffeeSize string

const (
	Small CoffeeSize = "Small"
	Medium CoffeeSize = "Medium"
	Large CoffeeSize = "Large"
)

type Order struct {
	ID         uint      `gorm:"privateKey" json:"id"`
	Name       string    `json:"name"`
	CoffeeName string    `json:"coffeeName"`
	Size       CoffeeSize    `json:"size"`
	Total      float64   `json:"total"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func NewOrder() *Order {
	return &Order{}
}
