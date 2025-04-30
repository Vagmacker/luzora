package email

import (
	"errors"
	"fmt"
	"net/mail"
	"regexp"
	"strings"
	"unicode/utf8"
)

const EmailMaxLength = 100

var (
	InvalidEmailChars = regexp.MustCompile(`[^a-zA-Z0-9+.@_~\-]`)
	ValidEmailSeq     = regexp.MustCompile(`^[a-zA-Z0-9+._~\-]+@[a-zA-Z0-9+._~\-]+(\.[a-zA-Z0-9+._~\-]+)+$`)
)

type Email struct {
	value string
}

func New(email string) (Email, error) {
	if strings.TrimSpace(email) == "" {
		return Email{}, errors.New("cannot be empty")
	}

	if strings.ContainsAny(email, " \t\n\r") {
		return Email{}, errors.New("cannot contain whitespace")
	}
	if strings.ContainsAny(email, `"'`) {
		return Email{}, errors.New("cannot contain quotes")
	}

	if rc := utf8.RuneCountInString(email); rc > EmailMaxLength {
		return Email{}, fmt.Errorf("cannot be a over %v characters in length", EmailMaxLength)
	}

	addr, err := mail.ParseAddress(email)
	if err != nil {
		email = strings.TrimSpace(email)
		msg := strings.TrimPrefix(strings.ToLower(err.Error()), "mail: ")

		switch {
		case strings.Contains(msg, "missing '@'"):
			return Email{}, errors.New("missing the @ sign")

		case strings.HasPrefix(email, "@"):
			return Email{}, errors.New("missing part before the @ sign")

		case strings.HasSuffix(email, "@"):
			return Email{}, errors.New("missing part after the @ sign")
		}

		return Email{}, errors.New(msg)
	}

	if addr.Name != "" {
		return Email{}, errors.New("cannot not include a name")
	}

	if matches := InvalidEmailChars.FindAllString(addr.Address, -1); len(matches) != 0 {
		return Email{}, fmt.Errorf("cannot contain: %v", matches)
	}

	if !ValidEmailSeq.MatchString(addr.Address) {
		_, end, _ := strings.Cut(addr.Address, "@")
		if !strings.Contains(end, ".") {
			return Email{}, errors.New("missing top-level domain, e.g. .com, .co.uk, etc")
		}

		return Email{}, errors.New("must be an email address, e.g. email@example.com")
	}

	return Email{value: addr.Address}, nil
}

func (e Email) Value() string {
	return e.value
}
