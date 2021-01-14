package main

import (
	"fmt"
	"github.com/ferza17/Design_Pattern/Builder/Builder_Facets/Model/Person_Model"
)

func main() {
	pb := Person_Model.NewPersonBuilder()
	pb.
		Lives().
		At("Ringroad Utara").
		In("Yogyakarta").
		WithPostCode("ABC123").
		Works().
		At("PT....").
		AsA("Freelance Programmer").
		Earning(5000000)

	person := pb.Build()

	fmt.Println("Person : ", person)
}
