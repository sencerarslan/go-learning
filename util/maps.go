package util

import (
	"fmt"
)

func Map() {
	example8()
	example9()
}

func example8() {
	a := make(map[string]int, 0)
	a["A"] = 1
	a["B"] = 2
	a["C"] = 3

	delete(a, "B")
	fmt.Println(a)
}

func example9() {
	b := map[string]int{
		"A": 1,
	}
	b["B"] = 2
	fmt.Println(b)
}
