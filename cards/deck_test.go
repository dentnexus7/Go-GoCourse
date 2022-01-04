package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, bot got %v", d[(len(d)-1)])
	}

}

func TestSaveToDeckAndNewDecktestFromFile(t *testing.T) {
	// Remove initial files
	os.Remove("_decktesting")

	// Create new deck
	deck := newDeck()
	// Save deck to hard drive file
	deck.saveToFile("_decktesting")

	// Create deck from hard drive file
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	// Remove used files
	os.Remove("_decktesting")
}
