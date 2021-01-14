package Person_Model

func (pjb *PersonJobBuilder) At(
	companyName string) *PersonJobBuilder {
	pjb.Person.CompanyName = companyName
	return pjb
}

func (pjb *PersonJobBuilder) AsA(
	position string) *PersonJobBuilder {
	pjb.Person.Position = position
	return pjb
}

func (pjb *PersonJobBuilder) Earning(
	annualIncome int) *PersonJobBuilder {
	pjb.Person.AnnualIncome = annualIncome
	return pjb
}
