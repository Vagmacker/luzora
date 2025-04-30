package category

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
	id        string
	name      string
	createdAt time.Time
	updatedAt time.Time
}

func New(name string) (*Category, error) {
	now := time.Now()
	c := &Category{
		id:        uuid.New().String(),
		name:      name,
		createdAt: now,
		updatedAt: now,
	}
	if err := c.Validate(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Category) Validate() error {
	if strings.TrimSpace(c.name) == "" {
		return ErrNameIsNotEmpty
	}
	if len(strings.TrimSpace(c.name)) < MinLength || len(strings.TrimSpace(c.name)) > MaxLength {
		return ErrNameSize
	}
	return nil
}

func (c *Category) Id() string {
	return c.id
}

func (c *Category) Name() string {
	return c.name
}

func (c *Category) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Category) UpdatedAt() time.Time {
	return c.updatedAt
}
