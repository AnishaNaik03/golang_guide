package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("expected deck length of 20, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades,but got %v", d[0])
	}

	if d[len(d)-1] != "four of club" {
		t.Errorf("Expected last card of Four of Clubs,but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveTofile("_deckstring")

	loadedDeck := newDeckFromFile("_deckstring")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck,got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
