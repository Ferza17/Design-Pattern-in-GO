package Person_Model

type person struct {
	// Address
	StreetAddress, Postcode, City string

	//job
	CompanyName, Position string
	AnnualIncome          int
}

type PersonBuilder struct {
	Person *person
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}
