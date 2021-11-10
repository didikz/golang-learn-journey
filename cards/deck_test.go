package main

import (
	"testing"
)

// Errorf reference: https://golangbyexample.com/errorf-function-golang/
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("Expected deck length of 20, but got %d", len(d))
	}

	if d[0] != "Ace of Spade" {
		t.Errorf("Expected first card is Ace of Spade, but got %s", d[0])
	}

	if d[len(d)-1] != "Five of Clubs" {
		t.Errorf("Expected last card is Clubs of Five, but got %s", d[len(d)-1])
	}
}
