package main

func main() {
	// var card string = "Ace of cards"
	// cards := "Ace of Spades"

	// cards := deck{"Ace of Spades", "King of Heart"}
	// cards = append(cards, "Six of Heart")

	// cards := newDeck()

	// handsDeck, remainingDeck := deal(cards, 8)

	// handsDeck.print()

	// fmt.Println("-----------")

	// remainingDeck.print()

	// fmt.Println(cards.deckToString())

	// cards.saveToFile("my_cards")

	// cards := deckFromFile("my_cards")
	// cards.shuffle()
	// cards.print()

	// fmt.Println(rand.Intn(100))
	// fmt.Println(rand.Intn(100))
	// fmt.Println(time.Now().UnixNano())

	// source := rand.NewSource(time.Now().UnixNano())

	// r := rand.New(source)

	cards := newDeck()
	cards.shuffle()
	cards.print()

}

// func newCard() string {
// 	return "Four of Diamonds"
// }
