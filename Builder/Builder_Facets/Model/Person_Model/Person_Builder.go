package Person_Model

// Initilize Struct
func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&person{}}
}

// Person Builder

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func (b *PersonBuilder) Build() *person {
	return b.Person
}
