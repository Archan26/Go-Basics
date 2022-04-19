package main

import "fmt"

func main() {
	cards := []string{"Ace of Spades", newCard()}
	cards = append(cards, "Six of Hearts")
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
