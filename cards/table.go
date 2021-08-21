package main

import "fmt"

type table []string

func createTable() table {
	tables := table{"table 1", "table 2"}

	return tables
}

func (t table) getFirstTable() string {
	return t[0]
}

func (t table) getTable(number int) string {
	return t[number]
}

func (t table) passDeck(deck []string) {
	for _, card := range deck {
		fmt.Println(card)
	}
}
