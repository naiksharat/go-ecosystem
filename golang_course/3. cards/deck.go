package main

import (
	"fmt"
	"strings"
)

type deck []string

func (d deck) print() {
	for card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Hears", "Diamonds", "Clubs", "Spades"}
	cardNumbers := []string{"Ace", "One", "Two", "Three"}

	for _, cardSuit := range cardSuits {
		for _, cardNumber := range cardNumbers {
			cards = append(cards, cardNumber+" of "+cardSuit)
		}
	}

	return cards
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",") //type conversion / type casting
}
