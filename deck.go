package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

// to convert our data to byte slice
//Step One -> convert deck to slice of strings and then to a single string
func (d deck) toString() string {
	// using the strings package to make slice of strings to a single string
	return strings.Join([]string(d), ",")
}

func (d deck) savetoFile(filename string) error {
	//write file is a funtion in the package ioutil which writes the data of pur program to harddsk
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
	//0666 ->permision that anyone can use this file
}

//extract the deck from local storage
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		//log the error amd quite the program
		fmt.Println("error:", err)
		//to quite the program use a lib os
		os.Exit(1)
	}
	//type conversion from slice of string to actual deck
	s := strings.Split(string(bs), ",")
	return deck(s)
	// type conversion what_we_want(what_we_have)
}
