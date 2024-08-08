package entity

//import "github.com/taranttini/study/go/pos-go-expert/fundacao-21-desafio/internal/entity"

type Order struct {
	Id    string `json:"id"`
	Data  string `json:"data"`
	Items []Item `json:"items"`
}

type Item struct {
	Id          string `json:"id"`
	OrderId     string
	Description string  `json:"description"`
	Qty         int     `json:"qty"`
	Value       float64 `json:"value"`
}

type OrderRepositoryInterface interface {
	Create(data string) (Order, error)
	AddItem(orderId string, descrition string, qty int, value float64) (Item, error)
	FindAllIncludeItems() ([]Order, error)
	//Get(orderId string) Order
	//List() []Order
}

type ItemRepositoryInterface interface {
	Create(orderId string, description string, qty int, value float64) (*Item, error)
	FindAll() ([]Item, error)
	FindByOrderId(orderId string) ([]Item, error)
}
