package coupon

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	ErrCodeIsNotEmpty = errors.New("")
)

type Coupon struct {
	ID        string
	Code      string
	Value     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(code string, value float64) (*Coupon, error) {
	now := time.Now()

	c := &Coupon{
		ID:        uuid.New().String(),
		Code:      code,
		Value:     value,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := c.Validate(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Coupon) Validate() error {
	if strings.TrimSpace(c.Code) == "" {
		return ErrCodeIsNotEmpty
	}
	return nil
}
