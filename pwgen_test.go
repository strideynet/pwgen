package pwgen

import "testing"

func TestCharacterSets(t *testing.T) {
	// run a bunch of sanity checks against the character sets, so we don't fall a letter short
	t.Run("lowercaseLetters", func(t *testing.T) {
		if len(lowercaseLetters) != 26 {
			t.Fatalf("expected lowercase letters to have 36 characters, was %d", len(lowercaseLetters))
		}
	})
	t.Run("numbers", func(t *testing.T) {
		if len(numbers) != 10 {
			t.Fatalf("expected numbers to have 10 characters, was %d", len(numbers))
		}
	})
}
