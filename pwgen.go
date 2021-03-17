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

type option func(*settings)

var (
	lowercaseCharacters = "abcdefghijklmnopqrstuvwxyz"
	uppercaseCharacters = strings.ToUpper(lowercaseCharacters)
	numberCharacters    = "0123456789"
	specialCharacters   = "@%!?*^&"
	allCharacters       = lowercaseCharacters + uppercaseCharacters + numberCharacters + specialCharacters
)

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

// pickRandomCharacters picks a count of random characters from a string
func pickRandomCharacters(set string, count int) string {
	str := ""
	for i := 0; i < count; i++ {
		str += string(set[rand.Intn(len(set))])
	}
	return str
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

	combinedSet := pickRandomCharacters(lowercaseCharacters, s.lowercase) +
		pickRandomCharacters(uppercaseCharacters, s.uppercase) +
		pickRandomCharacters(numberCharacters, s.number) +
		pickRandomCharacters(specialCharacters, s.special) +
		pickRandomCharacters(allCharacters, s.length-minimumLength) // Top up the character set with all characters to reach requested length

	runeSlice := []rune(combinedSet)
	rand.Shuffle(len(combinedSet), func(i, j int) {
		runeSlice[i], runeSlice[j] = runeSlice[j], runeSlice[i]
	})

	return string(runeSlice), nil
}
