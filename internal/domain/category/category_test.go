package category_test

import (
	"strings"
	"testing"

	"github.com/Vagmacker/luzora-api/internal/domain/category"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestGivenAValidParamsWhenCallsNewCategoryThenShouldReturnIt(t *testing.T) {
	expectedName := faker.Word()
	c, err := category.New(expectedName)
	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.NotNil(t, c.Id)
	assert.NotNil(t, c.CreatedAt)
	assert.NotNil(t, c.UpdatedAt)
	assert.Equal(t, c.Name, expectedName)
}

func TestGivenAnInvalidNameWhenCallsNewCategoryThenShouldReturnErr(t *testing.T) {
	expectedErrMessage := "name should not be empty"
	c, err := category.New("")
	assert.Nil(t, c)
	assert.Error(t, err, expectedErrMessage)
}

func TestGivenAnInvalidNameLengthMoreThan255WhenCallsNewCategoryReturnErr(t *testing.T) {
	expectedErrMessage := "name should be between 3 and 255 characters"
	c, err := category.New(faker.Word() + strings.Repeat(faker.Word(), category.MaxLength))
	assert.Nil(t, c)
	assert.Error(t, err, expectedErrMessage)
}

func TestGivenAnInvalidNameLengthLessThan3WhenCallsNewCategoryReturnErr(t *testing.T) {
	expectedErrMessage := "name should be between 3 and 255 characters"
	c, err := category.New("aa")
	assert.Nil(t, c)
	assert.Error(t, err, expectedErrMessage)
}
