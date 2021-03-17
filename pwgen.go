package pwgen

import (
	"errors"
)

type settings struct {
	length    int
	lowercase int
	number    int
	special   int
	uppercase int
}

type option func(*settings)

// Length specifies the length of the password to return.
func Length(l int) option {
	return func(s *settings) {
		s.length = l
	}
}

// Lowercase specifies the minimum number of lowercase characters to include.
func Lowercase(l int) option {
	return func(s *settings) {
		s.lowercase = l
	}
}

// Uppercase specifies the minimum number of uppercase characters to include.
func Uppercase(l int) option {
	return func(s *settings) {
		s.uppercase = l
	}
}

// Number specifies the minimum number of numbers to include.
func Number(l int) option {
	return func(s *settings) {
		s.number = l
	}
}

// Special sets the minimum number of special characters to include (@%!?*^&).
func Special(l int) option {
	return func(s *settings) {
		s.special = l
	}
}

func Generate(opts ...option) (string, error) {
	// start with default settings values
	s := &settings{
		length:    8,
		uppercase: 2,
		lowercase: 2,
		number:    2,
		special:   2,
	}
	// apply specified options to settings
	for _, opt := range opts {
		opt(s)
	}

	minimumLength := s.lowercase + s.number + s.special + s.uppercase
	if minimumLength > s.length {
		return "", errors.New("combined minimum lengths cannot exceed specified length")
	}

	return "", nil
}
