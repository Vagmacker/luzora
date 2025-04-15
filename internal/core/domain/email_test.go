package domain

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAValidEmailWhenCallsNewEmailShoudBeOK(t *testing.T) {
	expectedEmail := "johndoe@gmail.com"
	email, err := NewEmail(expectedEmail)
	assert.Nil(t, err)
	assert.Equal(t, expectedEmail, email.String())
}

func TestGivenAnEmptyEmailWhenCallsNewEmailShouldReturnError(t *testing.T) {
	email, err := NewEmail("")
	assert.Error(t, err)
	assert.Equal(t, "", email.String())
	assert.Equal(t, "cannot be empty", err.Error())
}

func TestGivenAnEmailWithWhitespaceWhenCallsNewEmailShouldReturnError(t *testing.T) {
	email, err := NewEmail("john doe@gmail.com")
	assert.Error(t, err)
	assert.Equal(t, "", email.String())
	assert.Equal(t, "cannot contain whitespace", err.Error())
}

func TestGivenAnEmailWithQuotesWhenCallsNewEmailShouldReturnError(t *testing.T) {
	email, err := NewEmail(`"john"@gmail.com`)
	assert.Error(t, err)
	assert.Equal(t, "", email.String())
	assert.Equal(t, "cannot contain quotes", err.Error())
}

func TestGivenAnEmailExceedingMaxLengthWhenCallsNewEmailShouldReturnError(t *testing.T) {
	longEmail := "a" + strings.Repeat("b", EmailMaxLength) + "@gmail.com"
	email, err := NewEmail(longEmail)
	assert.Error(t, err)
	assert.Equal(t, "", email.String())
	assert.Contains(t, err.Error(), "cannot be a over")
}

func TestGivenAnEmailWithoutAtSignWhenCallsNewEmailShouldReturnError(t *testing.T) {
	email, err := NewEmail("johndoegmail.com")
	assert.Error(t, err)
	assert.Equal(t, "", email.String())
	assert.Equal(t, "missing the @ sign", err.Error())
}

func TestGivenAnEmailStartingWithAtSignWhenCallsNewEmailShouldReturnError(t *testing.T) {
	email, err := NewEmail("@gmail.com")
	assert.Error(t, err)
	assert.Equal(t, "", email.String())
	assert.Equal(t, "missing part before the @ sign", err.Error())
}

func TestGivenAnEmailEndingWithAtSignWhenCallsNewEmailShouldReturnError(t *testing.T) {
	email, err := NewEmail("johndoe@")
	assert.Error(t, err)
	assert.Equal(t, "", email.String())
	assert.Equal(t, "missing part after the @ sign", err.Error())
}

func TestGivenAnEmailWithNameWhenCallsNewEmailShouldReturnError(t *testing.T) {
	email, err := NewEmail("John Doe <johndoe@gmail.com>")
	assert.Error(t, err)
	assert.Equal(t, "", email.String())
	assert.Equal(t, "cannot not include a name", err.Error())
}

func TestGivenAnEmailWithInvalidCharsWhenCallsNewEmailShouldReturnError(t *testing.T) {
	email, err := NewEmail("john#doe@gmail.com")
	assert.Error(t, err)
	assert.Equal(t, "", email.String())
	assert.Contains(t, err.Error(), "cannot contain:")
}

func TestGivenAnEmailWithoutTopLevelDomainWhenCallsNewEmailShouldReturnError(t *testing.T) {
	email, err := NewEmail("johndoe@gmail")
	assert.Error(t, err)
	assert.Equal(t, "", email.String())
	assert.Equal(t, "missing top-level domain, e.g. .com, .co.uk, etc", err.Error())
}

func TestGivenAnInvalidEmailFormatWhenCallsNewEmailShouldReturnError(t *testing.T) {
	email, err := NewEmail("johndoe@.com")
	assert.Error(t, err)
	assert.Equal(t, "", email.String())
	assert.Equal(t, "must be an email address, e.g. email@example.com", err.Error())
}

func TestGivenAValidEmailWithDifferentCasesWhenCallsNewEmailShouldBeOK(t *testing.T) {
	testCases := []string{
		"john.doe@gmail.com",
		"john+doe@gmail.com",
		"john_doe@gmail.com",
		"john-doe@gmail.com",
		"john~doe@gmail.com",
		"johndoe@sub.domain.com",
		"johndoe@domain.co.uk",
	}

	for _, tc := range testCases {
		t.Run(tc, func(t *testing.T) {
			email, err := NewEmail(tc)
			assert.Nil(t, err)
			assert.Equal(t, tc, email.String())
		})
	}
}
