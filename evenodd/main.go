package main

import (
	"fmt"
	"strconv"
)

func main() {
	numbers := newNumber()

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println(strconv.Itoa(number) + ": even")
		} else {
			fmt.Println(strconv.Itoa(number) + ": odd")
		}
	}
}
