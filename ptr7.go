package main

import "fmt"

func main() {

	var a int = 10
	var p *int = &a // p points to a
	fmt.Println(a)
	fmt.Println(p)

	// modifying values via ptr
	aa := 10
	change(&aa)
	fmt.Println(aa) // 100
	// passing by pointer allows modification of the original variable

	// Ptr with struct
	book := Book{Title: "Go in Action"}
	modify(&book)
	fmt.Println(book.Title) // Modified

	// Nil pointers
	// Dereferencing a nil pointer will panic. Always ensure a pointer is intialized before dereferencing
	var pp *int
	fmt.Println(pp) // nil

	// new Keyword
	ptr := new(int) // allocates memory, return *int
	*ptr = 42
	fmt.Println(*ptr) // 42

	// new(T) returns a pointer to a zeroed T
	// equivalent to:
	// var x T
	// return &x

	// Pointer to function
	f := sayHi
	f() // Hi
	// All function variables are actually references

	// No Pointer Arithmetic
	// Go doesn't allow pointer arithmetic (e.g., p++, p+1) -- unlike c/c++

	// Comparing pointers
	aaa := 5
	b := 5
	p1 := &aaa
	p2 := &b
	fmt.Println(p1 == p2) // false
	// Even if values are equal addresses are different
}

// Pointer
// A ptr is a variable that stores the memory address of another variable
// & -> "address of" operator
// * -> "dereference" operator (get the value at the address )

// Pointer types
// Go is strongly typed, so if a is an int, then p must be of type *int
// You cannot assign *string to a *int

// Modifying values via ptr
func change(val *int) {
	*val = 100
}

// Pointer with struct
type Book struct {
	Title string
}

func modify(b *Book) {
	b.Title = "Modified"
}

// You can use either a *book  receiver or pass &book to modify orignal struct fields

// Pointer Receivers in Methods
func (b *Book) Rename(newTitle string) {
	b.Title = newTitle
}

// Use pointer receivers when:
// You want to modify the struct
// To avoid copying (especially for large stucts)

// Pointers to Functions
func sayHi() {
	fmt.Println("hi")
}
