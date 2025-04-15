package domain

import (
	"strings"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestGivenAValidParamsWhenCallsNewCategoryThenShouldReturnIt(t *testing.T) {
	expectedName := faker.Word()
	c, err := NewCategory(expectedName)
	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.NotNil(t, c.ID)
	assert.NotNil(t, c.CreatedAt)
	assert.NotNil(t, c.UpdatedAt)
	assert.Equal(t, c.Name, expectedName)
}

func TestGivenAnInvalidNameWhenCallsNewCategoryThenShouldReturnErr(t *testing.T) {
	expectedErrMessage := "name should not be empty"
	c, err := NewCategory("")
	assert.Nil(t, c)
	assert.Error(t, err, expectedErrMessage)
}

func TestGivenAnInvalidNameLengthMoreThan255WhenCallsNewCategoryReturnErr(t *testing.T) {
	expectedErrMessage := "name should be between 3 and 255 characters"
	c, err := NewCategory(faker.Word() + strings.Repeat(faker.Word(), EmailMaxLength))
	assert.Nil(t, c)
	assert.Error(t, err, expectedErrMessage)
}

func TestGivenAnInvalidNameLengthLessThan3WhenCallsNewCategoryReturnErr(t *testing.T) {
	expectedErrMessage := "name should be between 3 and 255 characters"
	c, err := NewCategory("aa")
	assert.Nil(t, c)
	assert.Error(t, err, expectedErrMessage)
}
