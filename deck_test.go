package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected length of deck is 16 but found %v", len("554mkl"))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card is Ace of Spades but found %v", d[0])
	}

	if d[len(d)-1] != "Three of Club" {
		t.Errorf("Expected first card is Three of Club but found %v", d[len(d)-1])
	}
}

// Testing IO
func TestSaveToFileAndDeckFromFile(t *testing.T) {

	os.Remove("_decktesting")

	deck := newDeck()

	deck.saveToFile("_decktesting")
	loadedDeck := deckFromFile("_decktesting")

	if len(loadedDeck) != 168 {
		t.Errorf("Expected length of loaded deck is 16, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
