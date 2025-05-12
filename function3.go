package main

import "fmt"

func main() {
	fmt.Println(add(70, 1))
	fmt.Print("Divide function: ")
	rem, quo := divide(71, 3)

	fmt.Println("Remainder: ", rem, " Quotient: ", quo)

	// name return values
	fmt.Println(addd(1, 6))

	// Calling the variadic function
	fmt.Println(sum(1, 2, 4, 71)) // output 78
	fmt.Println(sum(10, 20))      //30
	fmt.Println(sum())            // 0

	// Passing variadic arguments
	arr := []int{1, 2, 3, 4, 71}
	// variadic functions allow you to pass a slice by using ... to expand a slice into individual elements.
	fmt.Println(sum(arr...))

	// variadic functions with multiple types
	printAll(1, "abhiii71", true)

}

// function syntax and return values
// in go the function syntax is straightforward. Here's how to define a function with parameters and return value:
func add(a int, b int) int {
	return a + b
}

// Multiple Return Values
// Go allows multiple return values, making it unique compared to many other languages.
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// Name return values
// Go allows you to specify named return values within the function signature. These values can be used as part
// of the return statement without explicitly naming them
func addd(a, b int) (sum int) {
	sum = a + b
	return // Implicitly returns the value of 'sum'
}

// Here sum is defined as the return value
// the return statement is implicitly returning the sum variable without needing to specify it
// Name return values are particularly useful for improving code clarity , especially when dealing with multiple return values or more complex logic.

// Variadic Functions
// A variadic function is a function that can accept a variable number of arguments of a specific type. In go, you
// define this using ...(ellipsis).
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// the ...int syntax allows the functions to accept zero or more integers as arguments
// inside the function, nums behave like a slice ([]int)
// You can call sum with any number of integers, and it will sum them up.

// Variadic Functions with Multiple Types
// Go doesn't support variadic functions with multiple argument types directly (like int, string, bool).
// However, you can use an interface{} to accept different types, though you'll need to handle types assertion inside the function
func printAll(values ...interface{}) {
	for _, v := range values {
		fmt.Println(v)
	}
}

// interface{} is Go's way of representing any type
// inside the function, you may need type assertion if  you need to work with specific types.
