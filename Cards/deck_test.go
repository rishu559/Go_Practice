package main

import (
	"os"
	"testing"
)

// func TestNewDeck(t *testing.T) {
// 	d := newDeck()

// 	if len(d) != 52 {
// 		t.Errorf("Expected Deck length of 52 but got %v", len(d))
// 	}
// }

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected Deck length of 52 but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting") 
}
