package domain

import (
	"time"

	"github.com/google/uuid"
)

// Category is an entity that represents a category of product
type Category struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCategory(name string) *Category {
	now := time.Now()
	return &Category{
		ID:        uuid.New().String(),
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
