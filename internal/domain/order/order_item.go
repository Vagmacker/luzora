package order

import "time"

// OrderItem is an entity that represents pivot table between order and product
type OrderItem struct {
	ID         string
	OrderID    string
	ProductID  string
	Quantity   int64
	TotalPrice float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
