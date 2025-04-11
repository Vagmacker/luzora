package domain

import (
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestGivenAValidParamsWhenCallsNewCustomerShoudBeOK(t *testing.T) {
	expectedName := faker.Name()
	expectedEmail, errEmail := NewEmail(faker.Email())
	customer, err := NewCustomer(expectedName, expectedEmail.String())
	assert.Nil(t, errEmail)
	assert.Nil(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, expectedEmail.String(), customer.Email.String())
	assert.Equal(t, expectedName, customer.Name)
	assert.NotNil(t, customer.CreatedAt)
	assert.NotNil(t, customer.UpdatedAt)
}
