package util

import "fmt"

var globalState = true

func VariableScope() {
	example22() // Block Scope
	example23() // Func Scope state ?
	example24() // Global Scope
	fmt.Println(globalState)
}
func example22() {
	state := true
	if state {
		// Block Scope
		var x = 10
		fmt.Println(x)
	}
	// fmt.Println(x)
}

func example23() {
	// Func Scope
	state := true
	if state {
		fmt.Println(state)
	}
	fmt.Println(state)
}
func example24() {
	// Global Scope
	if globalState {
		fmt.Println(globalState)
	}
	fmt.Println(globalState)
}
