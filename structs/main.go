package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	address   string
}

func main() {
	alex := person{"Alex", "Anderson", 30, "brooklyn"}                                // rely on properties order
	moshi := person{firstName: "Moshi", lastName: "Meong", age: 4, address: "Malang"} // specify based on property key

	// use this kind of declaration
	var molly person
	molly.firstName = "Molly"
	molly.lastName = "Ndut"
	molly.age = 5
	molly.address = "malang"

	fmt.Println(alex)
	fmt.Println(moshi)
	fmt.Printf("%+v", molly)
}
