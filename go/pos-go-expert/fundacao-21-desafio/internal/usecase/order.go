package usecase

import (
	"github.com/taranttini/study/go/pos-go-expert/fundacao-21-desafio/internal/entity"
)

type OrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	ItemRepository  entity.ItemRepositoryInterface
}

func NewOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	ItemRepository entity.ItemRepositoryInterface,
) *OrderUseCase {
	return &OrderUseCase{
		OrderRepository: OrderRepository,
		ItemRepository:  ItemRepository,
	}
}

func (c *OrderUseCase) Create(data string) (*entity.Order, error) {

	order, err := c.OrderRepository.Create(data)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (c *OrderUseCase) AddItem(orderId string, descrition string, qty int, value float64) (*entity.Item, error) {
	item, err := c.OrderRepository.AddItem(orderId, descrition, qty, value)

	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (c *OrderUseCase) GetOrders() (*[]entity.Order, error) {

	item, err := c.OrderRepository.FindAllIncludeItems()

	if err != nil {
		return nil, err
	}
	return &item, nil
}
