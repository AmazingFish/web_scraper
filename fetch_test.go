package deckMaker

import (
	"testing"
)

func TestGetMeaning(t *testing.T) {
	meaning, err := getMeaning("professor")
	if err != nil {
		t.Fatalf("Error while getting the meaning for professor. error: %v", err)
	}
	t.Logf("Professor = %v", meaning)
}
