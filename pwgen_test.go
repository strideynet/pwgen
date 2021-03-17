package pwgen

import "testing"

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
