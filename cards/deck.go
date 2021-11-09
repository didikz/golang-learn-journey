package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck
// which is a slice of strings
type deck []string

func getFileName() string {
	return "table_cards"
}

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

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckfromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(byteSlice), ","))
}

func (d deck) shuffle() {
	// create random generator seed source to make sure its always get different result
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		//swap value
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
