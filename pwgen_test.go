package pwgen

import (
	"strings"
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

// counts number of letters in string that belong to a set
func countSetInString(set string, str string) int {
	count := 0
	for _, l := range str {
		if strings.Contains(set, string(l)) {
			count++
		}
	}
	return count
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
			Options:       []Option{WithLength(2), WithLowercaseCount(16)},
			ExpectedError: "combined minimum lengths cannot exceed specified length",
		},
		{
			Name:          "too short",
			Options:       []Option{WithLength(-12)},
			ExpectedError: "a specified minimum character count cant be less than 0",
		},
		{
			Name: "hail mary",
			Options: []Option{
				WithLength(33), // one over, so that we get one random letter
				WithLowercaseCount(8),
				WithUppercaseCount(8),
				WithNumberCount(8),
				WithSpecialCount(8),
			},
			ExpectedLength:           33,
			ExpectedMinimumLowercase: 8,
			ExpectedMinimumUppercase: 8,
			ExpectedMinimumNumbers:   8,
			ExpectedMinimumSpecial:   8,
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
				return
			} else {
				if te.ExpectedError != "" {
					t.Fatalf("Expected error, got none")
				}
			}

			// Ensure correct number of characters are included
			if te.ExpectedLength != len(str) {
				t.Errorf("expected length %d, got %d", te.ExpectedLength, len(str))
			}

			lowerCount := countSetInString(lowercaseCharacters, str)
			if lowerCount < te.ExpectedMinimumLowercase {
				t.Errorf("expected at least %d lowercase characters, got %d", te.ExpectedMinimumLowercase, lowerCount)
			}

			upperCount := countSetInString(uppercaseCharacters, str)
			if upperCount < te.ExpectedMinimumUppercase {
				t.Errorf("expected at least %d uppercase characters, got %d", te.ExpectedMinimumUppercase, upperCount)
			}

			specialCount := countSetInString(specialCharacters, str)
			if specialCount < te.ExpectedMinimumSpecial {
				t.Errorf("expected at least %d specialcase characters, got %d", te.ExpectedMinimumSpecial, specialCount)
			}

			numberCount := countSetInString(numberCharacters, str)
			if numberCount < te.ExpectedMinimumNumbers {
				t.Errorf("expected at least %d number characters, got %d", te.ExpectedMinimumNumbers, numberCount)
			}
		})
	}
}
