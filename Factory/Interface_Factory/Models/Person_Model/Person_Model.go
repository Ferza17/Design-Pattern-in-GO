package Person_Model

import "fmt"

type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

type tiredPerson struct {
	name string
	age  int
}

func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{name: name, age: age}
	}
	return &person{
		name: name,
		age:  age,
	}
}

func (p *person) SayHello() {
	fmt.Printf("Hi my name is %s, I am %d years old\n", p.name, p.age)
}

func (p *tiredPerson) SayHello() {
	fmt.Println("Sorry i can't talk to you")
}
