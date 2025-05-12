package main

import "fmt"

// The any type and Empty Interface in
/*
In Go, interface{} is called the empty interface.
It can hold values of any type because all types implement zero methods, and interface{} requires zero methods.
As of Go 1.18, the keyword any is introduced as a type alias for interface{} for better readability.
*/

// Use cases:
// Functions that accepts any type:
func PrintAnything(v any) {
	fmt.Println(v)
}
func main() {

	var x interface{} // or var x any
	x = 42
	x = "hello"
	x = []int{1, 2, 3}
	fmt.Println("x: ", x)
	// Use cases:
	// Functions that accepts any type:

	// storing mixed types in slices or maps:
	_ = []any{1, 2, 5, 71}

	// You must use type assertion or type switches to access the actual value stored in an any:

	var v any = "abhiii71"
	str, ok := v.(string)
	if ok {
		fmt.Println("String: ", str)
	}

	// or
	switch val := v.(type) {
	case string:
		fmt.Println("String: ", val)
	case int:
		fmt.Println("Int: ", val)
	}
	// Type assertion allows you to extract the actual type from an interface{} or any value.
}
