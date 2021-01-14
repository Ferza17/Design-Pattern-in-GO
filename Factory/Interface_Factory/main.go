package main

import (
	PersonModel "github.com/ferza17/Design_Pattern/Factory/Interface_Factory/Models/Person_Model"
)

func main() {
	p := PersonModel.NewPerson("John", 134)
	p.SayHello()

}
