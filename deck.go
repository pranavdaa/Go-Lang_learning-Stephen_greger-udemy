package main

import "fmt"

// custome method to deck
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Dimonds", "Hearts", "c"}
	cardValue := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// type of funtion
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//this is how u make two grps from the deck of cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
