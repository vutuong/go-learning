package main

//một builder thì đủ cho việc xây dựng một particular objects
// Nhưng vẫn có trườn hợp cần nhiều hơn 1 builder
// ví dụ nếu muốn phân tách các process của việc xây dựng
// các aspects của một type nhất định

type Person struct {
	//address
	StreetAddress, Postcode, City string

	//job
	CompanyName, Position string
	AnnulaIncome          int
}

type PersonBuilder struct {
	person *Person
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}
func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (b *PersonAddressBuilder) At(StreetAddress string) *PersonAddressBuilder {
	it.person.StreetAddress = StreetAddress
	return it
}
func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	it.person.City = city
	return it
}

func (b *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	it.person.Postcode = postcode
	return it
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (pjb *PersonJobBuilder) At(StreetAddress string) *PersonAddressBuilder {
	it.person.StreetAddress = StreetAddress
	return it
}
func (pjb *PersonJobBuilder) In(city string) *PersonAddressBuilder {
	it.person.City = city
	return it
}

func (pjb *PersonJobBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	it.person.Postcode = postcode
	return it
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}
func main() {

}
