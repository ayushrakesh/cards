package main

import "fmt"

func main() {
	// var card string = "Ace of cards"
	// cards := "Ace of Spades"

	cards := []string{"Ace of Spades", "King of Heart", newCard()}
	cards = append(cards, "Six of Heart")

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Four of Diamonds"
}
