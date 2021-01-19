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

	p.age = age
	p.Fire(PropertyChange{"Age", p.age})
}

type TrafficManagement struct {
	o Observable
}

func (t *TrafficManagement) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Value.(int) >= 18 {
			fmt.Println("Congrats, you can drive now!")
			t.o.UnSubscribe(t)
		}
	}
}

func main() {
	p := NewPerson(15)
	t := &TrafficManagement{p.Observable}
	p.Subscribe(t)

	for i := 16; i < 20; i++ {
		fmt.Println("Setting the age to ", i)
		p.SetAge(i)
	}
}
