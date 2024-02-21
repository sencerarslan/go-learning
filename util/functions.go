package util

func Function() {

	addTotal := example10(20, 30)
	println(addTotal)

	total, dif, multiply := example11(20, 30)
	println(total, dif, multiply)

	var numbers = []int{10, 20, 30, 1, 2, 3, 4}
	sumResult := example12(numbers)
	println(sumResult)

	println(example13(1, 2, 3, 4))
}

func example10(x int, y int) int {
	return x + y
}

func example11(x int, y int) (int, int, int) {
	return x + y, x - y, x * y
}

func example12(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func example13(number ...int) int {
	sum := 0
	for _, value := range number {
		sum += value
	}
	return sum
}
