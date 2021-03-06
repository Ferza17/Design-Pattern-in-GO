package main

import (
	"container/list"
	"fmt"
)

// Observable, Observer

type Observable struct {
	sub *list.List
}

func (o *Observable) Subscribe(x Observer) {
	o.sub.PushBack(x)
}

func (o *Observable) UnSubscribe(x Observer) {
	for z := o.sub.Front(); z != nil; z = z.Next() {
		if z.Value.(Observer) == x {
			o.sub.Remove(z)
		}
	}
}

func (o *Observable) Fire(data interface{}) {
	for z := o.sub.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

type Observer interface {
	Notify(data interface{})
}

type PropertyChange struct {
	Name  string
	Value interface{}
}

type Person struct {
	Observable
	age int
}

func NewPerson(age int) *Person {
	return &Person{
		Observable{new(list.List)},
		age,
	}
}

func (p *Person) Age() int {
	return p.age
}

func (p *Person) SetAge(age int) {
	if age == p.age {
		return
	}

	oldCanVote := p.CanVote()

	p.age = age
	p.Fire(PropertyChange{"Age", p.age})

	if oldCanVote != p.CanVote() {
		p.Fire(PropertyChange{"CanVote", p.CanVote()})
	}
}

func (p *Person) CanVote() bool {
	return p.age >= 18
}

type ElectoralRoll struct {
}

func (e *ElectoralRoll) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Name == "CanVote" && pc.Value.(bool) {
			fmt.Println("Congratulations, you can vote!")
		}
	}
}

func main() {
	p := NewPerson(0)

	er := &ElectoralRoll{}
	p.Subscribe(er)

	for i := 10; i < 20; i++ {
		fmt.Println("Setting age to ", i)
		p.SetAge(i)
	}
}
