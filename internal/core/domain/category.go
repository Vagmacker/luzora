package domain

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	MinLength = 3
	MaxLength = 255
)

var (
	ErrNameIsNotEmpty = errors.New("name should not be empty")
	ErrNameSize       = errors.New("name should be between 3 and 255 characters")
)

// Category is an entity that represents a category of product
type Category struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCategory(name string) (*Category, error) {
	now := time.Now()
	c := &Category{
		ID:        uuid.New().String(),
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err := c.Validate()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Category) Validate() error {
	if strings.TrimSpace(c.Name) == "" {
		return ErrNameIsNotEmpty
	}
	if len(strings.TrimSpace(c.Name)) < MinLength || len(strings.TrimSpace(c.Name)) > MaxLength {
		return ErrNameSize
	}
	return nil
}
