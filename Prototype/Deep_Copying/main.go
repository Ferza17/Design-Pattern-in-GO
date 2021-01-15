package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}

func main() {
	john := Person{
		"John",
		&Address{
			"123",
			"London Road",
			"UK"}}

	// Bad Code
	//jane := john
	//jane.Name = "Jane" // ok
	//jane.Address.StreetAddress = "321 Baker st"
	// -->  The problem is John and Jane has the same pointer of Address

	// Good Code ( DEEP COPYING )
	jane := john
	jane.Address = &Address{
		john.Address.StreetAddress,
		john.Address.City,
		john.Address.Country,
	}
	jane.Name = "Jane" // ok
	jane.Address.StreetAddress = "321 Baker st"

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
