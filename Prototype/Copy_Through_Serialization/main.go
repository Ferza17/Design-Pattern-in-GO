package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	if err := e.Encode(p); err != nil {
		panic(err)
	}

	d := gob.NewDecoder(&b)
	result := Person{}

	if err := d.Decode(&result); err != nil {
		panic(err)
	}

	return &result

}

func main() {
	john := Person{
		"John",
		&Address{
			"123",
			"London Road",
			"UK"},
		[]string{"Maria", "Lucy"},
	}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker ST"
	jane.Friends = append(jane.Friends, "Karen")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
