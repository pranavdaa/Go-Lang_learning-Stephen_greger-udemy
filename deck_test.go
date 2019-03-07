//file to put all the test deck.go file

package main

import (
	"os"
	"testing"
)

//t is the test handeler is something goes wrong we tell t
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 20 but get %v", len(d))
		// % v to say what string we have got (the one whcih is nest to coma)
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first csrd of Ace of spads , but got %v", d[0])
	}
}

func TestSaveToDeckAndNewDecktFromFile(t *testing.T) {
	// remove file
	os.Remove("_decktesting")
	deck := newDeck()
	//save the file
	deck.savetoFile("_decktesting")
	//loading file from disk
	loadDeck := newDeckFromFile("_decktesting")

	if len(loadDeck) != 16 {
		t.Errorf("Expected 16 cards in deck,got %v", len(loadDeck))
	}
	os.Remove("_decktesting")
}
