package model

type Item struct {
	ID          string  `json:"id"`
	Description string  `json:"description"`
	Qty         int     `json:"qty"`
	Value       float64 `json:"value"`
	Order       *Order  `json:"order"`
}
