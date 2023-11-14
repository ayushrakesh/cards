package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//Custom type declaration
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Heart", "Club"}
	cardValues := []string{"Ace", "One", "Two", "Three"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards

}

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// convert deck to string
func (d deck) deckToString() string {

	return strings.Join([]string(d), ",")
}

//save deck to file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.deckToString()), 0666)
}
