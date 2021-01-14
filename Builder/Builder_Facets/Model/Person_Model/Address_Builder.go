package Person_Model

// Address Builder
func (it *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	it.Person.StreetAddress = streetAddress
	return it
}

func (it *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	it.Person.City = city
	return it
}

func (it *PersonAddressBuilder) WithPostCode(postcode string) *PersonAddressBuilder {
	it.Person.Postcode = postcode
	return it
}
