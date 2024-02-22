package util

func Interfaces() {
	example19()
}

type Book struct {
	desi int
}

func example19() {
	book1 := &Book{desi: 10}
	book2 := &Book{desi: 20}

	println(book1.ShippingCost())
	println(book2.ShippingCost())

	books := []Book{{desi: 10}, {desi: 20}, {desi: 30}}
	println(calculateTotalShippingCost(books))
}

func (book *Book) ShippingCost() int {
	return 5 + book.desi*2
}

func calculateTotalShippingCost(books []Book) int {
	total := 0
	for _, book := range books {
		total += book.ShippingCost()
	}
	return total
}
