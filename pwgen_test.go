package pwgen

import (
	"fmt"
	"testing"
)

func Test_characterSets(t *testing.T) {
	// run a bunch of sanity checks against the character sets, so we don't fall a letter short
	t.Run("lowercaseCharacters", func(t *testing.T) {
		if len(lowercaseCharacters) != 26 {
			t.Fatalf("expected lowercase letters to have 36 characters, was %d", len(lowercaseCharacters))
		}
	})
	t.Run("numbers", func(t *testing.T) {
		if len(numberCharacters) != 10 {
			t.Fatalf("expected numbers to have 10 characters, was %d", len(numberCharacters))
		}
	})
}

func Test_pickRandomCharacters(t *testing.T) {
	expectLength := 10
	str := pickRandomCharacters(allCharacters, expectLength)
	if len(str) != expectLength {
		t.Fatalf("expected generated string to have %d characters, was %d", expectLength, len(str))
	}

	secondStr := pickRandomCharacters(allCharacters, expectLength)
	if secondStr == str {
		t.Fatalf("expected second generated string to differ from first")
	}
}

func TestGenerate(t *testing.T) {
	tt := []struct {
		Name                     string
		Options                  []Option
		ExpectedError            string
		ExpectedLength           int
		ExpectedMinimumLowercase int
		ExpectedMinimumUppercase int
		ExpectedMinimumSpecial   int
		ExpectedMinimumNumbers   int
	}{
		{
			Name:                     "defaults",
			ExpectedLength:           8,
			ExpectedMinimumLowercase: 2,
			ExpectedMinimumUppercase: 2,
			ExpectedMinimumNumbers:   2,
			ExpectedMinimumSpecial:   2,
		},
		{
			Name:          "too long",
			Options:       []Option{Length(2), Lowercase(16)},
			ExpectedError: "combined minimum lengths cannot exceed specified length",
		},
		{
			Name: "hail mary",
			Options: []Option{
				Length(33), // one over, so that we get one random letter
				Lowercase(8),
				Uppercase(8),
				Number(8),
				Special(8),
			},
		},
	}

	for _, te := range tt {
		t.Run(te.Name, func(t *testing.T) {
			str, err := Generate(te.Options...)
			// Ensure returned error matches expectations
			if err != nil {
				if te.ExpectedError == "" {
					t.Fatalf("No error expected, got '%s'", err)
				} else if te.ExpectedError != err.Error() {
					t.Fatalf("Expected error '%s', got '%s'", te.ExpectedError, err)
				}
			} else {
				if te.ExpectedError != "" {
					t.Fatalf("Expected error, got none")
				}
			}

			// Ensure correct number of characters are included

		})
	}
}
