package order

import "gorm.io/gorm"

type OrderService interface {
	CreateOrder(orderData *InputCreateOrder) (*Order, error)
	GetOrders() ([]Order, error)
	GetOrder(id string) (*Order, error)
	UpdateOrder(id string, orderData *InputUpdateOrder) (*Order, error)
	DeleteOrder(id string) error
}

type orderService struct {
	db *gorm.DB
}

func NewOrderService(db *gorm.DB) OrderService {
	return &orderService{db}
}

func (os *orderService) CreateOrder(orderData *InputCreateOrder) (*Order, error) {
	newOrder := &Order{Name: orderData.Name, CoffeeName: orderData.CoffeeName, Size: orderData.Size, Total: orderData.Total}
	result := os.db.Create(newOrder)
	if result.Error != nil {
		return nil, result.Error
	}
	return newOrder, nil

}

func (os *orderService) GetOrders() ([]Order, error) {
	var orders []Order
	result := os.db.Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

func (os *orderService) GetOrder(id string) (*Order, error) {
	order := NewOrder()
	result := os.db.First(&order, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return order, nil
}

func (os *orderService) UpdateOrder(id string, orderData *InputUpdateOrder) (*Order, error) {
	updateOrder := NewOrder()
	result := os.db.First(updateOrder, id)
	if result.Error != nil {
		return nil, result.Error
	}
	updateData := Order{
		Name:       orderData.Name,
		CoffeeName: orderData.CoffeeName,
	}
	if orderData.Size != nil {
		updateData.Size = *orderData.Size
	}
	result = os.db.Model(updateOrder).Updates(updateData)

	if result.Error != nil {
		return nil, result.Error
	}

	return updateOrder, nil
}

func (os *orderService) DeleteOrder(id string) error {
	deleteOrder := NewOrder()
	result := os.db.First(deleteOrder, id)
	if result.Error != nil {
		return result.Error
	}
	os.db.Delete(deleteOrder)
	return nil
}
