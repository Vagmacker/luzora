package customer

import (
	"errors"
	"strings"
	"time"

	"github.com/Vagmacker/luzora-api/internal/domain/customer/email"
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

type Customer struct {
	id        string
	name      string
	email     email.Email
	createdAt time.Time
	updatedAt time.Time
}

func New(name, mail string) (*Customer, error) {
	now := time.Now()
	email, err := email.New(mail)
	if err != nil {
		return nil, err
	}
	return &Customer{
		id:        uuid.New().String(),
		name:      name,
		email:     email,
		createdAt: now,
		updatedAt: now,
	}, nil
}

func (c *Customer) Validate() error {
	if strings.TrimSpace(c.name) == "" {
		return ErrNameIsNotEmpty
	}
	if len(strings.TrimSpace(c.name)) < MinLength || len(strings.TrimSpace(c.name)) > MaxLength {
		return ErrNameSize
	}
	return nil
}

func (c *Customer) Id() string {
	return c.id
}

func (c *Customer) Name() string {
	return c.name
}

func (c *Customer) Email() email.Email {
	return c.email
}

func (c *Customer) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Customer) UpdatedAt() time.Time {
	return c.updatedAt
}
