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

type Person struct {
	Observable
	Name string
}

func NewPerson(name string) *Person {
	return &Person{
		Observable{new(list.List)},
		name,
	}
}

func (p *Person) CatchACold() {
	p.Fire(p.Name)
}

type DoctorService struct {
}

func (d *DoctorService) Notify(data interface{}) {
	fmt.Printf("A Doctor has been called for %s", data.(string))
}

func main() {
	p := NewPerson("John")
	ds := &DoctorService{}

	p.Subscribe(ds)

	p.CatchACold()
}
