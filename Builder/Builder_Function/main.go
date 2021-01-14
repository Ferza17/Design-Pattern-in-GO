package main

import (
	"github.com/ferza17/Design_Pattern/Builder/Builder_Function/Models/Person_Model"
	"fmt"
)

func main() {
	b := Person_Model.PersonBuilder{}
	p := b.Called("Fery").WorksAsA("Developer").Build()
	fmt.Println(p)
}
