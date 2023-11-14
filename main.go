package main

func main() {
	// var card string = "Ace of cards"
	// cards := "Ace of Spades"

	// cards := deck{"Ace of Spades", "King of Heart"}
	// cards = append(cards, "Six of Heart")

	cards := newDeck()

	// handsDeck, remainingDeck := deal(cards, 8)

	// handsDeck.print()

	// fmt.Println("-----------")

	// remainingDeck.print()

	// fmt.Println(cards.deckToString())

	cards.saveToFile("my_cards")
}

// func newCard() string {
// 	return "Four of Diamonds"
// }
