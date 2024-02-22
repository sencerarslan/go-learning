package util

import "fmt"

func Defer() {
	example20()
	example21()
}
func example20() {
	state := true
	defer fmt.Println("1. Defer")
	if state {
		return
	}
	defer fmt.Println("2. Defer")
}

func example21() {
	state := true
	defer cleanup()
	if state {
		// panic("Bir hata oluştu!")
	}
}
func cleanup() {
	fmt.Println("cleanup")
}
