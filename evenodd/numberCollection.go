package main

type numberCollection []int

func newNumber() numberCollection {
	numbers := numberCollection{}

	slices := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, slice := range slices {
		numbers = append(numbers, slice)
	}
	return numbers
}