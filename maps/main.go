package main

import "fmt"

func main() {
	// can't add values directly to movies map
	var movies map[string]string

	// use make function to initialize instead
	bandmaid := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"black": "#000000",
	}

	// adding item to map
	bandmaid["vocal"] = "Miku Kobato"
	bandmaid["guitar"] = "Kanami Tomo"

	// delete map item
	delete(bandmaid, "guitar")

	fmt.Println(colors)
	fmt.Println(movies)
	fmt.Println(bandmaid)

	// identify map key exists
	_, found := bandmaid["drum"]
	fmt.Println(found)
}
