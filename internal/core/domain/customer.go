package domain

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID        string
	Name      string
	Email     Email
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCustomer(name, mail string) (*Customer, error) {
	now := time.Now()
	email, err := NewEmail(mail)
	if err != nil {
		return nil, err
	}
	return &Customer{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
