package util

import "fmt"

func Struct() {
	example18()
}

type Customer struct {
	id      int64
	name    string
	age     int
	address Address
}
type Address struct {
	city     string
	district string
}

func example18() {
	var customer1 = Customer{id: 1, name: "Sencer", age: 30, address: Address{
		city: "İzmir", district: "Karşıyaka",
	}}
	var customer2 = Customer{id: 2, name: "Kubilay", age: 28, address: Address{
		city: "İzmir", district: "Bostanlı",
	}}

	customer1.ToString()
	ChangeName(&customer1)
	customer1.ToString()
	customer2.ToString()
}

func (customer *Customer) ToString() {
	fmt.Println(customer.name, customer.age, customer.address.city)
}

func ChangeName(customer *Customer) {
	customer.name = "Ali"
}
