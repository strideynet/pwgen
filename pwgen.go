package pwgen

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type settings struct {
	length    int
	lowercase int
	number    int
	special   int
	uppercase int
}

type Option func(*settings)

var (
	lowercaseCharacters = "abcdefghijklmnopqrstuvwxyz"
	uppercaseCharacters = strings.ToUpper(lowercaseCharacters)
	numberCharacters    = "0123456789"
	specialCharacters   = "@%!?*^&"
	allCharacters       = lowercaseCharacters + uppercaseCharacters + numberCharacters + specialCharacters
)

// WithLength specifies the length of the password to return.
func WithLength(l int) Option {
	return func(s *settings) {
		s.length = l
	}
}

// WithLowercaseCount specifies the minimum number of lowercase characters to include.
func WithLowercaseCount(l int) Option {
	return func(s *settings) {
		s.lowercase = l
	}
}

// WithUppercaseCount specifies the minimum number of uppercase characters to include.
func WithUppercaseCount(l int) Option {
	return func(s *settings) {
		s.uppercase = l
	}
}

// WithNumberCount specifies the minimum number of numbers to include.
func WithNumberCount(l int) Option {
	return func(s *settings) {
		s.number = l
	}
}

// WithSpecialCount sets the minimum number of special characters to include (@%!?*^&).
func WithSpecialCount(l int) Option {
	return func(s *settings) {
		s.special = l
	}
}

// pickRandomCharacters picks a count of random characters from a string
func pickRandomCharacters(set string, count int) string {
	str := ""
	for i := 0; i < count; i++ {
		str += string(set[rand.Intn(len(set))])
	}
	return str
}

// Generates creates a password as a string using the options specified. It may return an error when an invalid config is provided.
// By default, the password includes 8 characters, 2 lowercase, 2 uppercase, 2 numbers and 2 specials.
func Generate(opts ...Option) (string, error) {
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

	// ensure lengths are valid
	if s.length < 0 || s.lowercase < 0 || s.uppercase < 0 || s.number < 0 || s.special < 0 {
		return "", errors.New("a specified minimum character count cant be less than 0")
	}

	minimumLength := s.lowercase + s.number + s.special + s.uppercase
	if minimumLength > s.length {
		return "", errors.New("combined minimum lengths cannot exceed specified length")
	}

	// Produce an overall set of the characters in the password
	combinedSet := pickRandomCharacters(lowercaseCharacters, s.lowercase) +
		pickRandomCharacters(uppercaseCharacters, s.uppercase) +
		pickRandomCharacters(numberCharacters, s.number) +
		pickRandomCharacters(specialCharacters, s.special) +
		pickRandomCharacters(allCharacters, s.length-minimumLength) // Top up the character set with all characters to reach requested length

	// Shuffle the string so that there is not a predictable ordering to the characters
	runeSlice := []rune(combinedSet) // cast to rune slice, so we can actually swap characters
	rand.Shuffle(len(combinedSet), func(i, j int) {
		runeSlice[i], runeSlice[j] = runeSlice[j], runeSlice[i]
	})

	return string(runeSlice), nil
}
