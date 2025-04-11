package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAValidEmailWhenCallsNewEmailShoudBeOK(t *testing.T) {
	expectedEmail := "johndoe@gmail.com"
	email, err := NewEmail(expectedEmail)
	assert.Nil(t, err)
	assert.Equal(t, expectedEmail, email.String())
}
