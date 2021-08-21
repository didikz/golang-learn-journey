package main

import "fmt"

func main() {
	cards := newDeck()
	cardToString := cards.toString()

	fmt.Println(cardToString)
	fmt.Println(" ")

	hand, remainingCards := deal(cards, 2)

	hand.print()
	fmt.Println(" ")
	remainingCards.print()

	// slice byte
	greeting := "hi there!"
	fmt.Println([]byte(greeting))

	// try to create new type with name table
	tables := createTable()
	fmt.Println(tables.getTable(1))
	fmt.Println(" ")

	// passing another type
	tables.passDeck(hand)
	fmt.Println(" ")
}
