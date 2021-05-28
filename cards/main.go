package main

import "fmt"

func main() {
	cards := newDeck()
	cardToString := cards.toString()

	fmt.Println(cardToString)
	fmt.Println(" ")

	hand, remainingCards := deal(cards, 5)

	hand.print()
	fmt.Println(" ")
	remainingCards.print()

	// slice byte
	greeting := "hi there!"
	fmt.Println([]byte(greeting))
}
