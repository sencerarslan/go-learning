package util

import "fmt"

func Pointer() {
	example14()
	example15()

	var ex int = 10
	example16(&ex)
	fmt.Println(ex)

	var numbers = []int{1, 2, 3}
	example17(numbers)
	fmt.Println(numbers)
}

func example14() {

	var a int
	var p *int
	a = 10

	p = &a

	*p = 20

	fmt.Println(a, p, *p)
}
func example15() {
	var a int = 10
	var b int
	var p *int

	b = a
	p = &a

	*p = 20

	fmt.Println(a, b)
}
func example16(x *int) {
	*x += 12
}
func example17(numbers []int) {
	numbers[0] = 1000
}
