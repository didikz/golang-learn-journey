package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	cardToString := cards.toString()
	// save cards to file
	cards.saveToFile(getFileName())

	fmt.Println(cardToString)
	fmt.Println(" ")

	// get deck from file
	deckFromFile := newDeckfromFile(getFileName())
	fmt.Println("printing deck from file: ")
	deckFromFile.print()
	fmt.Println(" ")

	hand, remainingCards := deal(deckFromFile, 2)

	hand.print()
	fmt.Println(" ")
	remainingCards.print()
	fmt.Println(" ")

	// slice byte
	greeting := "hi there!"
	fmt.Println("printing byte of: ", greeting)
	fmt.Println([]byte(greeting))
	fmt.Println(" ")

	// try to create new type with name table
	tables := createTable()
	fmt.Println("Displaying specific table: ")
	fmt.Println(tables.getTable(1))
	fmt.Println(" ")

	// passing another type
	tables.passDeck(hand)
	fmt.Println(" ")
}
