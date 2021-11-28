package main

import "fmt"

// as long as type implement these functions, it will automatically promoted of bot type
type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanihsBot struct{}

func main() {
	eb := englishBot{}
	sb := spanihsBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there"
}

func (spanihsBot) getGreeting() string {
	return "Hola"
}
