package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()
	hand, deck := deal(cards, 5)

	hand.print()
	deck.print()

	fmt.Println(cards.toString())

}

func newCard() string {
	return "Five of Diamonds"
}
