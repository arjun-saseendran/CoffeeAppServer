package coffee

import "time"

type Order struct {
	ID         uint      `gorm:"privateKey" json:"id"`
	Name       string    `json:"name"`
	CoffeeName string    `json:"coffeeName"`
	Size       string    `json:"size"`
	Total      float64   `json:"total"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func NewOrder() *Order {
	return &Order{}
}
