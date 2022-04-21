package main

func main() {
	cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	// fmt.Println(cards.toString())

	// cards.saveTofile("my_cards")

	// cards := newDeckFromFile("my_cards")

	cards.shuffle()
	cards.print()
}
