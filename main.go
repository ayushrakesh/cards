package main

func main() {
	// var card string = "Ace of cards"
	// cards := "Ace of Spades"

	cards := deck{"Ace of Spades", "King of Heart"}
	cards = append(cards, "Six of Heart")

	cards.print()

}

// func newCard() string {
// 	return "Four of Diamonds"
// }
