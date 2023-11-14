package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

	//writing file to disk -> before we have to convert it to byte slice
	return ioutil.WriteFile(filename, []byte(d.deckToString()), 0666)
}

func deckFromFile(filename string) deck {

	// Read file from disk which returns byte slice
	bs, err := ioutil.ReadFile(filename)

	if err != nil {

		// log the error
		fmt.Println("Error  ->", err)

		// exit program with status 1 -> means something went wrong
		os.Exit(1)
	}

	// spliting the string made from byte slice into the slice of strings using seperator string ","
	s := strings.Split(string(bs), ",")

	return deck(s)
}
