package main

import "fmt"

type contact struct {
	email string
	phone string
}

type person struct {
	firstName string
	lastName  string
	age       int
	address   string
	contact   contact
}

func main() {
	alex := person{"Alex", "Anderson", 30, "brooklyn", contact{"alex@gmail.com", "01234556"}}                                         // rely on properties order
	moshi := person{firstName: "Moshi", lastName: "Meong", age: 4, address: "Malang", contact: contact{"moshi@gmail.com", "0123123"}} // specify based on property key

	// use this kind of declaration
	var molly person
	molly.firstName = "Molly"
	molly.lastName = "Ndut"
	molly.age = 5
	molly.address = "malang"
	molly.contact = contact{}

	fmt.Println(alex)
	fmt.Println(moshi)
	fmt.Printf("%+v", molly)
}
