package customer_test

import (
	"testing"

	"github.com/Vagmacker/luzora-api/internal/domain/customer"
	"github.com/Vagmacker/luzora-api/internal/domain/customer/email"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestGivenAValidParamsWhenCallsNewCustomerShoudBeOK(t *testing.T) {
	// Given
	expectedName := faker.Name()
	expectedEmail, err := email.New(faker.Email())
	assert.Nil(t, err)

	// When
	customer, err := customer.New(expectedName, expectedEmail.Value())

	// Then
	assert.Nil(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, expectedEmail.Value(), customer.Email().Value())
	assert.Equal(t, expectedName, customer.Name())
	assert.NotNil(t, customer.CreatedAt)
	assert.NotNil(t, customer.UpdatedAt)
}
