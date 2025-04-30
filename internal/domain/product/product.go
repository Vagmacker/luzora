package product

import (
	"time"

	"github.com/google/uuid"
)

// Product is an entity that represents a product
type Product struct {
	ID          string
	Name        string
	Description string
	Image       string
	Price       float64
	CostPrice   float64
	Stock       int64
	CategoryId  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func New(name, description, image string, price, costPrice float64, stock int64) *Product {
	now := time.Now()
	return &Product{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Image:       image,
		Price:       price,
		CostPrice:   costPrice,
		Stock:       stock,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
