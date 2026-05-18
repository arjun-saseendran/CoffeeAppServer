package order

type InputCreateOrder struct {
	Name       string  `json:"name" binding:"required"`
	CoffeeName string  `json:"coffeeName" binding:"required"`
	Size       string  `json:"size" bindig:"required"`
	Total      float64 `json:"total" binding:"required"`
}

type InputUpdateOrder struct {
	Name       string `json:"name"`
	CoffeeName string `json:"coffeeName"`
	Size       string `json:"size"`
}

func NewInputCreateOrder() *InputCreateOrder {
	return &InputCreateOrder{}
}

func NewInputUpdateOrder() *InputUpdateOrder {
	return &InputUpdateOrder{}
}
