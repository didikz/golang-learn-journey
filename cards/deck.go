package main

import (
	"fmt"
	"strings"
)

// create a new type of 'deck
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spade", "Diamonds", "Hearts", "Clubs"}
	cardNumbers := []string{"Ace", "Two", "Three", "Four", "Five"}

	for _, suite := range cardSuites {
		for _, number := range cardNumbers {
			cards = append(cards, number+" of "+suite)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) hello() string {
	return "Hellow from deck"
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
