package main

import (
	"github.com/ferza17/Design_Pattern/Factory/Factory_Function/Models/Person_Model"
	"fmt"
)

func main() {

	p := Person_Model.NewPerson("John", 33)
	fmt.Println(p)

}
